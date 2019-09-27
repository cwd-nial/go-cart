package main

import (
	"fmt"
	"github.com/cwd-nial/go-cart/service"
)

var appName = "cartservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}
