package services

import (
	"errors"
	"fmt"
	"rest/services/models"
)

var (
	ErrAlreadyExist = errors.New("error: item already exist")
)

var Products []models.Products
var CountID = 1

func ExistCode(code int) bool {
	for _, product := range Products {
		if product.CodeValue == code {
			return true
		}
	}
	return false
}
func Create(name string, quantity int, code int,
	published bool, expiration string,
	price float64) (prod models.Products, err error) {

	if ExistCode(code) == true {
		err = fmt.Errorf("%w. %s", ErrAlreadyExist, "code not unique")
		return
	}

	prod = models.Products{
		ID:          CountID,
		Name:        name,
		Quantity:    quantity,
		CodeValue:   code,
		IsPublished: published,
		Expiration:  expiration,
		Price:       price,
	}
	Products = append(Products, prod)
	CountID++
	return
}
func Get() []models.Products {
	return Products
}
func GetByID(id int) (prod models.Products, err error) {
	for _, product := range Products {
		if product.ID == id {
			prod = product
			return
		}
	}
	err = errors.New("error: product not exist")

	return
}
