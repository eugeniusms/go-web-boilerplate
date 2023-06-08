package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
    host     = os.Getenv("DB_HOST")
    port     = os.Getenv("DB_PORT")
    user     = os.Getenv("DB_USER")
    password = os.Getenv("DB_PASSWORD")
    dbname   = os.Getenv("DB_NAME")
)

func TestOpenConnection(t *testing.T) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}