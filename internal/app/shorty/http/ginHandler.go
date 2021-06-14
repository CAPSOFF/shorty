package http

import (
	"net/http"

	shorty "github.com/amartha-shorty/pkg/shorty"
	"github.com/gin-gonic/gin"
)

type ginHandler struct {
	shortenController shorty.Controller
}

const basePath = "api/v1/amartha/shorty"

// NewGinHandler will initialize http endpoint for delivery
func NewGinHandler(router *gin.Engine, shortenController shorty.Controller) {
	handler := &ginHandler{shortenController: shortenController}
	route := router.Group(basePath)
	{
		route.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "test API")
		})
		route.POST("/shorten", handler.HandleShorten)
		// route.GET("/:shortcode", handler.HandleShorten)
		// route.GET("/:shortcode/stats", handler.HandleShorten)
	}
}

func (gh *ginHandler) HandleShorten(c *gin.Context) {
	var request struct {
		URL       string `json:"url"`
		ShortCode string `json:"shortCode"`
	}

	c.Header("Content-Type", "application/json")
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "cannot decode request",
		})
		return
	}

	gh.shortenController.Shorten(c)
	// controller here

	c.JSON(http.StatusCreated, gin.H{
		"shortCode": "test-short123",
	})
}
