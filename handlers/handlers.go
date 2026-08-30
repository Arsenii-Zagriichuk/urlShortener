package handlers

import (
	"net/http"
	"urlShortener/services"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	Service *services.Service
}

type URLRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func (h *URLHandler) GenerateURL(c *gin.Context) {
	var req URLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.Service.GenerateURL(req.URL)

	c.JSON(http.StatusOK, result)
}

// GetOGUrl get request handler
func (h *URLHandler) GetOGUrl(c *gin.Context) {

	shortCode := c.Param("code")

	originalURL, exists := h.Service.GetOGUrl(shortCode)

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "short link not found"})
		return
	}

	// 3. Issue the redirect to the target website
	c.Redirect(http.StatusFound, originalURL)

}
