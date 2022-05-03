package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/denisandreenko/http-rest-api/internal/app/store/sqlstore"
	"github.com/denisandreenko/http-rest-api/internal/utils"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const _sessionKey = "SESSION_KEY"

func Start(config *Config) error {
	sessionKey := utils.Getenv(_sessionKey)
	if sessionKey == "" {
		sessionKey = string(securecookie.GenerateRandomKey(32))
		utils.Setenv(_sessionKey, sessionKey)
	}
	sessionsStore := sessions.NewCookieStore([]byte(sessionKey))

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

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
