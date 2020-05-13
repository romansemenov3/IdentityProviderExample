package repository

import (
	"common"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var DB *sql.DB

type config struct {
	Database databaseConfig `yaml:"database"`
}

type databaseConfig struct {
	Url string `yaml:"url"`
}

func init() {
	cfg := config{}
	err := common.ReadConfig(&cfg)
	if err != nil {
		panic(err)
	}

	DB, err = sql.Open("pgx", cfg.Database.Url)
	if err != nil {
		panic(err)
	}
}
