package product

import (
	"errors"
	"fmt"

	//"fmt"
	"rest/internal/domain"
	"rest/pkg/store"
)

var (
	ErrNotFound = errors.New("item not found")
)

type Repository interface {
	// read
	Get() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	ExistCode(code int) (bool, error)
	// write
	Create(domain.Product) (int, error)
	Update(domain.Product) (domain.Product, error)
	//PartialUpdate(domain.Product) (domain.Product, error)
	Delete(id int) error
}
type repository struct {
	storage store.Storage
	//db *[]domain.Product
	// config
}

func NewRepository(storage store.Storage) Repository {
	return &repository{storage: storage}
}

func (r *repository) Get() ([]domain.Product, error) {
	return r.storage.GetStorage()
}

func (r *repository) GetByID(id int) (product domain.Product, err error) {
	db, err := r.storage.GetStorage()
	if err != nil {
		return
	}
	for _, product := range db {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("%w. %s", ErrNotFound, "product does not exist")
}

func (r *repository) ExistCode(code int) (bool, error) {
	db, err := r.storage.GetStorage()
	if err != nil {
		return false, err
	}
	for _, product := range db {
		if product.CodeValue == code {
			return true, nil
		}
	}
	return false, nil
}
func (r *repository) Create(product domain.Product) (lastID int, err error) {
	db, err := r.storage.GetStorage()
	if err != nil {
		return
	}
	lastID = len(db) + 1
	product.ID = lastID
	db = append(db, product)
	err = r.storage.SetStorage(db)
	if err != nil {
		return
	}
	return lastID, nil
}

func (r *repository) Update(productUpdate domain.Product) (domain.Product, error) {
	db, err := r.storage.GetStorage()
	if err != nil {
		return domain.Product{}, err
	}
	for i, product := range db {
		if product.ID == productUpdate.ID {
			flag, err := r.ExistCode(productUpdate.CodeValue)
			if err != nil {
				return domain.Product{}, err
			}

			if flag && product.CodeValue != productUpdate.CodeValue {
				return domain.Product{}, errors.New("code value already exists")
			}
			db[i] = productUpdate
			err = r.storage.SetStorage(db)
			if err != nil {
				return domain.Product{}, err
			}
			return productUpdate, nil
		}
	}
	return domain.Product{}, fmt.Errorf("%w. %s", ErrNotFound, "product does not exist")
}

//func (r *repository) PartialUpdate(productUpdate domain.Product) (domain.Product, error) {
//	for i, product := range *r.db {
//		if product.ID == productUpdate.ID {
//			if !r.ExistCode(productUpdate.CodeValue) && product.CodeValue != productUpdate.CodeValue {
//				return domain.Product{}, errors.New("code value already exists")
//			}
//			(*r.db)[i] = productUpdate
//			return productUpdate, nil
//		}
//	}
//	return domain.Product{}, fmt.Errorf("%w. %s", ErrNotFound, "product does not exist")
//}

func (r *repository) Delete(id int) (err error) {
	db, err := r.storage.GetStorage()
	if err != nil {
		return err
	}
	for i, product := range db {
		if product.ID == id {
			db = append(db[:i], db[i+1:]...)
			err = r.storage.SetStorage(db)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("%w. %s", ErrNotFound, "product does not exist")
}
