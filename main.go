package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type users struct {
	ID   string `json:"userId"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func main() {

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	for results.Next() {
		var users users
		// for each row, scan the result into our tag composite object
		err := results.Scan(&users.ID, &users.Name, &users.Role)
		if err != nil {
			log.Printf("error in scanner: %v", err) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(users.Name)
	}

	handleRequests()

}
