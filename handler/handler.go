package handler

import (
	"log"
	"net/http"
	"os"

	m "github.com/darias-developer/test-ms-beer/middleware"
	r "github.com/darias-developer/test-ms-beer/router"
	u "github.com/darias-developer/test-ms-beer/util"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* RouterManager maneja las rutas y puertos */
func RouterManager() {

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/beers", r.BeerFindAll).Methods("GET")
	newRouter.HandleFunc("/beers", m.ValidateBeerAdd(r.BeerAdd)).Methods("POST")
	newRouter.HandleFunc("/beers/{id}", m.ValidateBeerFindById(r.BeerFindById)).Methods("GET")
	newRouter.HandleFunc("/beers/{id}/boxprice", m.ValidateBeerBoxPriceById(r.BeerBoxPriceById)).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		u.LogWarn.Printf("puerto no encontrado. se utilizara el puerto por defecto")
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(newRouter)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
