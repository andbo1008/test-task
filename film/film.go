package film

import "time"

type Film struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"releaseDate"`
	createdAt   time.Time `json:"createdAt"`
}
