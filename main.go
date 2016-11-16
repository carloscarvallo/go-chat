package main

import (
	"flag"
	"fmt"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		fmt.Println("is host")
	} else {
		// go run main.go <ip>
		fmt.Println("is guest")
	}
}
