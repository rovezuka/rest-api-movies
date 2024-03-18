package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/rovezuka/rest-api-movies/internal/database"
	"github.com/rovezuka/rest-api-movies/internal/models"
)

// @Summary Создание актера
// @Description Создает нового актера на основе полученных данных
// @Tags actors
// @Accept json
// @Produce json
// @Param actor body models.Actor true "Информация о новом актере"
// @Success 201 {object} models.Actor "Информация о созданном актере"
// @Failure 400 {string} string "Неверное тело запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /actors [post]
func CreateActor(w http.ResponseWriter, r *http.Request) {
	// Обработка тела запроса
	var actor models.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		log.Println("не удалось создать нового актера")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вставка информации об актере в базу данных
	db, err := database.ConnectToDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	err = db.QueryRow("INSERT INTO actors (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id", actor.Name, actor.Gender, actor.BirthDate).Scan(&actor.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actor)
}

// @Summary Обновление информации об актере
// @Description Обновляет информацию об актере по его идентификатору
// @Tags actors
// @Accept json
// @Produce json
// @Param actorID path int true "Идентификатор актера"
// @Param actor body models.Actor true "Новая информация об актере"
// @Success 204 "Успешное обновление информации об актере"
// @Failure 400 {string} string "Неверный ID актера или неверное тело запроса"
// @Failure 404 {string} string "Актер не найден"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /actors/{actorID} [put]
func PutActor(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр {actorID} из URL
	vars := mux.Vars(r)
	actorID := vars["actorID"]

	// Парсим actorID в int
	id, err := strconv.Atoi(actorID)
	if err != nil {
		log.Println("неверный ID актера")
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	// Проверяем, что запрос содержит тело с новой информацией об актере
	var newActor models.Actor
	err = json.NewDecoder(r.Body).Decode(&newActor)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectToDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// Проверяем существование актера с указанным ID
	var existingActorID int
	err = db.QueryRow("SELECT id FROM actors WHERE id = $1", id).Scan(&existingActorID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("актер не найден при обновлении информации")
			http.Error(w, "Actor not found", http.StatusNotFound)
		} else {
			log.Println("ошибка при поиске актера")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Обновляем информацию об актере в базе данных
	_, err = db.Exec("UPDATE actors SET name = $1, gender = $2, birth_date = $3 WHERE id = $4",
		newActor.Name, newActor.Gender, newActor.BirthDate, id)
	if err != nil {
		log.Println("ошибка при обновлении информации актера")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusNoContent)
}

// @Summary Удаление актера
// @Description Удаляет актера по его идентификатору
// @Tags actors
// @Accept json
// @Produce json
// @Param actorID path int true "Идентификатор актера"
// @Success 204 "Успешное удаление актера"
// @Failure 400 {string} string "Неверный ID актера"
// @Failure 404 {string} string "Актер не найден"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /actors/{actorID} [delete]
func DeleteActor(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметра {actorID} из URL
	vars := mux.Vars(r)
	actorID := vars["actorID"]

	// Парсинг actorID в int
	id, err := strconv.Atoi(actorID)
	if err != nil {
		log.Println("неверный ID актера при удалении")
		http.Error(w, "Неверный ID актера", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectToDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// Проверка существования актера с указанным ID
	var existingActorID int
	err = db.QueryRow("SELECT id FROM actors WHERE id = $1", id).Scan(&existingActorID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("актер не найден при удалении")
			http.Error(w, "Актер не найден", http.StatusNotFound)
		} else {
			log.Println("произошла ошибка при удалении актера")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Удаление информации об актере из базы данных
	_, err = db.Exec("DELETE FROM actors WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusNoContent)
}

// @Summary Добавление фильма
// @Description Добавляет новый фильм на основе полученных данных
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body models.Movie true "Информация о новом фильме"
// @Success 201 "Успешное добавление фильма"
// @Failure 400 {string} string "Неверное тело запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movies [post]
func AddMovie(w http.ResponseWriter, r *http.Request) {
	// Парсинг данных о фильме из тела запроса
	var newMovie models.Movie
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		log.Println("ошибка при добавлении фильма")
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectToDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	actorsString := "{" + newMovie.Actors + "}"

	// Вставка новой информации о фильме в базу данных
	_, err = db.Exec("INSERT INTO movies (title, description, release_date, rating, actors) VALUES ($1, $2, $3, $4, $5)",
		newMovie.Title, newMovie.Description, newMovie.ReleaseDate, newMovie.Rating, actorsString)
	if err != nil {
		log.Println("ошибка при вставке фильма в базу данных")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusCreated)
}

// @Summary Получение списка фильмов
// @Description Получает список фильмов с возможностью поиска и сортировки
// @Tags movies
// @Accept json
// @Produce json
// @Param search query string false "Поиск по названию фильма или имени актера"
// @Param sortBy query string false "Сортировка по названию (title), рейтингу (rating) или дате выпуска (release_date)"
// @Success 200 {array} models.Movie "Список фильмов"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movies [get]
func GetMovies(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры запроса для поиска и сортировки
	searchQuery := r.URL.Query().Get("search")
	sortBy := r.URL.Query().Get("sortBy")

	// Подготовка SQL-запроса для получения списка фильмов с возможностью поиска
	query := "SELECT m.* FROM movies m"

	if searchQuery != "" {
		query += fmt.Sprintf(" JOIN movie_actors ma ON m.id = ma.movie_id "+
			"JOIN actors a ON a.id = ma.actor_id WHERE m.title ILIKE '%%%s%%' OR a.name ILIKE '%%%s%%'", searchQuery, searchQuery)
	}

	// Добавление условия сортировки в SQL-запрос
	if sortBy != "" {
		query += " ORDER BY " + sortBy
	}

	db, err := database.ConnectToDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// Выполнение SQL-запроса
	rows, err := db.Query(query)
	if err != nil {
		log.Println("ошибка при получении фильма")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Слайс для хранения фильмов
	var movies []models.Movie

	// Обработка результатов запроса
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating, &movie.Actors)
		if err != nil {
			log.Println("ошибка про получении фильма из базы данных")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}

	// Преобразование списка фильмов в JSON и отправка ответа
	json.NewEncoder(w).Encode(movies)
}

// @Summary Обновление информации о фильме
// @Description Обновляет информацию о фильме по его идентификатору
// @Tags movies
// @Accept json
// @Produce json
// @Param movieID path int true "Идентификатор фильма"
// @Param movie body models.Movie true "Новая информация о фильме"
// @Success 204 "Успешное обновление информации о фильме"
// @Failure 400 {string} string "Неверный ID фильма или неверное тело запроса"
// @Failure 404 {string} string "Фильм не найден"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movies/{movieID} [put]
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметра {movieID} из URL
	vars := mux.Vars(r)
	movieID := vars["movieID"]

	// Парсинг movieID в int
	id, err := strconv.Atoi(movieID)
	if err != nil {
		log.Println("неверный ID фильма при обновлении информации")
		http.Error(w, "Неверный ID фильма", http.StatusBadRequest)
		return
	}

	// Парсинг данных о фильме из тела запроса
	var updatedMovie models.Movie
	err = json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectToDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// Проверка существования фильма с указанным ID
	var existingMovieID int
	err = db.QueryRow("SELECT id FROM movies WHERE id = $1", id).Scan(&existingMovieID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("фильм не найден")
			http.Error(w, "Фильм не найден", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	actorsString := "{" + updatedMovie.Actors + "}"

	// Обновление информации о фильме в базе данных
	_, err = db.Exec("UPDATE movies SET title = $1, description = $2, release_date = $3, rating = $4, actors = $5 WHERE id = $6",
		updatedMovie.Title, updatedMovie.Description, updatedMovie.ReleaseDate, updatedMovie.Rating, actorsString, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusNoContent)
}

// @Summary Удаление фильма
// @Description Удаляет фильм по его идентификатору
// @Tags movies
// @Accept json
// @Produce json
// @Param movieID path int true "Идентификатор фильма"
// @Success 204 "Успешное удаление фильма"
// @Failure 400 {string} string "Неверный ID фильма"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movies/{movieID} [delete]
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметра {movieID} из URL
	vars := mux.Vars(r)
	movieID := vars["movieID"]

	// Парсинг movieID в int
	id, err := strconv.Atoi(movieID)
	if err != nil {
		log.Println("неверный ID фильма при удалении")
		http.Error(w, "Неверный ID фильма", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectToDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// Удаление информации о фильме из базы данных
	_, err = db.Exec("DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusNoContent)
}

// @Summary Получение информации о фильме по ID
// @Description Получает информацию о фильме по его идентификатору
// @Tags movies
// @Accept json
// @Produce json
// @Param movieID path int true "Идентификатор фильма"
// @Success 200 {object} models.Movie "Информация о фильме"
// @Failure 400 {string} string "Неверный ID фильма"
// @Failure 404 {string} string "Фильм не найден"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /movies/{movieID} [get]
func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметра {movieID} из URL
	vars := mux.Vars(r)
	movieID := vars["movieID"]

	// Парсинг movieID в int
	id, err := strconv.Atoi(movieID)
	if err != nil {
		log.Println("неверный ID фильма при получении по ID")
		http.Error(w, "Неверный ID фильма", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectToDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// Запрос информации о фильме из базы данных
	var movie models.Movie
	err = db.QueryRow("SELECT id, title, description, release_date, rating FROM movies WHERE id = $1", id).Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Фильм не найден", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Конвертация информации о фильме в формат JSON и отправка клиенту
	json.NewEncoder(w).Encode(movie)
}

// @Summary Получение списка актеров с фильмами
// @Description Получает список актеров с перечислением фильмов, в которых они снимались
// @Tags actors
// @Accept json
// @Produce json
// @Success 200 {array} models.ActorWithMovies "Список актеров с фильмами"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /actors/movies [get]
func GetActorsWithMovies(w http.ResponseWriter, r *http.Request) {
	// Подготовка SQL-запроса для получения списка актеров с их участием в фильмах
	query := `
		SELECT a.id, a.name, a.gender, a.birth_date, array_agg(m.title ORDER BY m.title) AS movies
		FROM actors a
		JOIN movie_actors ma ON a.id = ma.actor_id
		JOIN movies m ON ma.movie_id = m.id
		GROUP BY a.id, a.name
	`

	db, err := database.ConnectToDB(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// Выполнение SQL-запроса
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Слайс для хранения актеров с их участием в фильмах
	var actors []models.ActorWithMovies

	// Обработка результатов запроса
	for rows.Next() {
		var actor models.ActorWithMovies
		err := rows.Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.BirthDate, pq.Array(&actor.Movies))
		if err != nil {
			log.Println("ошибка про поиске актера в базе данных")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		actors = append(actors, actor)
	}

	// Преобразование списка актеров в JSON и отправка ответа
	json.NewEncoder(w).Encode(actors)
}
