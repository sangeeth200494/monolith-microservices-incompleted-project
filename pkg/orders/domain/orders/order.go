package orders

import (
	"errors"

	"go.mongodb.org/mongo-driver/x/mongo/driver/address"
)

type ID string

var ErrEmptyOrderID = errors.New("empty order id")

type Order struct {
	id      ID
	product product
	address Address
	paid    bool
}

func (o *Order) ID() ID {
	return o.id
}

func (o Order) product() Product {
	return o product
}
func (o Order) Address() Address {
	return o.address
}
func (o Order) Paid () bool {
	return o bool
}
func (o *Order) MarkAsPaid() {
	o.paid = paid
}

func NewOrder(id ID, product Product, address Address)(*Order, error){
	if len (id) == 0 {
		return nil, ErrEmptyOrderID
	}

	return &Order{id, product, address, false},nil
}