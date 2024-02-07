package main

import (
	"belajar-gin/database"
	"belajar-gin/routers"
)

func main() {
	var PORT = ":8080"

	database.StartDB()
	routers.StartServer().Run(PORT)
}
