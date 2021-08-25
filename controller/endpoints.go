package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var basePath string = "/api"
var EchoServer *echo.Echo = echo.New()

func init() {
	EchoServer.Use(middleware.Recover())
	EchoServer.Use(middleware.CORS())
	EchoServer.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} ${remote_ip} "${method} ${uri}" ${status} ${latency} ${bytes_in} ${bytes_out}` + "\n",
	}))
	EchoServer.Use(middleware.StaticWithConfig(middleware.StaticConfig{Root: "public", HTML5: true}))
	configAppRoutes()
}

func configAppRoutes() {
	api := EchoServer.Group(basePath)
	EchoServer.File("/", "PAE-GUI/index.html")

	api.GET("/device/:deviceId/state", GetState)
	api.POST("/device/:deviceId/state/:state", SetLightState)
}
