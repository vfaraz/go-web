package product

import (
	"errors"
	"rest/internal/domain"
)

var (
	ErrAlreadyExist = errors.New("already exist")
)

// controller
type Service interface {
	Get() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	Create(name string, quantity int, codeValue int, isPublished bool,
		expiration string, price float64) (domain.Product, error)
	Update(id int, name string, quantity int, codeValue int, isPublished bool,
		expiration string, price float64) (domain.Product, error)
	PartialUpdate(domain.Product) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	// repo
	repo Repository

	// external api's
	// ...
}

func NewService(rp Repository) Service {
	return &service{repo: rp}
}

// read
func (sv *service) Get() ([]domain.Product, error) {
	return sv.repo.Get()
}

func (sv *service) GetByID(id int) (domain.Product, error) {
	return sv.repo.GetByID(id)
}

// write
func (sv *service) Create(name string, quantity int, codeValue int, isPublished bool,
	expiration string, price float64) (domain.Product, error) {
	flag, err := sv.repo.ExistCode(codeValue)
	if err != nil {
		return domain.Product{}, err
	}
	if flag {
		return domain.Product{}, ErrAlreadyExist
	}

	product := domain.Product{
		Name:        name,
		Quantity:    quantity,
		CodeValue:   codeValue,
		IsPublished: isPublished,
		Expiration:  expiration,
		Price:       price,
	}
	lastID, err := sv.repo.Create(product)
	if err != nil {
		return domain.Product{}, err
	}

	product.ID = lastID

	return product, nil
}

func (sv *service) Update(id int, name string, quantity int, codeValue int, isPublished bool,
	expiration string, price float64) (domain.Product, error) {

	product := domain.Product{
		ID:          id,
		Name:        name,
		Quantity:    quantity,
		CodeValue:   codeValue,
		IsPublished: isPublished,
		Expiration:  expiration,
		Price:       price}

	product, err := sv.repo.Update(product)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (sv *service) PartialUpdate(product domain.Product) (domain.Product, error) {

	//product := domain.Product{
	//	Name:        name,
	//	Quantity:    quantity,
	//	CodeValue:   codeValue,
	//	IsPublished: isPublished,
	//	Expiration:  expiration,
	//	Price:       price}

	//newProduct, err := sv.repo.PartialUpdate(product)
	newProduct, err := sv.repo.Update(product)
	if err != nil {
		return domain.Product{}, err
	}
	return newProduct, nil
}

func (sv *service) Delete(id int) (err error) {
	err = sv.repo.Delete(id)
	if err != nil {
		return
	}
	return
}
