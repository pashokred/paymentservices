package app

import (
	"log"
	"net/http"
	"paymentservices/api/controllers"
)

func RunApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products/{id}", controllers.GetButtons)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
