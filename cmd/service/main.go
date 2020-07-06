package main

import (
	"physiobot/cmd/router"
	"physiobot/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/getsentry/sentry-go"
)

func main() {
	/*
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://bb55a2c65af446b39d99e00957cba3dc@o396472.ingest.sentry.io/5249867",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	*/

	server := echo.New()

	//middleware
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	server.Use(middleware.Recover())
	server.Use(middleware.BodyLimit("2M"))
	server.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 7,
	}))

	//cors
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	server = router.Init(server)

	configuration := config.GetConfig()
	server.Logger.Fatal(server.Start(":" + configuration.App.Port))
}
