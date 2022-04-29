package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/denisandreenko/http-rest-api/internal/app/store/sqlstore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	sessionKey := securecookie.GenerateRandomKey(32)
	sessionsStore := sessions.NewCookieStore(sessionKey)
	store := sqlstore.New(db)
	s := newServer(store, sessionsStore)

	return http.ListenAndServe(config.BindAddr, s)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
