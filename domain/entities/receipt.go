package entities

import (
	"fmt"
	"time"
)

type ReceiptId uint
type ReceiptItemId uint
type Amount uint

type ReceiptItem struct {
	id     ReceiptItemId
	name   string
	amount Amount
	qty    int
}

func (ri ReceiptItem) Id() ReceiptItemId {
	return ri.id
}

func (ri ReceiptItem) Name() string {
	return ri.name
}

func (ri ReceiptItem) Amount() Amount {
	return ri.amount
}

func (ri ReceiptItem) Qty() int {
	return ri.qty
}

func (ri ReceiptItem) TotalAmount() Amount {
	return ri.amount * Amount(ri.qty)
}

func NewReceiptItem(id ReceiptItemId, name string, amount Amount, qty int) ReceiptItem {
	return ReceiptItem{id: id, name: name, amount: amount, qty: qty}
}

type Receipt interface {
	Id() ReceiptId
	BudgetId() string
	Amount() Amount
	FactAmount() Amount
	Description() string
	Labels() []string
	Items() []ReceiptItem
	Date() time.Time
}

type receipt struct {
	id          ReceiptId
	budgetId    string
	amount      Amount
	factAmount  Amount
	description string
	labels      []string
	items       []ReceiptItem
	date        time.Time
}

func (r receipt) Id() ReceiptId {
	return r.id
}

func (r receipt) BudgetId() string {
	return r.budgetId
}

func (r receipt) Amount() Amount {
	return r.amount
}

func (r receipt) FactAmount() Amount {
	return r.factAmount
}

func (r receipt) Description() string {
	return r.description
}

func (r receipt) Labels() []string {
	return r.labels
}

func (r receipt) Items() []ReceiptItem {
	return r.items
}

func (r receipt) Date() time.Time {
	return r.date
}

func NewReceipt(id ReceiptId, budgetId string, amount Amount, description string, labels []string, items []ReceiptItem, date time.Time) (Receipt, error) {
	r := receipt{
		id: id, budgetId: budgetId, amount: amount, description: description, labels: labels, date: date, items: make([]ReceiptItem, 0, len(items)),
	}
	for _, ri := range items {
		err := addItem(&r, ri)
		if err != nil {
			return nil, err
		}
	}
	return &r, nil
}

func addItem(r *receipt, ri ReceiptItem) error {
	updatedFactAmount := r.factAmount + ri.TotalAmount()
	if r.amount < updatedFactAmount {
		return fmt.Errorf("invalid receipt amount: %v < %v", r.amount, updatedFactAmount)
	}
	r.factAmount = updatedFactAmount
	r.items = append(r.items, ri)
	return nil
}
