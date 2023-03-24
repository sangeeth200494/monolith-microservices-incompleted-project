package products

import "errors"

var ErrNotFound = errors.New("product not found")

type repository interface {
	save(*product) error
	ByID(ID) (*Product, error)
}
