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

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	w.WriteHeader(http.StatusNoContent)
}

// StartRouter function
func StartRouter() *mux.Router {
	router := mux.NewRouter()
	// router.Methods("OPTIONS").HandlerFunc(preflightHandler)

	router.HandleFunc("/signup", withConfig(service.Signup)).Methods("POST")
	router.HandleFunc("/login", withConfig(service.Login)).Methods("POST")
	return router
}
