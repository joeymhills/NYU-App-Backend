package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/spatialcurrent/go-stringify/pkg/stringify"
)

// DSN=65xjbvp99e06f6krzt0x:pscale_pw_ztGVHxT3MSn3zTpg4741B1a9EYn7NZXiOCbVgJtFzxV@tcp(aws.connect.psdb.cloud)/nyu-db?tls=true&interpolateParams=true

type NullString string

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, ok := value.(string)
	if !ok {
		return errors.New("Column is not a string")
	}
	*s = NullString(strVal)
	return nil
}
func (s NullString) Value() (driver.Value, error) {
	if len(s) == 0 { // if nil or empty string
		return nil, nil
	}
	return string(s), nil
}

type Award struct {
	Id             string     `json:"id"`
	Name           string     `json:"name"`
	Institution    string     `json:"institution"`
	Outcome        string     `json:"outcome"`
	ServiceLine    string     `json:"serviceLine"`
	ExtSource      string     `json:"extSource"`
	IntSource      string     `json:"intSource"`
	Messaging      string     `json:"messaging"`
	Comments       string     `json:"comments"`
	Frequency      string     `json:"frequency"`
	NotifDate      string     `json:"notifDate"`
	Cmcontact      string     `json:"cmcontact"`
	Sourceatr      string     `json:"sourceatr"`
	Wherepubint    string     `json:"wherepubint"`
	Promotionlim   string     `json:"promotionlim"`
	EffectiveDate  string     `json:"effectiveDate"`
	ExpirationDate string     `json:"expirationDate"`
	CreatedAt      string     `json:"createdAt"`
	Imgurl1        NullString `json:"imgurl1,omitempty"`
	Imgurl2        NullString `json:"imgurl2,omitempty"`
	Imgurl3        NullString `json:"imgurl3,omitempty"`
	Imgurl4        NullString `json:"imgurl4,omitempty"`
	Supported      bool       `json:"supported"`
}
type Employee struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
type DB struct {
	*sql.DB
}
type SearchBody struct {
	Search string `json:"search"`
}
type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

func searchAwards(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var s SearchBody
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// log.Println(s)

		//sql query where name like %s%

		awards := []Award{}
		results, err := db.Query("SELECT * FROM accolade")
		if err != nil {
			panic(err.Error())
		}
		for results.Next() {
			var award Award
			err = results.Scan(&award.Id, &award.Name, &award.Institution, &award.Outcome, &award.ServiceLine,
				&award.ExtSource, &award.IntSource, &award.Messaging, &award.Comments, &award.Frequency, &award.NotifDate,
				&award.Cmcontact, &award.Sourceatr, &award.Wherepubint, &award.Promotionlim, &award.EffectiveDate,
				&award.ExpirationDate, &award.Imgurl1, &award.Imgurl2, &award.Imgurl3, &award.Imgurl4, &award.Supported, &award.CreatedAt)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your apps
			}
			awardStruct := Award{
				Id: award.Id, Name: award.Name, Institution: award.Institution, Outcome: award.Outcome, ServiceLine: award.ServiceLine,
				ExtSource: award.ExtSource, IntSource: award.IntSource, Messaging: award.Messaging, Comments: award.Comments, Frequency: award.Frequency,
				NotifDate: award.NotifDate, Cmcontact: award.Cmcontact, Sourceatr: award.Sourceatr, Wherepubint: award.Wherepubint, Promotionlim: award.Promotionlim,
				EffectiveDate: award.EffectiveDate, ExpirationDate: award.ExpirationDate, Imgurl1: award.Imgurl1, Imgurl2: award.Imgurl2, Imgurl3: award.Imgurl3,
				Imgurl4: award.Imgurl4, Supported: award.Supported, CreatedAt: award.CreatedAt,
			}
			awards = append(awards, awardStruct)
		}

		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Origin", "https://nyu-award.vercel.app")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(awards)
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
