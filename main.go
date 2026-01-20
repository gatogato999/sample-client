package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		ShowMainToolUsage()
		os.Exit(1)
	}

	mainCommand := os.Args[1]
	var err error
	switch mainCommand {
	case "base":
		err = PingUrl()
	case "register":
		err = RegisterFlags(os.Args[2:])
	case "loggin":
		err = LogginUserFlags(os.Args[2:])
	case "list":
		err = ListUsers()
	case "query":
		err = SearchUserFlags(os.Args[2:])
	default:
		ShowMainToolUsage()
		os.Exit(1)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func LogginUserFlags(args []string) error {
	flagset := flag.NewFlagSet("login", flag.ExitOnError)
	email := flagset.String("email", "", "user valid email")
	password := flagset.String("password", "", "user password")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	return LogginUser(*email, *password)
}

func RegisterFlags(args []string) error {
	flagset := flag.NewFlagSet("register", flag.ExitOnError)
	fName := flagset.String("fName", "", "new user first name")
	lName := flagset.String("lName", "", "new user last name")
	email := flagset.String("email", "", "new user valid email ")
	password := flagset.String("password", "", "new user password (>8 characters)")
	phone := flagset.String("phone", "", "new user phone number")
	age := flagset.Int64("age", 0, "new user age (> 18) ")
	job := flagset.String("job", "", "new user job title ")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if *fName == "" || *lName == "" || *email == "" || *password == "" || *phone == "" ||
		*job == "" {
		return fmt.Errorf("flags can't be empty")
	}
	if *age < 18 {
		return fmt.Errorf("age under 18 is not allowed")
	}
	if len(*password) < 8 {
		return fmt.Errorf("short password , minimum 8 ch")
	}
	return Register(*fName, *lName, *email, *password, *phone, *job, *age)
}

func SearchUserFlags(args []string) error {
	flagset := flag.NewFlagSet("search", flag.ExitOnError)
	email := flagset.String("email", "", "user email to search for ")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if *email == "" {
		return fmt.Errorf("email flag is required")
	}
	return SearchUser(*email)
}
