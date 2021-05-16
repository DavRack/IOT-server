package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type state struct {
	// true encendido, false apagado
	LightState bool
}

var State state = state{
	LightState: false,
}

func SetLightState(c echo.Context) error {
	lightState := c.Param("state")

	if lightState == "on" {
		State.LightState = true
		return c.JSON(http.StatusOK, nil)
	}
	if lightState == "off" {
		State.LightState = false
		return c.JSON(http.StatusOK, nil)
	}
	return nil
}

func GetState(c echo.Context) (err error) {
	fmt.Println(State)
	return c.JSON(http.StatusOK, &State)
}
