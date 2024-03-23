package services

import (
	"time"

	"github.com/DubrovskijRD/budget_assistant_go/domain/entities"
	"github.com/DubrovskijRD/budget_assistant_go/domain/interfaces"
	"github.com/DubrovskijRD/budget_assistant_go/domain/interfaces/repositories"
)

type ReceiptService struct {
	uow interfaces.UoW
}

func NewReceiptService(uow interfaces.UoW) ReceiptService {
	return ReceiptService{uow: uow}
}

// todo replace repositories repositories.ReceiptAdd with other dto
func (s *ReceiptService) AddReceipt(rIn repositories.ReceiptAdd) (entities.Receipt, error) {
	// todo validate rIn
	var receipt entities.Receipt
	err := s.uow.Do(func(tx interfaces.TX) error {
		r, err := tx.ReceiptRepo().Add(rIn)
		if err != nil {
			return err
		}
		receipt = r
		return err
	})
	return receipt, err
}

func (s *ReceiptService) GetLabels(budgetId string) ([]string, error) {
	var labels []string
	err := s.uow.Do(func(tx interfaces.TX) error {
		l, err := tx.ReceiptRepo().LabelList(budgetId)
		if err != nil {
			return err
		}
		labels = l
		return nil

	})
	return labels, err

}

func (s *ReceiptService) GetReceipts(budgetId string, labels []string, dateFrom time.Time, dateTo time.Time) ([]entities.Receipt, error) {
	var receipts []entities.Receipt
	err := s.uow.Do(func(tx interfaces.TX) error {
		r, err := tx.ReceiptRepo().List(budgetId, labels, dateFrom, dateTo)
		if err != nil {
			return err
		}
		receipts = r
		return nil
	})
	return receipts, err

}
