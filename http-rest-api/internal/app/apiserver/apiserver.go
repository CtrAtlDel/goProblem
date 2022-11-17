package apiserver

import (
	"database/sql"
	sqlstore "http-rest-api/internal/app/store/sqlStore"
	"net/http"
)

func Start(c *Config) error {
	db, err := newDB(c.DataBaseUrl)
	if err != nil {
		return err
	}
	store := sqlstore.New(db)
	srv := newServer(store)
	return http.ListenAndServe(c.BindAddr, srv)
}

func newDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {

	}

	return db, nil
}
