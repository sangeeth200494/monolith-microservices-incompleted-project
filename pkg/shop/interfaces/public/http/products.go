package http

import (
	"net/http"

	"github.com/gin-gonic/gin/render"
)

func AddRoutes(router *chi.Mux, productReadModel productsReadmodel) {
	resource := productsResorce{productReadModel}
	router.Get("/products", resorce.GetAll)
}

type productReadModel interface {
	AllProducts() ([]products.Product, error)
}

type productResoce struct {
	readModel productsReadmodel
}

type productView struct {
	ID          string    `json:"id"`
	Name        string    `json:"name`
	Description string    `json:"description"`
	Price       priceView `json:"price"`
}

type priceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func privateViewFromPrice(p price.Price) priceView {
	return priceView{p.Cents(), p.Currency()}
}

func (p productResorces) GetAll(w http.ResponseWriter, r *http.Request) {

	products, err := p.readModel.AllProducts()
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}

	view := []productView{}
	for _, product := range products {
		view = append(view, productView{
			string(product.ID()),
			product.Name(),
			product.Description(),
			priceViewFromPrice(product.Price()),
		})
	}
	render.Respond(w, r, view)
}
