package models

// ActorWithMovies представляет информацию об актере с его фильмами
type ActorWithMovies struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Gender    string   `json:"gender"`
	BirthDate string   `json:"birth_date"`
	Movies    []string `json:"movies"`
}

// Actor структура для актера
type Actor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
}

// Movie структура для фильма
type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Rating      float64 `json:"rating"`
	Actors      string  `json:"actors"`
}
