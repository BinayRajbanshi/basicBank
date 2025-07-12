package util

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// log.Println(pgErr.Code)
		return pgErr.Code
	}
	return ""
}
