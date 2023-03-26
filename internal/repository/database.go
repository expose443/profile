package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/with-insomnia/profile/internal/config"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "profile"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func InstancePostgres(cfg *config.PostgresInfo) (*sql.DB, error) {
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	err = createTable(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTable(db *sql.DB) error {
	user := `CREATE TABLE IF NOT EXISTS users(
		user_id SERIAL,
		username VARCHAR(60) NOT NULL,
		password VARCHAR NOT NULL
	)`

	query := []string{}
	query = append(query, user)
	for _, v := range query {

		_, err := db.Exec(v)
		if err != nil {
			return err
		}
	}
	query1 := `INSERT INTO users(username, password) VALUES($1, $2)`
	_, err := db.Exec(query1, "abdu", "pass")
	if err != nil {
		return err
	}
	return nil
}
