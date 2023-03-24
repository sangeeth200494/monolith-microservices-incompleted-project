package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/* creating order microservices and return the order router handler*/
func main() {
	log.Println("Starting the orders microservices")

	ctx := cmd.context()

	r , closeFn := createOrderMicroservices() /* main function exept func main* , creating microservices */

	defer closeFn()

	server := &http.server{Addr: os.Getenv("SHOP_ORDER_SERVICE_BIND_ADDR"), Handler: r}

	go func ()  {
		if err := server.ListenAndServe(); err!= http.ErrServerClosed{
			panic(err)
		}
	}()

		<-ctx.Done()

		log.Println("closing order microservice")

		if err := server.Close(); err != nill{

			panic(err)
		}
	}
}

func createOrderMicroservices()(router *chi.mux, closeFn func()){
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	shopHTTPClint := orders_infra_produc.NewHTTPClint(OS.Getenv("SHOP_PRODUCT_SERVICE_ADDR"))

	r := cmd.CreateRouter()

	orders_public_hrrp.AddRouts(r, ordersService, ordersRepo)
	orders_private_htTP.AddRouts(r, ordersService, ordersRepo)  

	return r, func () {
		
	}
}