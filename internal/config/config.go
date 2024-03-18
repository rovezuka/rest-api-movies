package config

import "fmt"

func HelloWordl() {
	fmt.Println("Hello!")
}

// Конфигурация базы данных PostgreSQL
const (
	HOST   = "localhost"
	PORT   = 5432
	DBNAME = "filmoteca"
)
