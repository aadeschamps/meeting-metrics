package models

import (
	"database/sql"
)

// MeetingModel contains db and performs actions
type MeetingModel struct {
	Db *sql.DB
}

// Meeting respresents a meeting from the db
type Meeting struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	TeamID string `json:"team_id"`
}
