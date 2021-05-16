package main

import (
	"fmt"
	"iot/controller"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	serverPort := 8080
	go func() {
		if err := controller.EchoServer.Start(fmt.Sprintf(":%v", serverPort)); err != nil && err != http.ErrServerClosed {
			panic("Error fatal")
		}
	}()

	// esperar hasta que se envie una señal para detener la aplicación
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
