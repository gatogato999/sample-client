package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"hashed_password"`
	Phone     string `json:"phone"`
	Age       int    `json:"age"`
	Job       string `json:"job"`
}

func main() {
	var err error
	var url string
	var jwtToken string

	// NOTE: base route /
	url = "http://localhost:8080/"
	err = pingUrl(url)
	checkError(err)

	// NOTE: insert a user to the database
	url = "http://localhost:8080/register"
	user := User{
		FirstName: "user1",
		LastName:  "user1",
		Email:     "user1@gmail.com",
		Password:  "passpass",
		Phone:     "123456789",
		Age:       29,
		Job:       "user1job",
	}
	err = insertUser(url, user)
	checkError(err)

	// NOTE: loggin : /auth
	url = "http://localhost:8080/auth"
	err = logginUser(url, user.Email, user.Password)
	checkError(err)

	// NOTE: get users route /users

	url = "http://localhost:8080/users"
	jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQGdtYWlsLmNvbSIsImV4cCI6MTc2ODUwNzA2NywiaWF0IjoxNzY4NTA2MTY3LCJzdWIiOiJ1c2VyMUBnbWFpbC5jb20ifQ.EkMrOQkpttbel7kj-pVBM7d6T8LIEf8FLIlP2X8oBmw"
	err = getAllUsers(url, "mo@gmail.com", jwtToken)
	checkError(err)

	// NOTE: get search for user : /query
	url = "http://localhost:8080/query"
	jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQGdtYWlsLmNvbSIsImV4cCI6MTc2ODUwNzA2NywiaWF0IjoxNzY4NTA2MTY3LCJzdWIiOiJ1c2VyMUBnbWFpbC5jb20ifQ.EkMrOQkpttbel7kj-pVBM7d6T8LIEf8FLIlP2X8oBmw"
	err = searchUserByEmail(url, "mo@gmail.com", jwtToken)
	checkError(err)
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

func getAllUsers(url, email, jwtToken string) error {
	userInputs := map[string]string{"email": email}
	token := "Bearer " + jwtToken

	reqBody, err := json.Marshal(userInputs)
	checkError(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	checkError(err)
	req.Header.Add("Authorization", token)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	checkError(err)

	getResponse(res)

	return nil
}

func searchUserByEmail(url, email, jwtToken string) error {
	token := "Bearer " + jwtToken

	url = url + fmt.Sprint("/", email)
	req, err := http.NewRequest("POST", url, nil)
	checkError(err)

	req.Header.Add("Authorization", token)
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
		os.Exit(1)
	}
}

func getResponse(res *http.Response) {
	defer res.Body.Close()

	fmt.Println("\n-------------------------\nResponse status:", res.Status)

	var data any
	err := json.NewDecoder(res.Body).Decode(&data)
	checkError(err)

	b, err := json.Marshal(data)
	checkError(err)
	fmt.Println("\n", string(b))
}
