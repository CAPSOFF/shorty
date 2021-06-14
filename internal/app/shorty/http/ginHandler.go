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
		route.GET("/:shortcode", handler.HandleShortCode)
		route.GET("/:shortcode/stats", handler.HandleShortCodeStats)
	}
}

func (gh *ginHandler) HandleShorten(c *gin.Context) {
	var request struct {
		URL       string `json:"url"`
		ShortCode string `json:"shortCode"`
	}

	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "cannot decode request",
		})
		return
	}

	shortCode, httpStatusCode, err := gh.shortenController.Shorten(c, request.URL, request.ShortCode)
	if err != nil {
		c.AbortWithStatusJSON(httpStatusCode, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(httpStatusCode, gin.H{
		"shortCode": shortCode,
	})
}

func (gh *ginHandler) HandleShortCode(c *gin.Context) {
	var request struct {
		ShortCode string `uri:"shortcode" binding:"required"`
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}

	url, httpStatusCode, err := gh.shortenController.ShortCode(c, request.ShortCode)
	if err != nil {
		c.AbortWithStatusJSON(httpStatusCode, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}

	c.Header("Location", url)
	c.JSON(httpStatusCode, "")
}

func (gh *ginHandler) HandleShortCodeStats(c *gin.Context) {
	var request struct {
		ShortCode string `uri:"shortcode" binding:"required"`
	}

	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}

	shortyData, httpStatusCode, err := gh.shortenController.ShortCodeStats(c, request.ShortCode)
	if err != nil {
		c.AbortWithStatusJSON(httpStatusCode, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(httpStatusCode, gin.H{
		"startDate":     shortyData.StartDate,
		"lastSeenDate":  shortyData.LastSeenDate,
		"redirectCount": shortyData.RedirectCount,
	})
}
