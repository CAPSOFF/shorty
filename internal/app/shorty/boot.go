package amartha

import (
	"time"

	_shortyController "github.com/amartha-shorty/pkg/shorty/controller"
	_shortyRepository "github.com/amartha-shorty/pkg/shorty/repository"
	"github.com/gin-gonic/gin"

	"github.com/amartha-shorty/internal/app/shorty/http"
)

// Start to initialize sirclo-berat application
func Start(opt Options) error {
	shortyRepository := _shortyRepository.NewShortyRepository()
	shortyController := _shortyController.NewShortyController(time.Second*2, shortyRepository)

	engine := gin.New()
	http.NewGinHandler(engine, shortyController)

	return engine.Run(":7777")
}
