package internalsql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error conneting to database %+v\n", err)
	}

	return db, nil
}
