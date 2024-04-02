package db

import (
	"fmt"
	"go-fiber-api-docker/pkg/common/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(c *config.Config) *sqlx.DB {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.DBPort, c.User, c.Password, c.DBname)

	db, err := sqlx.Open("postgres", url)

	if err != nil {
		log.Fatalln(err)
	}

	migrateDB(db)

	return db
}

func migrateDB(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT,
		description TEXT,
		price DECIMAL,
		stock BIGINT
	)`)

	if err != nil {
		log.Fatalln(err)
	}
}
