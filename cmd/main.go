package main

import (
	"log"

	todoHttp "github.com/Domish4/GoHttp"
	"github.com/Domish4/GoHttp/package/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todoHttp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error curred while running http server: %s", err.Error())
	}

}
