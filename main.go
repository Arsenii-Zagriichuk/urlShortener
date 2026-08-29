package main

import (
	"urlShortener/handlers"
	"urlShortener/models"
	"urlShortener/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	model := &models.SavedURL{}
	service := &services.Service{Model: model}
	handler := &handlers.URLHandler{Service: service}
	r.POST("/shorten", handler.GenerateURL)
}
