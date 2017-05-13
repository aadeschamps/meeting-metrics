package models

import (
	"database/sql"
)

// TeamModel contains db and performs actions
type TeamModel struct {
	Db *sql.DB
}

// Team respresents a group from the db
type Team struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Users []*User `json:"users"`
}
