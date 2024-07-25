package api

import "database/sql"

type ApiServer struct {
	db *sql.DB
}

func NewApiServer(db *sql.DB) *ApiServer {
	return &ApiServer{
		db: db,
	}
}
