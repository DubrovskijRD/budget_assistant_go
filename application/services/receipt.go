package services

import (
	"time"

	"github.com/DubrovskijRD/budget_assistant_go/domain/entities"
	"github.com/DubrovskijRD/budget_assistant_go/domain/interfaces/repositories"
)

type ReceiptService struct {
	receiptRepo repositories.ReceiptRepo
}

func NewReceiptService(repo repositories.ReceiptRepo) ReceiptService {
	return ReceiptService{receiptRepo: repo}
}

// todo replace repositories repositories.ReceiptAdd with other dto
func (s *ReceiptService) AddReceipt(rIn repositories.ReceiptAdd) (entities.Receipt, error) {
	// todo: add uow
	// todo validate rIn
	r, err := s.receiptRepo.Add(rIn)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *ReceiptService) GetLabels(budgetId string) ([]string, error) {
	labels, err := s.receiptRepo.LabelList(budgetId)
	if err != nil {
		return nil, err
	}
	return labels, nil
}

func (s *ReceiptService) GetReceipts(budgetId string, labels []string, dateFrom time.Time, dateTo time.Time) ([]entities.Receipt, error) {
	receipts, err := s.receiptRepo.List(budgetId, labels, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}
	return receipts, nil
}
