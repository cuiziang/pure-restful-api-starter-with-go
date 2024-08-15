package main

import "github.com/cuiziang/pure-restFul-api-starter-with-go/internal/server"

func main() {
	err := server.NewServer().SetupRoutes().Start()
	if err != nil {
		panic(err)
	}
}
