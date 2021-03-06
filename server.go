package main

import (
	"net/http"

	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/cgi-microservice/handlers"
	"github.com/byuoitav/hateoas"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	port := ":8010"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	// GET requests
	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))

	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	secure.GET("/:address/volume/up", handlers.VolumeUp)
	secure.GET("/:address/volume/down", handlers.VolumeDown)
	secure.GET("/:address/volume/toggle-mute", handlers.Mute)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)

}
