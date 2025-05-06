package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Songs     interface {
		Insert(song *Song) error
		Get(id int64) (*Song, error)
		GetAll(title string, length int, filters Filters) ([]*Song, Metadata, error)
		Update(song *Song) error
		Delete(id int64) error
	}
	Favorites interface {
		Add(userID, songID int) error
		Remove(userID, songID int) error
		GetAll(userID int) ([]*Song, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Songs:     SongModel{DB: db},
		Favorites: FavoriteModel{DB: db},
	}
}
