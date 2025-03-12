package main

import (
	"auction_service_api/internal/config"
)

func main() {
	conf := config.NewConfig()
	db := config.NewPgDb(conf)
	_ = db
}
