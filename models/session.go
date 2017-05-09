package models

import (
	"database/sql"
)

// SessionModel contains db and performs actions
type SessionModel struct {
	Db *sql.DB
}

// Session respresents a session from the db
type Session struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	TeamID string `json:"team_id"`
}
