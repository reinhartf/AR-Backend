package context

import (
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

// OpenDB creates and establishes the connection to PostgreSQL
func OpenDB(config *Config) (*pgx.Conn, error) {
	log.Println("Database is connecting... ")
	conn, err := pgx.ParseDSN(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName))
	if err != nil {
		panic(err.Error())
	}
	db, err := pgx.Connect(conn)

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}
