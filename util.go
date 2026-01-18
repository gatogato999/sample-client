package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func ParseArgs(args []string) map[string]string {
	if len(args)%2 != 0 {
		fmt.Println("error : there is an arg without value")
		os.Exit(1)
	}
	inputs := make(map[string]string)
	for i := 0; i < len(args); i += 2 {
		inputs[args[i]] = args[i+1]
	}
	return inputs
}

func ShowMainToolUsage() {
	fmt.Println("\nsample-client base")
	fmt.Println("sample-client register fname lname email pass phone age job ")
	fmt.Println("sample-client loggin email <xxxx@xxx> password <xxx>")
	fmt.Println("sample-client list")
	fmt.Println("sample-client query <email> ")
	os.Exit(1)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetResponse(res *http.Response) {
	defer res.Body.Close()

	fmt.Println("\nResponse status:", res.Status)

	var data any
	err := json.NewDecoder(res.Body).Decode(&data)
	CheckError(err)

	b, err := json.MarshalIndent(data, "", "")
	CheckError(err)
	fmt.Println("\n", string(b))
}

func SaveJwt(cookies []*http.Cookie) error {
	for _, c := range cookies {
		if c.Name == "jwt_token" {
			err := os.MkdirAll(".secrets", 0o700)
			if err != nil {
				return err
			}
			return os.WriteFile(".secrets/jwt_token", []byte(c.Value), 0o600)
		}
	}
	return fmt.Errorf("cookie not found")
}

func Loadjwt() (*http.Cookie, error) {
	b, err := os.ReadFile(".secrets/jwt_token")
	if err != nil {
		return nil, err
	}
	return &http.Cookie{Name: "jwt_token", Value: string(b), Path: "/"}, nil
}

func AddJwt(req *http.Request) error {
	c, err := Loadjwt()
	if err != nil {
		return err
	}
	req.AddCookie(c)
	return nil
}
