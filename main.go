package main

import (
	"fmt"
	"go-api/routes"
)

func main() {
	fmt.Println("Running server....")

	routes.InitializeRouter()

}
