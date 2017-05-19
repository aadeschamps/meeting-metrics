package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	// postgres driver for db
	_ "github.com/lib/pq"

	"os"

	"github.com/aadeschamps/meeting-metrics/controllers"
	"github.com/aadeschamps/meeting-metrics/middlewares"
	"github.com/aadeschamps/meeting-metrics/models"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	// instantiate router, database and middleware layer
	router := mux.NewRouter()

	var (
		username = os.Getenv("MMAPI_DB_USERNAME")
		password = os.Getenv("MMAPI_DB_PASSWORD")
		dbname   = os.Getenv("MMAPI_DB_NAME")
	)

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s",
		username, password, dbname)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	n := negroni.New(negroni.HandlerFunc(middlewares.Logger))

	// make models
	userModel := models.UserModel{Db: db}
	teamModel := models.TeamModel{Db: db}

	// make controllers
	userController := controllers.UserController{User: &userModel}
	teamController := controllers.TeamsController{Team: &teamModel}
	sessionController := controllers.SessionController{User: &userModel}

	// instantiate all api handlers
	api := router.PathPrefix("/api/v1/").Subrouter()
	api.HandleFunc("/login", sessionController.Create).Methods("POST")
	api.HandleFunc("/users", userController.Index).Methods("GET")
	api.HandleFunc("/users", userController.Create).Methods("POST")
	api.HandleFunc("/users/{id}", userController.Show).Methods("GET")
	api.HandleFunc("/users/{id}", userController.Update).Methods("PUT")

	api.HandleFunc("/groups/{id}", teamController.Show).Methods("GET")

	// instantiate static and index files serving
	fs := http.FileServer(http.Dir("./assets"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fs))
	router.PathPrefix("/").HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./assets/views/index.html")
	})

	// pass to negroni handler
	n.UseHandler(router)

	// create the server
	svr := &http.Server{
		Addr:           ":8080",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// prints server informations and starts it
	fmt.Println("Starting the server on localhost" + svr.Addr)
	svr.ListenAndServe()
}
