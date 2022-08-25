package context

import (
	"database/sql"
	"fmt"
)

func Connect(host, name string) *sql.DB {
	db, err := sql.Open("sqlserver", fmt.Sprintf("server=%s;database=%s;trusted_connection=yes", host, name))

	if err != nil {
		panic(err)
	}
	return db
}
