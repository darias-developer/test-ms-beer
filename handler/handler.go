package handler

import (
	"log"
	"net/http"
	"os"

	m "github.com/darias-developer/test-ms-beer/middleware"
	r "github.com/darias-developer/test-ms-beer/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* RouterManager maneja las rutas y puertos */
func RouterManager() {

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/beers", m.CheckDB(r.BeerFindAll)).Methods("GET")
	newRouter.HandleFunc("/beers", m.CheckDB(r.BeerAdd)).Methods("POST")
	newRouter.HandleFunc("/beers/{id}", m.CheckDB(r.BeerFindById)).Methods("GET")
	newRouter.HandleFunc("/beers/{id}/boxprice", m.CheckDB(r.BeerBoxPriceById)).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(newRouter)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
