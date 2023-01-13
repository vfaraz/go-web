package store

import (
	"encoding/json"
	"os"
	"rest/internal/domain"
)

type jsonStorage struct {
	pathFile string
}

type Storage interface {
	GetStorage() ([]domain.Product, error)
	SetStorage([]domain.Product) error
}

func NewStorage(pathFile string) Storage {
	return jsonStorage{pathFile: pathFile}
}

// GetStorage implements Storage
func (s jsonStorage) GetStorage() ([]domain.Product, error) {
	var products []domain.Product
	file, err := os.ReadFile(s.pathFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// SetStorage implements Storage
func (s jsonStorage) SetStorage(products []domain.Product) error {
	bytes, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathFile, bytes, 0644)

}
