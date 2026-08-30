package main

import (
	"urlShortener/handlers"
	"urlShortener/models"
	"urlShortener/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	model := &models.SavedURL{Urls: make(map[string]string)}
	service := &services.Service{Model: model}
	handler := &handlers.URLHandler{Service: service}
	r.POST("/shorten", handler.GenerateURL)
	r.GET("/:code", handler.GetOGUrl)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
