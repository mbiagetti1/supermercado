package services

import (
	"errors"
	"example/services/models"
	"fmt"
)

var (
	ErrAlreadyExist = errors.New("error: item already exist")
)

var Products []models.Product
var NextID int

// read
func Get() []models.Product {
	return Products
}

func GetById(id int) models.Product {
	var product models.Product
	for _, p := range Products {
		if p.Id == id {
			product = p
			break
		}
	}
	return product
}

func GetPriceGt(price float64) (result []models.Product) {
	for _, p := range Get() {
		if p.Price > price {
			result = append(result, p)
		}
	}
	return
}

func ExistsProductId(id int) bool {
	for _, p := range Products {
		if p.Id == id {
			return true
		}
	}
	return false

}

func ExistProductName(name string) bool {
	for _, p := range Products {
		if p.Name == name {
			return true
		}
	}

	return false
}

// write
func Create(id int, name string, quantity int, code_value string, is_published bool, expiration string, price float64) (models.Product, error) {
	// validations
	if ExistProductName(name) {
		return models.Product{}, fmt.Errorf("%w. %s", ErrAlreadyExist, "product name not unique")
	}

	product := models.Product{
		Id:           NextID,
		Name:         name,
		Quantity:     quantity,
		Code_value:   code_value,
		Is_published: is_published,
		Expiration:   expiration,
		Price:        price,
	}

	Products = append(Products, product)
	NextID++
	return product, nil
}
