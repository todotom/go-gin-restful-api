package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresClient struct {
	database *sqlx.DB
}

func ProvidePostgresClient() PostgresClient {
	database, err := sqlx.Connect(
		"postgres",
		"user=postgres password=postgres host=localhost port=5432 database=postgres sslmode=disable",
	)
	if err != nil {
		log.Fatalln(err)
	}
	return PostgresClient{database}
}
