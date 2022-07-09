package repository

import (
	"rest-api/film"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PostgresDB struct {
	db *sqlx.DB
}

func NewFilmPostgres(db *sqlx.DB) *PostgresDB {
	return &PostgresDB{
		db: db,
	}
}

func (p *PostgresDB) CreateFilm(f film.Film) (int, error) {
	logrus.Print(f)
	var id int
	creatAt := time.Now().UTC()

	query := "INSERT INTO films (title, releaseDate, createdAt) values ($1, $2, $3) RETURNING id;"
	row := p.db.QueryRow(query, f.Title, f.ReleaseDate, creatAt)
	if err := row.Scan(&id); err != nil {
		logrus.Fatal(err)
		return 0, err
	}
	return id, nil
}

func (p *PostgresDB) GetById(id string) (film.Film, error) {
	var f film.Film
	query := "SELECT id, title, releaseDate FROM films WHERE id = $1"
	row := p.db.QueryRow(query, id)
	if err := row.Scan(&f.Id, &f.Title, &f.ReleaseDate); err != nil {
		logrus.Error(err)
	}
	f.ReleaseDate.Format(time.RFC3339)
	return f, nil
}

func (p *PostgresDB) GetAll() ([]film.Film, error) {
	var f film.Film
	var arrF []film.Film

	query := "SELECT id, title, releaseDate FROM films"
	row, err := p.db.Query(query)
	if err != nil {
		logrus.Fatal(err)
	}
	for row.Next() {
		if err := row.Scan(&f.Id, &f.Title, &f.ReleaseDate); err != nil {
			logrus.Error(err)
		}
		f.ReleaseDate.Format(time.RFC3339)
		logrus.Print(f.ReleaseDate)
		arrF = append(arrF, f)

	}
	return arrF, nil
}
