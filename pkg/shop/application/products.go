package application

import "errors"

type productReadModel interface {
	AllProducts() ([]products.Product, err)
}

type productsService struct {
	repo      products.Repository
	readModel productReadModel
}

func NewProductsService(repo products.Repository, readModel productReadModel) ProductsService {
	return ProductsService{repo, readModel}

}

func (s ProductsService) AllProducts() ([]products.Product, error) {
	return s.readModel.AllProducts()

}

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	PriceCent     int
	PriceCurrency string
}

func (s ProductsService) AddProduct(cms AddProductCommand) error {

	price, err := price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)
	if err != nil {
		return errors.Wrap(err, "invalid product price")
	}

	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)

	if err != nil {
		return errors.Wrap(err, "cannot create product")
	}

	if err := s.repo.save(p); err != nil {
		return errors.Wrap(err, "cannot saave product")
	}

	return nil
}
