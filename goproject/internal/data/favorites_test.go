package data

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestFavoriteModel_AddAndRemove(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:110755@localhost:5432/goproject?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM favorites")
	if err != nil {
		t.Fatal(err)
	}

	favoriteModel := FavoriteModel{DB: db}

	userID := 1
	songID := 111

	err = favoriteModel.Add(userID, songID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	songs, err := favoriteModel.GetAll(userID)
	if err != nil {
		t.Fatalf("expected no error on GetAll, got %v", err)
	}

	if len(songs) != 1 {
		t.Fatalf("expected 1 favorite, got %d", len(songs))
	}

	err = favoriteModel.Remove(userID, songID)
	if err != nil {
		t.Fatalf("expected no error on Remove, got %v", err)
	}
}
