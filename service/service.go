package service

import (
	"rest-api/film"
	"rest-api/repository"
)

type Service struct {
	ServiceFilm
}

type ServiceFilm interface {
	CreateFilm(film film.Film) (int, error)
	GetById(id string) (film.Film, error)
	GetAll() ([]film.Film, error)
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		ServiceFilm: NewFilmService(repository.Films),
	}
}
