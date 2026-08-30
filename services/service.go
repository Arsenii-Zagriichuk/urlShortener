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
	// define allowed chars for the new url
	const allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var newURL strings.Builder
	newURL.Grow(newURLLength) // optimize for memory

	// populate the new url picking a random char from the allowedChars
	for range newURLLength {
		randomIndex := rand.IntN(len(allowedChars))
		newURL.WriteByte(allowedChars[randomIndex])
	}
	// write the short url to the existing map
	shortCode := newURL.String()
	s.Model.Urls[shortCode] = url
	return shortCode
}

func (s *Service) GetOGUrl(newUrl string) (string, bool) {
	url, ok := s.Model.Urls[newUrl]
	if ok {
		return url, true
	}
	return "error has occurred", false
}
