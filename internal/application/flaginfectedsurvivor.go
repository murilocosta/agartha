package application

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type FlagInfectedSurvivor struct {
	survRepo  domain.SurvivorRepository
	infecRepo domain.InfectionRepository

	invServ *domain.InventoryService
}

func NewFlagInfectedSurvivor(survRepo domain.SurvivorRepository, infecRepo domain.InfectionRepository, invServ *domain.InventoryService) *FlagInfectedSurvivor {
	return &FlagInfectedSurvivor{survRepo, infecRepo, invServ}
}

func (ucase *FlagInfectedSurvivor) Invoke(infection *dto.ReportedInfection) error {
	validate := validator.New()

	if err := validate.Struct(infection); err != nil {
		msg := core.NewErrorMessage(dto.ReportInfectedSurvivorFailed, "report infected survivor failed", http.StatusBadRequest)
		msg.AddErrorDetail(err, dto.SurvivorWriteErrorBuilder)
		return msg
	}

	if infection.ReportedID == infection.ReporteeID {
		return core.NewErrorMessage(dto.SurvivorCannotBeFlagged, "cannot flag yourself", http.StatusBadRequest)
	}

	reportee, err := ucase.survRepo.FindByID(infection.ReporteeID)
	if err != nil {
		return buildSurvivorNotFoundError(infection.ReporteeID)
	}

	reported, err := ucase.survRepo.FindByID(infection.ReportedID)
	if err != nil {
		return buildSurvivorNotFoundError(infection.ReportedID)
	}

	if duplicated := ucase.infecRepo.HasReported(reportee.ID, reported.ID); duplicated {
		return core.NewErrorMessage(dto.SurvivorAlreadyFlagged, "survivor cannot be flagged twice", http.StatusBadRequest)
	}

	if count, err := ucase.infecRepo.CountReports(reported.ID); err != nil {
		return err
	} else if count == (domain.FlagInfectedSurvivorMax + 1) {
		if err := ucase.flagSurvivorAsInfectedAndTransferInventory(reportee, reported); err != nil {
			return err
		}
	} else if count > domain.FlagInfectedSurvivorMax {
		detail := fmt.Sprintf("cannot flag a survivor more than %d times", count)
		return core.NewErrorMessage(dto.SurvivorCannotBeFlagged, detail, http.StatusBadRequest)
	}

	return ucase.infecRepo.Save(reportee, reported, infection.Annotation)
}

func (ucase *FlagInfectedSurvivor) flagSurvivorAsInfectedAndTransferInventory(reportee *domain.Survivor, reported *domain.Survivor) error {
	// Should transfer inventory to reportee
	err := ucase.invServ.TransferInventory(reported.ID, reportee.ID)

	if err != nil {
		detail := fmt.Sprintf("cannot transfer inventory from '%s' to '%s'", reported.Name, reportee.Name)
		return core.NewErrorMessage(dto.SurvivorInventoryTransferFailed, detail, http.StatusBadRequest)
	}

	reported.Infected = true
	err = ucase.survRepo.Save(reported)
	if err != nil {
		detail := fmt.Sprintf("cannot flag '%s' as infected", reported.Name)
		return core.NewErrorMessage(dto.SurvivorCannotBeFlagged, detail, http.StatusBadRequest)
	}

	return nil
}

func buildSurvivorNotFoundError(survivorID uint) error {
	detail := fmt.Sprintf("could not find survivor with ID %d", survivorID)
	return core.NewErrorMessage(dto.SurvivorNotFound, detail, http.StatusBadRequest)
}
