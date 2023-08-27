package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
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
type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", os.Getenv("DSN"))

	if err != nil {

		log.Fatalf("failed to connect: %v", err)

	}

	defer db.Close()

	users := []User{}

	results, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User
		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		person := User{
			Id: user.Id, Name: user.Name, Email: user.Email, Password: user.Password, Role: user.Role,
		}
		users = append(users, person)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/getusers", getUsers)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}

	http.ListenAndServe("0.0.0.0:"+port, nil)

	log.Println("listening and serving")

	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("server closed\n")
	// } else if err != nil {
	// 	fmt.Printf("error starting server: %s\n", err)
	// 	os.Exit(1)
	// }
}
