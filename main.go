package main

import (
	"github.com/rjribeiro/goBackend/database"
	"github.com/rjribeiro/goBackend/routes"
)

func main() {
	database.ConnectToDB()
	routes.HandleRequests()
}
