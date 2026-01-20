package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func ShowMainToolUsage() {
	fmt.Println("\nsample-client base")
	fmt.Println("sample-client register fname lname email pass phone age job ")
	fmt.Println("sample-client loggin email <xxxx@xxx> password <xxx>")
	fmt.Println("sample-client list")
	fmt.Println("sample-client query <email> ")
}

func GetResponse(res *http.Response) error {
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		b, _ := io.ReadAll(res.Body)
		return fmt.Errorf("server error: %s", b)
	} else {
		fmt.Println("\nResponse status:", res.Status)

		var data any
		err := json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			return err
		}

		b, err := json.MarshalIndent(data, "", "")
		if err != nil {
			return err
		}
		fmt.Println("\n", string(b))
		return nil
	}
}

func getConfigDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "sample-client"), nil
}

func SaveJwt(cookies []*http.Cookie) error {
	dir, err := getConfigDir()
	if err != nil {
		return err
	}
	for _, c := range cookies {
		if c.Name == "jwt_token" {
			if err := os.MkdirAll(dir, 0o700); err != nil {
				return err
			}
			return os.WriteFile(filepath.Join(dir, "jwt_token"), []byte(c.Value), 0o600)
		}
	}
	return fmt.Errorf("cookie not found")
}

func Loadjwt() (*http.Cookie, error) {
	dir, err := getConfigDir()
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(filepath.Join(dir, "jwt_token"))
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
