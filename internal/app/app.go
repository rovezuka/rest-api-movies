package app

import (
	"net/http"

	"github.com/rovezuka/rest-api-movies/internal/transport"
)

func RunServer(port string) {
	r := transport.HandleFuncs()
	http.ListenAndServe(port, r)
}
