package transport

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rovezuka/rest-api-movies/internal/services"
)

// HandleFuncs регистрирует обработчики для всех эндпоинтов и возвращает роутер
func HandleFuncs() *mux.Router {
	router := mux.NewRouter()

	// Документация Swagger
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs"))))

	// Регистрация обработчиков с аутентификацией и авторизацией
	router.HandleFunc("/actors", services.CreateActor).Methods("POST")
	router.HandleFunc("/actors/{actorID}", services.PutActor).Methods("PUT")
	router.HandleFunc("/actors/{actorID}", services.DeleteActor).Methods("DELETE")

	router.HandleFunc("/movies-create", services.AddMovie).Methods("POST")
	router.HandleFunc("/actors-list", services.GetActorsWithMovies).Methods("GET")
	router.HandleFunc("/movies-list", services.GetMovies).Methods("GET")
	router.HandleFunc("/movies/{movieID}", services.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{movieID}", services.DeleteMovie).Methods("DELETE")

	return router
}
