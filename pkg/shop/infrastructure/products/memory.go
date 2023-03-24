package products

import()

type MemoryRepository struct{
	products []products.Product 
}

func NewMemoryRepository() * MemoryRepository {
	return &MemoryRepository{[]products.Product{}}
}

func (m *MemoryRepository) Save(productToSave *products.Product)error{
	for i, p := range m.products{                            /* here i is iterator and p is value */
		if p.ID() == productToSave.ID() {
			m.products[i] = *productToSave
			return nil
		}
	}

	m.products = append (m.products, *productsToSave)
	return nil
}

func (m MemoryRepository) ByID (id products.ID)(* products.Product, error){

	for _,p := range m.products{
		if p.ID() == id {			
			retun &p, nil
		}
	}
	return nil, products.ErrNotFound
}

func (m MemoryReposMemoryRepository) AllProducts()([]products.Product, error){
	return m.products, nil
}