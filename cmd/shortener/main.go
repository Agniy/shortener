package main

import (
	"fmt"
	"github.com/Agniy/shortener/internal/app/config"
	"github.com/Agniy/shortener/internal/app/handler"
	"net/http"
)

func main() {
	cfg := config.GetConfig()
	mux := http.NewServeMux()
	mux.Handle(`/`, http.HandlerFunc(handler.MainPage))

	fmt.Println("Starting server on port:", cfg.IP+":"+cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		panic(err)
	}
}
