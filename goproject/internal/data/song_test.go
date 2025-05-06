package data

import (
	"testing"

	"goproject/internal/validator"

	"github.com/stretchr/testify/assert"
)

func TestValidateSong(t *testing.T) {
	v := validator.New()

	song := &Song{
		Id:       1,
		Title:    "Test Song",
		Album_id: 100,
	}

	ValidateSong(v, song)
	assert.True(t, v.Valid(), "должен быть валидным, если все поля заполнены корректно")

	v = validator.New()
	song = &Song{
		Id:       2,
		Title:    "",
		Album_id: 100,
	}

	ValidateSong(v, song)
	assert.False(t, v.Valid(), "должен быть невалидным при пустом названии песни")
	assert.Contains(t, v.Errors, "title", "ошибка должна содержать ключ 'title'")

	v = validator.New()
	song = &Song{
		Id:       0,
		Title:    "Another Song",
		Album_id: 100,
	}

	ValidateSong(v, song)
	assert.False(t, v.Valid(), "должен быть невалидным при ID < 1")
	assert.Contains(t, v.Errors, "id", "ошибка должна содержать ключ 'id'")

	v = validator.New()
	song = &Song{
		Id:    3,
		Title: "Third Song",
	}

	ValidateSong(v, song)
	assert.False(t, v.Valid(), "должен быть невалидным при отсутствии albumId")
	assert.Contains(t, v.Errors, "albumId", "ошибка должна содержать ключ 'albumId'")
}
