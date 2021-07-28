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

var devices = make(map[string]state)

func SetLightState(c echo.Context) error {
	lightState := c.Param("state")
	deviceId := c.Param("deviceId")

	if lightState == "on" {
		devices[deviceId] = state{LightState: true}
		return c.JSON(http.StatusOK, nil)
	}
	if lightState == "off" {
		devices[deviceId] = state{LightState: false}
		return c.JSON(http.StatusOK, nil)
	}
	return nil
}

func GetState(c echo.Context) (err error) {
	deviceId := c.Param("deviceId")
	State, ok := devices[deviceId]
	if !ok {
		fmt.Println("adding", deviceId, "to database")
		devices[deviceId] = state{LightState: false}
	}
	fmt.Println(State)
	return c.JSON(http.StatusOK, &State)
}
