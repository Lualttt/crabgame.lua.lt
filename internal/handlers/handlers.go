package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Lualttt/crabgame.lua.lt/web"
)

func Join(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(web.Templates, "join.tmpl")
	if err != nil {
		http.Error(w, "Internal server error!", http.StatusInternalServerError)
		slog.Error("join page error", "error", err)
		return
	}

	lobbyID, err := strconv.Atoi(r.PathValue("lobbyID"))
	if err != nil {
		http.Error(w, "Internal server error!", http.StatusInternalServerError)
		slog.Error("join page error", "error", err)
		return
	}

	playerID, err := strconv.Atoi(r.PathValue("playerID"))
	if err != nil {
		http.Error(w, "Internal server error!", http.StatusInternalServerError)
		slog.Error("join page error", "error", err)
		return
	}

	err = t.Execute(w, fmt.Sprintf("steam://joinlobby/1782210/%d/%d", lobbyID, playerID))
	if err != nil {
		http.Error(w, "Internal server error!", http.StatusInternalServerError)
		slog.Error("join page error", "error", err)
		return
	}
}
