package data

import (
	"context"
	"database/sql"
	"time"
)

type FavoriteModel struct {
	DB *sql.DB
}

func (f FavoriteModel) Add(userID, songID int) error {
	query := `
		INSERT INTO favorites (user_id, song_id)
		VALUES ($1, $2)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := f.DB.ExecContext(ctx, query, userID, songID)
	return err
}

func (f FavoriteModel) Remove(userID, songID int) error {
	query := `
		DELETE FROM favorites
		WHERE user_id = $1 AND song_id = $2
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := f.DB.ExecContext(ctx, query, userID, songID)
	return err
}

func (f FavoriteModel) GetAll(userID int) ([]*Song, error) {
	query := `
		SELECT s.song_id, s.title, s.length, s.album_id
		FROM favorites
		INNER JOIN song s ON favorites.song_id = s.song_id
		WHERE favorites.user_id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := f.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*Song

	for rows.Next() {
		var song Song
		err := rows.Scan(&song.Id, &song.Title, &song.Length, &song.Album_id)
		if err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}
