package controllers

import (
	"net/http"

	"encoding/json"

	"fmt"

	"github.com/aadeschamps/meeting-metrics/models"
)

// SessionController creates sessions
type SessionController struct {
	User *models.UserModel
}

// Create auth's the user
func (c *SessionController) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	_, err := c.User.Authenticate(&user)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Fprint(w, "ok")
	}
}
