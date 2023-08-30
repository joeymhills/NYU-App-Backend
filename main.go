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

type Award struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Institution    string `json:"institution"`
	Outcome        string `json:"outcome"`
	ServiceLine    string `json:"serviceLine"`
	ExtSource      string `json:"extSource"`
	IntSource      string `json:"intSource"`
	Messaging      string `json:"messaging"`
	Comments       string `json:"comments"`
	Frequency      string `json:"frequency"`
	NotifDate      string `json:"notifDate"`
	Cmcontact      string `json:"cmcontact"`
	Sourceatr      string `json:"sourceatr"`
	Wherepubint    string `json:"wherepubint"`
	Promotionlim   string `json:"promotionlim"`
	EffectiveDate  string `json:"effectiveDate"`
	ExpirationDate string `json:"expirationDate"`
	Imgurl1        string `json:"imgurl1"`
	Imgurl2        string `json:"imgurl2"`
	Imgurl3        string `json:"imgurl3"`
	Imgurl4        string `json:"imgurl4"`
	Supported      bool   `json:"supported"`
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

func searchAwards(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// awards := []Award{}
		search := r.Body
		log.Println(search)

		// results, err := db.Query("SELECT * FROM accolade")
		// if err != nil {
		// 	panic(err.Error())
		// }
		// for results.Next() {
		// 	var award Award
		// 	err = results.Scan(&.Id, &user.Name, &user.Email, &user.Role)
		// 	if err != nil {
		// 		panic(err.Error()) // proper error handling instead of panic in your apps
		// 	}
		// 	person := User{
		// 		Id: user.Id, Name: user.Name, Email: user.Email, Role: user.Role,
		// 	}
		// 	awards = append(awards, award)
		// }

		// w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
		// w.Header().Set("Access-Control-Allow-Origin", "https://nyu-award.vercel.app")
		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(users)
	}
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
				panic(err.Error()) // proper error handling instead of panic in your apps
			}
			person := User{
				Id: user.Id, Name: user.Name, Email: user.Email, Role: user.Role,
			}
			users = append(users, person)
		}

		// w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST")
		// meoww
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Origin", "https://nyu-award.vercel.app")
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
	http.HandleFunc("/search", searchAwards(db))

	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}

	http.ListenAndServe("0.0.0.0:"+port, nil)

	log.Println("listening and serving")

}
