package main

import (
	"database/sql"
	"fin/handler"
	"flag"
	"log/slog"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	listenAddr := *flag.String("addr", "127.0.0.1:6969", "Interface addr and port to bind to")
	dbURL := *flag.String("db", "postgres://postgres:postgres@localhost/fin?sslmode=disable", "Database url")
	flag.Parse()

	db, err := connectToDatabase(dbURL)
	shouldNotErrorOut(err)

	routes := handler.SetupRouter(db)
	server := &http.Server{
		Addr:    listenAddr,
		Handler: routes,
	}

	slog.Info("Starting web server", "addr", listenAddr)
	err = server.ListenAndServe()
	shouldNotErrorOut(err)
}

func connectToDatabase(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func shouldNotErrorOut(err error) {
	if err != nil {
		panic(err)
	}
}
