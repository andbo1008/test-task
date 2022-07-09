package repository

import (
	"rest-api/film"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Films
}

type Films interface {
	CreateFilm(f film.Film) (int, error)
	GetById(id string) (film.Film, error)
	GetAll() ([]film.Film, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Films: NewFilmPostgres(db),
	}
}
