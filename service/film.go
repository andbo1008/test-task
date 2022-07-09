package service

import (
	"rest-api/film"
	"rest-api/repository"

	"github.com/sirupsen/logrus"
)

type FilmServise struct {
	repository repository.Films
}

func NewFilmService(repository repository.Films) *FilmServise {
	return &FilmServise{repository: repository}
}

func (s *FilmServise) CreateFilm(f film.Film) (int, error) {
	logrus.Print("service :", f)
	return s.repository.CreateFilm(f)
}
func (s *FilmServise) GetById(id string) (film.Film, error) {
	return s.repository.GetById(id)
}
func (s *FilmServise) GetAll() ([]film.Film, error) {
	return s.repository.GetAll()
}
