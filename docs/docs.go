// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/actors": {
            "post": {
                "description": "Создает нового актера на основе полученных данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actors"
                ],
                "summary": "Создание актера",
                "parameters": [
                    {
                        "description": "Информация о новом актере",
                        "name": "actor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Actor"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Информация о созданном актере",
                        "schema": {
                            "$ref": "#/definitions/models.Actor"
                        }
                    },
                    "400": {
                        "description": "Неверное тело запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/actors/movies": {
            "get": {
                "description": "Получает список актеров с перечислением фильмов, в которых они снимались",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actors"
                ],
                "summary": "Получение списка актеров с фильмами",
                "responses": {
                    "200": {
                        "description": "Список актеров с фильмами",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ActorWithMovies"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/actors/{actorID}": {
            "put": {
                "description": "Обновляет информацию об актере по его идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actors"
                ],
                "summary": "Обновление информации об актере",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор актера",
                        "name": "actorID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Новая информация об актере",
                        "name": "actor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Actor"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Успешное обновление информации об актере"
                    },
                    "400": {
                        "description": "Неверный ID актера или неверное тело запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Актер не найден",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет актера по его идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actors"
                ],
                "summary": "Удаление актера",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор актера",
                        "name": "actorID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Успешное удаление актера"
                    },
                    "400": {
                        "description": "Неверный ID актера",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Актер не найден",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "description": "Получает список фильмов с возможностью поиска и сортировки",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Получение списка фильмов",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Поиск по названию фильма или имени актера",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Сортировка по названию (title), рейтингу (rating) или дате выпуска (release_date)",
                        "name": "sortBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список фильмов",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Movie"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавляет новый фильм на основе полученных данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Добавление фильма",
                "parameters": [
                    {
                        "description": "Информация о новом фильме",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Успешное добавление фильма"
                    },
                    "400": {
                        "description": "Неверное тело запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/movies/{movieID}": {
            "get": {
                "description": "Получает информацию о фильме по его идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Получение информации о фильме по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор фильма",
                        "name": "movieID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Информация о фильме",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    },
                    "400": {
                        "description": "Неверный ID фильма",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Фильм не найден",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновляет информацию о фильме по его идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Обновление информации о фильме",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор фильма",
                        "name": "movieID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Новая информация о фильме",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Успешное обновление информации о фильме"
                    },
                    "400": {
                        "description": "Неверный ID фильма или неверное тело запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Фильм не найден",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет фильм по его идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Удаление фильма",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор фильма",
                        "name": "movieID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Успешное удаление фильма"
                    },
                    "400": {
                        "description": "Неверный ID фильма",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Actor": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.ActorWithMovies": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "movies": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "release_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
