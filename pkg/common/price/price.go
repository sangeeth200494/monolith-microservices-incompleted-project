package price

import "errors"

var (
	ErrPriceTooLow     = errors.New("price must be greater than 0")
	ErrInvalidCurrency = errors.New("Invalid currency")
)

type Price struct {
	Cents    uint
	Currency string
}

func NewPrice(cents uint, currency string) (Price, error) {

	if cents <= 0 {
		return Price{}, ErrPriceTooLow
	}

	if len(currency) != 3 {
		return Price{}, ErrInvalidCurrency
	}

	return Price{cents, currency}, nil
}

func NewPriceP(cents unit, currency string) Price {
	p, err := NewPrice(cents, currency)

	if err != nil {
		panic(err)
	}

	return p
}

func (p Price) Cents() unit {
	return p.cents
}

func (p Price) Currency() string {
	return p.currency
}
