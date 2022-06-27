package connections

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewConnection() *sql.DB {
	client, err := sql.Open("mysql", "root:12345@tcp(localhost:3306)/mercado_fresco")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
