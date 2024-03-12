package repositories

import (
	"time"

	"github.com/DubrovskijRD/budget_assistant_go/domain/entities"
)

type ReceiptItemAdd struct {
	Name   string
	Amount entities.Amount
	Qty    int
}

type ReceiptAdd struct {
	BudgetId    string
	Amount      entities.Amount
	FactAmount  entities.Amount
	Items       []ReceiptItemAdd
	Description string
	Labels      []string
	Date        time.Time
}

type ReceiptRepo interface {
	List(budgetId string, labels []string, dateFrom time.Time, dateTo time.Time) ([]entities.Receipt, error)
	LabelList(budgetId string) ([]string, error)
	Add(r ReceiptAdd) (entities.Receipt, error)
	Delete(budgetId string, id entities.ReceiptId) (entities.ReceiptId, error)
}
