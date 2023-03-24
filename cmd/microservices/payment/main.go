package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Starting the payments microservice")
	
	defer log.Println("closing payments microservices")

	ctx := cmd.Context()

	paymentsInterface := createPaymentsMicroservice()
	if err:= paymentsInterface.Run(ctx); err !=nill {
		panic(err)
	}
}

func createPaymentsMicroservice() amqp.paymentsInterface{
	cmd.WaitForService(os.Getenv("SHOP_RABITMQ_ADDR"))

	paymentsService := payments_app.NewPaymentsService(
		payments_infra_orders.NewHTTPClint(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)

	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR"))
		os.Getenv("SHOP_RABITMQ_ORDERS_TO_PAY_QUEUE")
		paymentsService,
	)

	if err!= nil{
		panic(err)

		return payments 
	}
}