package controller

import (
	"sync"
	"time"
)

type userType struct {
	id       int
	email    string
	password string
}

type userSingleton struct {
	userMap    map[string]userType
	userAmount int
	mux        sync.Mutex
}

type visitSingleton struct {
	count int
	mux   sync.Mutex
}

var userSingle userSingleton
var visitSingle visitSingleton

func getVisitNumber(c chan int) {
	visitSingle.mux.Lock()
	defer visitSingle.mux.Unlock()
	visitSingle.count++

	time.Sleep(time.Second * 5)
	c <- visitSingle.count
}

func getUser(email string) (userType, bool) {
	user, isFound := userSingle.userMap[email]

	time.Sleep(time.Second * 3)
	return user, isFound
}

// SetupUser function
func SetupUser() {
	userSingle.userMap = make(map[string]userType)
}

// Signup function
func Signup(email string, password string) (int, string) {
	userSingle.mux.Lock()
	defer userSingle.mux.Unlock()

	if _, isFound := getUser(email); !isFound {
		userSingle.userAmount++
		userSingle.userMap[email] = userType{
			id:       userSingle.userAmount,
			email:    email,
			password: password,
		}
		return userSingle.userAmount, "OK"
	}
	return 0, "USER EXISTS"
}

// Login function
func Login(email string, password string) (int, int) {
	c := make(chan int)
	go getVisitNumber(c)
	var userID int
	if user, isFound := getUser(email); isFound && user.password == password {
		userID = user.id
	}
	return userID, <-c
}
