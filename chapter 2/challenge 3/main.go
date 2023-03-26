package main

import (
	"chapter2-challenge3/database"
	routers "chapter2-challenge3/routes"

	_ "github.com/lib/pq"
)

func main() {

	database.DatabaseConnection()
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
