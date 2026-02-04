package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func PingUrl() error {
	url := "http://localhost:8080/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if err = GetResponse(res); err != nil {
		return err
	}

	return nil
}

func LogginUser(email, password string) error {
	url := "http://localhost:8080/auth"
	userInputs := map[string]any{
		"email":           email,
		"hashed_password": password,
	}
	reqBody, err := json.Marshal(userInputs)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	err = SaveJwt(res.Cookies())
	if err = GetResponse(res); err != nil {
		return err
	}

	return nil
}

func SearchUser(email string) error {
	url := fmt.Sprint("http://localhost:8080/query/", email)
	println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	err = AddJwt(req)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if err = GetResponse(res); err != nil {
		return err
	}

	return nil
}
