package services

import (
	"math/rand/v2"
	"strings"
	"urlShortener/models"
)

type Service struct {
	Model *models.SavedURL
}

func (s *Service) GenerateURL(url string) string {
	const newURLLength = 6 // length of a new url
	var newURL strings.Builder
	newURL.Grow(newURLLength) // optimize for memory

	// populate the new url
	for range newURLLength {
		randomIndex := rand.IntN(len(url))
		newURL.WriteByte(url[randomIndex])
	}
	// write the short url to the existing map
	shortCode := newURL.String()
	s.Model.Urls[shortCode] = url
	return shortCode
}
