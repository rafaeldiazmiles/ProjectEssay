package main

import (
	"fmt"
	"log"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

// Responsible for initializing and starting our gRPC server
func Run() error {
	fmt.Println("User service starting...")
	return nil
}
