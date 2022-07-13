package logger

import (
	"context"
	"database/sql"
	"github.com/douglmendes/mercado-fresco-round-go/connections"
)

const createQuery = "insert into logs (level, timestamp, caller, msg) values(?, ?, ?, ?)"

func CreateLog(ctx context.Context, level, timestamp, caller, msg string) error {
	var db *sql.DB
	db = connections.NewConnection()

	_, err := db.ExecContext(ctx, createQuery, level, timestamp, caller, msg)

	return err
}
