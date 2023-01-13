package product

import (
	"errors"
	"fmt"
	"rest/internal/domain"
)

var (
	ErrNotFound = errors.New("item not found")
)

type Repository interface {
	// read
	Get() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	ExistCode(code int) bool
	// write
	Create(domain.Product) (int, error)
	Update(int, domain.Product) (domain.Product, error)
	PartialUpdate(domain.Product) (domain.Product, error)
	Delete(id int) error
}
type repository struct {
	db *[]domain.Product
	// config
	lastID int
}

func NewRepository(db *[]domain.Product, lastID int) Repository {
	return &repository{db: db, lastID: lastID}
}

func (r *repository) Get() ([]domain.Product, error) {
	return *r.db, nil
}
func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, product := range *r.db {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("%w. %s", ErrNotFound, "product does not exist")
}
func (r *repository) ExistCode(code int) bool {
	for _, product := range *r.db {
		if product.CodeValue == code {
			return true
		}
	}
	return false
}
func (r *repository) Create(product domain.Product) (int, error) {
	r.lastID++
	product.ID = r.lastID
	*r.db = append(*r.db, product)
	return r.lastID, nil
}

func (r *repository) Update(id int, productUpdate domain.Product) (domain.Product, error) {
	for i, product := range *r.db {
		if product.ID == id {
			if !r.ExistCode(productUpdate.CodeValue) && product.CodeValue != productUpdate.CodeValue {
				return domain.Product{}, errors.New("code value already exists")
			}
			productUpdate.ID = id
			(*r.db)[i] = productUpdate
			return productUpdate, nil
		}
	}
	return domain.Product{}, fmt.Errorf("%w. %s", ErrNotFound, "product does not exist")
}

func (r *repository) PartialUpdate(productUpdate domain.Product) (domain.Product, error) {
	for i, product := range *r.db {
		if product.ID == productUpdate.ID {
			(*r.db)[i] = productUpdate
			return productUpdate, nil
		}
	}
	return domain.Product{}, fmt.Errorf("%w. %s", ErrNotFound, "product does not exist")
}

func (r *repository) Delete(id int) (err error) {
	for i, product := range *r.db {
		if product.ID == id {
			*r.db = append((*r.db)[:i], (*r.db)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("%w. %s", ErrNotFound, "product does not exist")
}
