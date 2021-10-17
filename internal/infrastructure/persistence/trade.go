package persistence

import (
	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/domain"
)

type postgresTradeRepository struct {
	db *gorm.DB
}

func NewPostgresTradeRepository(db *gorm.DB) domain.TradeRepository {
	return &postgresTradeRepository{db}
}

func (repo *postgresTradeRepository) Save(trade *domain.Trade) error {
	tx := repo.db.Save(trade)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *postgresTradeRepository) UpdateTradeStatus(tradeID uint, status domain.TradeStatus) error {
	return repo.UpdateTradeStatusWithAnnotation(tradeID, status, "")
}

func (repo *postgresTradeRepository) UpdateTradeStatusWithAnnotation(tradeID uint, status domain.TradeStatus, annotation string) error {
	updateFields := &domain.Trade{Status: status, Annotation: annotation}

	tx := repo.db.Model(&domain.Trade{}).Where("id = ?", tradeID).Updates(updateFields)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *postgresTradeRepository) UpdateResourceItem(resID uint, quantity uint) error {
	updateFields := map[string]interface{}{"quantity": quantity}

	tx := repo.db.Model(&domain.Resource{}).Where("id = ?", resID).Updates(updateFields)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *postgresTradeRepository) FindByID(tradeID uint) (*domain.Trade, error) {
	var trade domain.Trade

	tx := repo.db.First(&trade, tradeID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	err := repo.db.Model(&trade).Association("Sender").Find(&trade.Sender)
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&trade).Association("Receiver").Find(&trade.Receiver)
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&trade.Sender).Association("TradeResources").Find(&trade.Sender.TradeResources)
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&trade.Receiver).Association("TradeResources").Find(&trade.Receiver.TradeResources)
	if err != nil {
		return nil, err
	}

	return &trade, nil
}

func (repo *postgresTradeRepository) FindResourceByID(resID uint) (*domain.Resource, error) {
	var res domain.Resource

	err := repo.db.Where("id = ?", resID).First(&res).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&res).Association("Item").Find(&res.Item)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
