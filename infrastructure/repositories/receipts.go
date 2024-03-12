package repositories

import (
	"time"

	"github.com/DubrovskijRD/budget_assistant_go/domain/entities"
	"github.com/DubrovskijRD/budget_assistant_go/domain/interfaces/repositories"
	"github.com/DubrovskijRD/budget_assistant_go/infrastructure/models"
	"gorm.io/gorm"
)

func ReceiptModel(ra repositories.ReceiptAdd) *models.Receipt {
	r := models.Receipt{
		BudgetId:    ra.BudgetId,
		Amount:      ra.Amount,
		FactAmount:  ra.FactAmount,
		Description: ra.Description,
		Items:       ItemsModels(ra.Items),
		Date:        ra.Date,
	}
	r.Labels = LabelsModels(ra.Labels, r)
	return &r
}

func ItemsModels(items []repositories.ReceiptItemAdd) []models.ReceiptItem {
	list := make([]models.ReceiptItem, 0, len(items))
	for _, i := range items {
		list = append(list, models.ReceiptItem{Name: i.Name, Amount: i.Amount, Qty: i.Qty})
	}
	return list
}

func LabelsModels(items []string, receipt models.Receipt) []models.ReceiptLabel {
	list := make([]models.ReceiptLabel, 0, len(items))
	for _, i := range items {
		list = append(list, models.ReceiptLabel{Name: i, ReceiptID: receipt.ID})
	}
	return list
}

func ReceiptFromModel(m models.Receipt) (entities.Receipt, error) {
	labels := make([]string, 0, len(m.Labels))
	for _, i := range m.Labels {
		labels = append(labels, i.Name)
	}
	return entities.NewReceipt(m.ID, m.BudgetId, m.Amount, m.Description, labels, ReceiptItemsFromModel(m.Items), m.Date)

}

func ReceiptItemsFromModel(items []models.ReceiptItem) []entities.ReceiptItem {
	list := make([]entities.ReceiptItem, 0, len(items))
	for _, i := range items {
		list = append(list, entities.NewReceiptItem(i.ID, i.Name, i.Amount, i.Qty))
	}
	return list
}

type LabeslResult struct {
	Name string
}

type ReceiptRepositoryImpl struct {
	db gorm.DB
}

func NewReceiptRepo(db gorm.DB) *ReceiptRepositoryImpl {
	return &ReceiptRepositoryImpl{db: db}
}

func (r *ReceiptRepositoryImpl) List(budgetId string, labels []string, dateFrom time.Time, dateTo time.Time) ([]entities.Receipt, error) {
	var result []models.Receipt
	q := r.db.Model(&models.Receipt{}).Preload("Labels").Preload("Items").Where(models.Receipt{BudgetId: budgetId})
	if len(labels) > 0 {
		q = q.Joins("join receipt_labels on receipts.id = receipt_labels.receipt_id").Where("receipt_labels.name in ?", labels)
	}
	// TODO: date filter
	q.Debug().Find(&result)
	receipts := make([]entities.Receipt, 0, len(result))
	for _, i := range result {
		r, err := ReceiptFromModel(i)
		if err != nil {
			return nil, err
		}
		receipts = append(receipts, r)
	}
	return receipts, nil
}

func (r *ReceiptRepositoryImpl) LabelList(budgetId string) ([]string, error) {
	var names []string
	r.db.Debug().Model(&models.ReceiptLabel{}).Distinct().Joins("join receipts on receipts.id = receipt_labels.receipt_id").Where("receipts.budget_id = ?", budgetId).Pluck("Name", &names)

	return names, nil
}

func (r *ReceiptRepositoryImpl) Add(ra repositories.ReceiptAdd) (entities.Receipt, error) {
	model := ReceiptModel(ra)
	r.db.Debug().Create(&model)
	return ReceiptFromModel(*model)
}

func (r *ReceiptRepositoryImpl) Delete(budgetId string, id entities.ReceiptId) (entities.ReceiptId, error) {
	result := r.db.Debug().Delete(&models.Receipt{ID: id, BudgetId: budgetId})
	return id, result.Error

}
