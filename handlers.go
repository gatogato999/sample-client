package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func PingUrl() error {
	url := "http://localhost:8080/"
	req, err := http.NewRequest("GET", url, nil)
	CheckError(err)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	CheckError(err)

	GetResponse(res)

	return nil
}

func Register(args []string) error {
	input := ParseArgs(args)
	url := "http://localhost:8080/register"
	age, err := strconv.Atoi(input["age"])
	CheckError(err)
	userInputs := map[string]any{
		"firstName":       input["firstName"],
		"lastName":        input["lastName"],
		"email":           input["email"],
		"hashed_password": input["password"],
		"phone":           input["phone"],
		"age":             age,
		"job":             input["job"],
	}
	reqBody, err := json.Marshal(userInputs)

	CheckError(err)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	CheckError(err)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	CheckError(err)

	GetResponse(res)

	return nil
}

func LogginUser(args []string) error {
	url := "http://localhost:8080/auth"
	flags := ParseArgs(args)
	userInputs := map[string]any{
		"email":           flags["email"],
		"hashed_password": flags["password"],
	}
	reqBody, err := json.Marshal(userInputs)
	CheckError(err)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	CheckError(err)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	CheckError(err)

	SaveJwt(res.Cookies())
	GetResponse(res)

	return nil
}

func ListUsers() error {
	url := "http://localhost:8080/users"

	req, err := http.NewRequest("GET", url, nil)
	CheckError(err)
	AddJwt(req)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	CheckError(err)

	GetResponse(res)

	return nil
}

func SearchUser(args []string) error {
	input := ParseArgs(args)
	url := fmt.Sprint("http://localhost:8080/query/", input["email"])
	println(url)
	req, err := http.NewRequest("POST", url, nil)
	CheckError(err)

	AddJwt(req)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	CheckError(err)

	GetResponse(res)

	return nil
}
