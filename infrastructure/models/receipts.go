package models

import (
	"time"

	"github.com/DubrovskijRD/budget_assistant_go/domain/entities"
	"gorm.io/gorm"
)

type ReceiptLabel struct {
	Name      string             `gorm:"primaryKey"`
	ReceiptID entities.ReceiptId `gorm:"primaryKey"`
}

type ReceiptItem struct {
	gorm.Model
	ID        entities.ReceiptItemId
	Name      string
	Amount    entities.Amount
	Qty       int
	ReceiptID entities.ReceiptId
}

type Receipt struct {
	gorm.Model
	ID          entities.ReceiptId
	BudgetId    string
	Amount      entities.Amount
	FactAmount  entities.Amount
	Description string
	Labels      []ReceiptLabel
	Items       []ReceiptItem
	Date        time.Time
}
