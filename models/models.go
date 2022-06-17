package models

import (
	"database/sql"
	"time"
)

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Movie struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Year        int            `json:"year"`
	ReleaseDate time.Time      `json:"release_date"`
	Runtime     int            `json:"runtime"`
	Rating      int            `json:"rating"`
	MPAARating  string         `json:"mpaa_rating"`
	CreateAt    time.Time      `json:"-"`
	UpdateAt    time.Time      `json:"-"`
	MovieGenre  map[int]string `json:"genres"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreateAt  time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`
}

type MovieGenre struct {
	ID       int       `json:"-"`
	MovieID  int       `json:"-"`
	GenreID  int       `json:"-"`
	Genre    Genre     `json:"genre"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}
