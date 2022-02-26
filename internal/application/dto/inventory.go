package dto

import "github.com/murilocosta/agartha/internal/domain"

type ItemRead struct {
	ID     uint              `json:"id,omitempty"`
	Name   string            `json:"name"`
	Icon   string            `json:"icon"`
	Price  int32             `json:"price"`
	Rarity domain.ItemRarity `json:"rarity"`
}

type InventoryItemRead struct {
	ResourceID uint      `json:"resource_id"`
	Quantity   uint      `json:"quantity"`
	Item       *ItemRead `json:"item"`
}

type InventoryRead struct {
	ID         uint                 `json:"id"`
	SurvivorID uint                 `json:"survivor_id"`
	Items      []*InventoryItemRead `json:"items"`
}

func ConvertToInventoryRead(inv *domain.Inventory) *InventoryRead {
	var invItems []*InventoryItemRead
	for _, res := range inv.Resources {
		item := &InventoryItemRead{
			ResourceID: res.ID,
			Quantity:   res.Quantity,
			Item: &ItemRead{
				ID:     res.Item.ID,
				Name:   res.Item.Name,
				Icon:   res.Item.Icon,
				Price:  res.Item.Price,
				Rarity: res.Item.Rarity,
			},
		}
		invItems = append(invItems, item)
	}

	return &InventoryRead{
		ID:         inv.ID,
		SurvivorID: inv.OwnerID,
		Items:      invItems,
	}
}

func ConvertToItemRead(item *domain.Item) *ItemRead {
	return &ItemRead{
		ID:     item.ID,
		Name:   item.Name,
		Icon:   item.Icon,
		Price:  item.Price,
		Rarity: item.Rarity,
	}
}
