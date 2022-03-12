package main

import (
	"AllureResultParser/pkg"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	shutdown := make(chan error, 1)

	handler := pkg.Handler{}
	router := handler.InitRoutes()

	go func() {
		err := http.ListenAndServe(":3000", router)
		shutdown <- err
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

}
