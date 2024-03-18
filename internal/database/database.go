package database

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/rovezuka/rest-api-movies/internal/config"
)

func ConnectToDB(r *http.Request) (*sql.DB, error) {
	username, password := r.Header.Get("username"), r.Header.Get("password")
	db, err := sql.Open("postgres", "postgres://"+username+":"+password+"@"+config.HOST+":"+strconv.Itoa(config.PORT)+"/"+config.DBNAME)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
