package infrastructure

import (
	"errors"

	"github.com/DubrovskijRD/budget_assistant_go/domain/interfaces"
	"github.com/DubrovskijRD/budget_assistant_go/domain/interfaces/repositories"

	repoImpl "github.com/DubrovskijRD/budget_assistant_go/infrastructure/repositories"
	"gorm.io/gorm"
)

type transaction struct {
	repo repositories.ReceiptRepo
}

func (t transaction) ReceiptRepo() repositories.ReceiptRepo {
	return t.repo
}

type UoW struct {
	db gorm.DB
}

func NewUow(db gorm.DB) UoW {
	return UoW{db: db}
}

func (u *UoW) Do(fn func(interfaces.TX) error) (err error) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknow transaction panic")
			}
		}
	}()
	err = fn(transaction{repoImpl.NewReceiptRepo(*tx)})
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
