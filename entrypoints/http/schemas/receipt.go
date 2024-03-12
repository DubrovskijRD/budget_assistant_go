package schemas

import (
	"time"

	"github.com/DubrovskijRD/budget_assistant_go/domain/entities"
)

type Receipt struct {
	ID          entities.ReceiptId `json:"id,omitempty"`
	BudgetId    string             `json:"budget_id"`
	Amount      entities.Amount    `json:"amount"`
	FactAmount  entities.Amount    `json:"fact_amount,omitempty"`
	Description string             `json:"description"`
	Labels      []string           `json:"labels"`
	Items       []ReceiptItem      `json:"items"`
	Date        time.Time          `json:"date"`
}

type ReceiptItem struct {
	ID     entities.ReceiptItemId `json:"id,omitempty"`
	Name   string                 `json:"name"`
	Amount entities.Amount        `json:"amount"`
	Qty    int                    `json:"qty"`
}

func RespFromReceipt(r entities.Receipt) Receipt {
	return Receipt{
		ID:          r.Id(),
		BudgetId:    r.BudgetId(),
		Amount:      r.Amount(),
		FactAmount:  r.FactAmount(),
		Description: r.Description(),
		Labels:      r.Labels(),
		Items:       itemsFromReceipt(r.Items()),
		Date:        r.Date(),
	}
}

func RespFromReceiptList(receipts []entities.Receipt) []Receipt {
	list := make([]Receipt, 0, len(receipts))
	for _, r := range receipts {
		list = append(list, RespFromReceipt(r))
	}
	return list
}

func itemsFromReceipt(items []entities.ReceiptItem) []ReceiptItem {
	list := make([]ReceiptItem, 0, len(items))
	for _, i := range items {
		list = append(list, ReceiptItem{ID: i.Id(), Amount: i.Amount(), Qty: i.Qty()})
	}
	return list
}
