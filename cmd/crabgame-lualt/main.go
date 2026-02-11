package main

import (
	"log/slog"
	"net/http"

	"github.com/Lualttt/crabgame.lua.lt/internal/handlers"
)

func main() {
	slog.Info("starting http server on 0.0.0.0:8080")

	http.HandleFunc("GET /{lobbyID}/{playerID}", handlers.Join)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	slog.Error(err.Error())
}
