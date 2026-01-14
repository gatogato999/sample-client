package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// NOTE: base route /
	var err error
	var url string

	url = "http://localhost:8080/"
	err = pingUrl(url)
	checkError(err)

	// NOTE: get users route /users
	url = "http://localhost:8080/users"
	err = getAllUsers(url, "mo@gmail.com")
	checkError(err)
}

func getAllUsers(url, email string) error {
	e := map[string]string{"email": email}

	jsonE, err := json.Marshal(e)
	checkError(err)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonE))
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
