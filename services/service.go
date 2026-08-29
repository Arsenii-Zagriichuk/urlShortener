package services

import (
	"math/rand/v2"
	"strings"
	"urlShortener/models"
)

type Service struct {
	Model *models.SavedURL
}

func (s *Service) GenerateURL(url string) models.SavedURL {
	const newURLLength = 6 // length of a new url
	var newURL strings.Builder
	newURL.Grow(newURLLength) // optimize for memory

	// populate the new url
	for range newURLLength {
		randomIndex := rand.IntN(len(url))
		newURL.WriteByte(url[randomIndex])
	}
	// put a new map into a struct
	m := make(map[string]string)
	m[newURL.String()] = url
	return models.SavedURL{
		Urls: m,
	}
}
