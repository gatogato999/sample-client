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

func SearchUserFlags(args []string) error {
	flagset := flag.NewFlagSet("search", flag.ExitOnError)
	email := flagset.String("email", "", "user email ")
	password := flagset.String("password", "", "user password")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if *email == "" {
		return fmt.Errorf("email flag is required")
	}
	return SearchUser(*email, *password)
}
