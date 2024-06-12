package main

import (
	"fmt"

	"github.com/vashu992/MVC-Pattern/api"
	"github.com/vashu992/MVC-Pattern/controller"
)

func main() {

	api := api.ApiRoutes{}

	api.StartApp(controller.Server{})
	fmt.Printf("main server = %v\n", api)
}
