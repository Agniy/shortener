package main

import (
	"fmt"
	"github.com/Agniy/shortener/internal/app/config"
	"github.com/Agniy/shortener/internal/app/handler"
	"github.com/Agniy/shortener/internal/app/middleware"
	"github.com/Agniy/shortener/internal/app/models"
	"github.com/Agniy/shortener/internal/app/storage"

	"net/http"
)

func main() {
	cfg := config.GetConfig()

	// create db connection and migrate models
	// ------------------------------------------
	psql, err := storage.GetDbClient()
	if err != nil {
		panic(err)
	}
	err = models.MigrateAllModels(psql)
	if err != nil {
		fmt.Println(err)
	}
	// ------------------------------------------

	mux := http.NewServeMux()
	mux.Handle(`/api/shorten`, middleware.LoggingMiddleware(http.HandlerFunc(handler.MainPage)))

	fmt.Println("Starting server on port:", cfg.IP+":"+cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		panic(err)
	}
}
