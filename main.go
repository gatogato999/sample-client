package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// NOTE: base route /

	url := "http://localhost:8080/"
	err := pingUrl(url)
	checkError(err)
}

func pingUrl(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	checkError(err)
	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)

	var data map[string]any
	err = json.NewDecoder(res.Body).Decode(&data)
	checkError(err)

	b, err := json.MarshalIndent(data, "", "")
	checkError(err)
	fmt.Println(string(b))

	return nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
