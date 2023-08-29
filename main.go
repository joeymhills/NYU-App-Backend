package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/spatialcurrent/go-stringify/pkg/stringify"
)

// DSN=65xjbvp99e06f6krzt0x:pscale_pw_ztGVHxT3MSn3zTpg4741B1a9EYn7NZXiOCbVgJtFzxV@tcp(aws.connect.psdb.cloud)/nyu-db?tls=true&interpolateParams=true

type Accolade struct {
	id             string `json:"id"`
	name           string `json:"name"`
	institution    string `json:"institution"`
	outcome        string `json:"outcome"`
	serviceLine    string `json:"serviceLine"`
	extSource      string `json:"extSource"`
	intSource      string `json:"intSource"`
	messaging      string `json:"messaging"`
	comments       string `json:"comments"`
	frequency      string `json:"frequency"`
	notifDate      string `json:"notifDate"`
	cmcontact      string `json:"cmcontact"`
	sourceatr      string `json:"sourceatr"`
	wherepubint    string `json:"wherepubint"`
	promotionlim   string `json:"promotionlim"`
	effectiveDate  string `json:"effectiveDate"`
	expirationDate string `json:"expirationDate"`
	imgurl1        string `json:"imgurl1"`
	imgurl2        string `json:"imgurl2"`
	imgurl3        string `json:"imgurl3"`
	imgurl4        string `json:"imgurl4"`
}
type Employee struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
type DB struct {
	*sql.DB
}
type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

func getUsers(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		users := []User{}

		results, err := db.Query("SELECT id, email, name, role FROM user")
		if err != nil {
			panic(err.Error())
		}
		for results.Next() {
			var user User
			err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Role)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			person := User{
				Id: user.Id, Name: user.Name, Email: user.Email, Role: user.Role,
			}
			users = append(users, person)
		}

		// w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
func main() {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/getusers", getUsers(db))

	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}

	http.ListenAndServe("0.0.0.0:"+port, nil)

	log.Println("listening and serving")

}
