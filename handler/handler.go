package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/darias-developer/test-ms-beer/middleware"
	"github.com/darias-developer/test-ms-beer/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* RouterManager maneja las rutas y puertos */
func RouterManager() {

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/beers", middleware.CheckDB(router.SearchBeers)).Methods("GET")
	newRouter.HandleFunc("/beers", middleware.CheckDB(router.AddBeers)).Methods("POST")
	newRouter.HandleFunc("/beers/{id}", middleware.CheckDB(router.SearchBeerById)).Methods("GET")
	newRouter.HandleFunc("/beers/{id}/boxprice", middleware.CheckDB(router.BoxBeerPriceById)).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(newRouter)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
