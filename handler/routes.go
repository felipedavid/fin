package handler

import (
	"database/sql"
	"net/http"
)

func SetupRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", renderLoginPage)

	return mux
}
