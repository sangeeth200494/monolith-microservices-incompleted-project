package application

type productsService interface {
}

type paymentsService interface {
}

type OrdersService struct {
}

func NewOrdersService() {

}

type placeOrderCommand struct {
}

func (s OrdersService) PlaceOrder(cmd placeOrderCommand) error {

}

type MarkOrderAsPaidCommand struct {
}

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {

}

func (s OrdersService) OrderByID(id orders.ID) (orders.Order, error) {

}
