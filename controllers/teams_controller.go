package controllers

import (
	"net/http"

	"fmt"

	"github.com/aadeschamps/meeting-metrics/models"
)

// TeamsController exports all methods needed to act on users
type TeamsController struct {
	Team *models.TeamModel
}

// Show retrieves a specific team by id
func (c *TeamsController) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, groups!")
}
