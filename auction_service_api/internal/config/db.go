package config

import "github.com/jmoiron/sqlx"

func NewPgDb(conf Config) *sqlx.DB {
	db, err := sqlx.Open("postgres", conf.Get("SDN"))
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
