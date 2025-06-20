package logic

import (
	db "medicare/internal/dao/sqlc" // SQLC generated code
	// "medicare/utility/config"       // your postgres.go\
	"github.com/jackc/pgx/v5/pgxpool"
)

var Queries db.Querier

func InitSQLC(connPool *pgxpool.Pool) {
	Queries = db.New(connPool) // config.DB is *sql.DB
}
