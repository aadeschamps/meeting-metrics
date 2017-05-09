package controllers

import (
	"net/http"

	"fmt"

	"github.com/aadeschamps/giftexchangeapi/models"
)

// GroupController exports all methods needed to act on users
type TeamsController struct {
	Group *models.TeamModel
}

// Show retrieves a specific group by id
func (c *TeamsController) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, groups!")
}
