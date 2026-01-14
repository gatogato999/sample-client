package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID        int64  ` json:"id"`
	FirstName string ` json:"firstName"`
	LastName  string ` json:"lastName"`
	Email     string ` json:"email"`
	Password  string ` json:"password"`
	Phone     string ` json:"phone"`
	Age       int    ` json:"age"`
	Job       string ` json:"job"`
}

func main() {
	var err error
	var url string

	// NOTE: base route /
	// url = "http://localhost:8080/"
	// err = pingUrl(url)
	// checkError(err)

	// NOTE: get users route /users
	// url = "http://localhost:8080/users"
	// err = getAllUsers(url, "mo@gmail.com")
	// checkError(err)

	// NOTE:  insert a user to the database
	// url = "http://localhost:8080/register"
	user := User{
		FirstName: "ftestData",
		LastName:  "ltestData",
		Email:     "momto@gmal.com",
		Password:  "pass",
		Phone:     "12345678",
		Age:       89,
		Job:       "testData",
	}
	// err = insertUser(url, user)
	// checkError(err)

	// NOTE: loggin
	url = "http://localhost:8080/auth"
	err = logginUser(url, user.Email, user.Password)
	checkError(err)
}

func logginUser(url, email, password string) error {
	userInputs := map[string]string{"email": email, "hashed_password": password}
	reqBody, err := json.Marshal(userInputs)
	checkError(err)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	checkError(err)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	checkError(err)

	getResponse(res)

	return nil
}

func insertUser(url string, user User) error {
	reqBody, err := json.Marshal(user)
	checkError(err)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	checkError(err)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	checkError(err)

	getResponse(res)

	return nil
}

func getAllUsers(url, email string) error {
	userInputs := map[string]string{"email": email}

	reqBody, err := json.Marshal(userInputs)
	checkError(err)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(reqBody))
	checkError(err)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	checkError(err)

	getResponse(res)

	return nil
}

func pingUrl(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	checkError(err)

	getResponse(res)

	return nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getResponse(res *http.Response) {
	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)

	var data any
	err := json.NewDecoder(res.Body).Decode(&data)
	checkError(err)

	b, err := json.MarshalIndent(data, "", "")
	checkError(err)
	fmt.Println("\n", string(b))
}
