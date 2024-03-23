package interfaces

import "github.com/DubrovskijRD/budget_assistant_go/domain/interfaces/repositories"

type TX interface {
	ReceiptRepo() repositories.ReceiptRepo
	// commit
	// rollback
}

type UoW interface {
	Do(fn func(TX) error) error
}
