package controllers

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/aadeschamps/meeting-metrics/models"
	"github.com/gorilla/mux"
)

// UserController exports all methods needed to act on users
type UserController struct {
	User *models.UserModel
}

// Index retrieves and sends all users
func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := c.User.GetAll()
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

// Show retrieves a specific user by id
func (c *UserController) Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userToFind := models.User{ID: params["id"]}
	user, err := c.User.Get(&userToFind)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	user.Password = ""
	json.NewEncoder(w).Encode(user)
}

// Create creates a user and sends it
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	user, err := c.User.Create(&newUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Update updates a user and sends it
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userToUpdate := models.User{ID: params["id"]}
	err := json.NewDecoder(r.Body).Decode(&userToUpdate)
	user, err := c.User.Update(&userToUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Authenticate authenticates a user and sends it
func (c *UserController) Authenticate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

}
