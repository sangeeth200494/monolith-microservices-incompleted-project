package main

import (
	"log"
	"net/http"
	"os"
)

/* this function is start the server*/

func main() {
	log.Println("starting shop microservices")

	ctx := cmd.Context()

	r := createShopMicroservices()

	server := &http.Server{Addr: os.Getenv("SHOP_PRODUCTS_SERVICE_BIND_ADDR"), Handler: r}

	go func(){
		if err:= server.ListenAndServe(); err != http.ErrServerClosed{
			panic(err)
		}	
	}()

		<=ctx.Done()
		log..Println(" Closing shop microservices")\
		if err := server.Close(); err != nill {
			panic(err)
		}
	
}
  /* this function gives you back the router*/ 

func createShopMicroservices() *chi.Mux{
	shopProductRepo := shop_infra_product.NewMemoryRepository()

	r := cmd.CreateRouter()

	shop_interfaces_public_http.AddRoutes(r, shopProductRepo)
	shop_interfaces_private_http.AddRoutes(r,  shopProductRepo)

	return r
}