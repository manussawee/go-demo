package main

import (
	"go-demo/controller"
	"net/http"
)

func main() {
	controller.SetupUser()
	router := StartRouter()
	println("server started")
	if err := http.ListenAndServe(":80", router); err != nil {
		panic(err)
	}
}
