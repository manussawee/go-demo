package main

import (
	"go-demo/service"
	"net/http"

	"github.com/gorilla/mux"
)

func withConfig(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With")
		f(res, req)
	}
}

// StartRouter function
func StartRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/signup", withConfig(service.Signup)).Methods("POST")
	router.HandleFunc("/login", withConfig(service.Login)).Methods("POST")
	return router
}
