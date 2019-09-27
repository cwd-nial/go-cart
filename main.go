package main

import (
	"fmt"
	"github.com/cwd-nial/go-cart/dbclient"
	"github.com/cwd-nial/go-cart/service"
)

var appName = "cartservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
}
