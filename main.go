package main

import (
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		ShowMainToolUsage()
	}

	mainCommand := os.Args[1]
	println(mainCommand)
	switch mainCommand {
	case "base":
		PingUrl()
	case "register":
		Register(os.Args[2:])
	case "loggin":
		LogginUser(os.Args[2:])
	case "list":
		ListUsers()
	case "query":
		SearchUser(os.Args[2:])
	default:
		ShowMainToolUsage()
	}
}
