package models

import (
	"database/sql"
)

// CategoryModel contains db and performs actions
type CategoryModel struct {
	Db *sql.DB
}

// Category respresents a category from the db
type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	MeetingID string `json:"meeting_id"`
}
