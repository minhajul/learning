package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConn()

	defer db.Close()

	rows, err := db.Query("SELECT id, first_name, email FROM users")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	var users = []User{}
	for rows.Next() {
		var (
			id        int
			firstName string
			email     string
		)
		err = rows.Scan(&id, &firstName, &email)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, User{ID: id, Firstname: firstName, Email: email})
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(users)
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "butterfly"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func handleAllRequests() {
	http.HandleFunc("/", getUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleAllRequests()
}
