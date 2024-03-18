package main

import (
	_ "github.com/lib/pq"
	"github.com/rovezuka/rest-api-movies/internal/app"
)

func main() {
	app.RunServer(":8080")
}
