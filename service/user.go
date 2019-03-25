package service

import (
	"encoding/json"
	"fmt"
	"go-demo/controller"
	"io/ioutil"
	"net/http"
)

// Signup function
func Signup(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body request
	json.Unmarshal(b, &body)

	type response struct {
		UserID int    `json:"user_id"`
		Status string `json:"status"`
	}
	var data response
	data.UserID, data.Status = controller.Signup(body.Email, body.Password)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(jsonData)
}

// Signup function
func Login(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body request
	json.Unmarshal(b, &body)

	type response struct {
		UserID      int `json:"user_id"`
		VisitNumber int `json:"visit_number"`
	}
	var data response
	data.UserID, data.VisitNumber = controller.Login(body.Email, body.Password)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(jsonData)
}
