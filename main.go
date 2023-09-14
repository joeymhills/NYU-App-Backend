package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/spatialcurrent/go-stringify/pkg/stringify"
)


// type NullString struct {
// 	sql.NullString
// }

// func (ns *NullString) MarshalJSON() ([]byte, error) {
// 	if !ns.Valid {
// 		return []byte("null"), nil
// 	}
// 	return json.Marshal(ns.String)
// }

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
	CreatedAt      string `json:"createdAt"`
	Imgurl1        string `json:"imgurl1"`
	Imgurl2        string `json:"imgurl2"`
	Imgurl3        string `json:"imgurl3"`
	Imgurl4        string `json:"imgurl4"`
	Supported      bool   `json:"supported"`
}

type BackupAward struct {
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
	DeletedAt      string `json:"deletedAt"`
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
type FindId struct {
	Id    string    `json:"id"`
}

func recentAwards(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		awards := []Award{}
		results, err := db.Query("SELECT id, name, institution, outcome, serviceLine, extSource, intSource, messaging, comments, frequency, notifDate, cmcontact, sourceatr, wherepubint, promotionlim, IFNULL(expirationDate,''), IFNULL(effectiveDate,''), IFNULL(imgurl1,''),IFNULL(imgurl2,''),IFNULL(imgurl3,''), IFNULL(imgurl4,''), supported, createdAt FROM accolade ORDER BY createdat DESC LIMIT 4")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			panic(err.Error())
		}
		for results.Next() {
			var award Award
			err = results.Scan(&award.Id, &award.Name, &award.Institution, &award.Outcome, &award.ServiceLine,
				&award.ExtSource, &award.IntSource, &award.Messaging, &award.Comments, &award.Frequency, &award.NotifDate,
				&award.Cmcontact, &award.Sourceatr, &award.Wherepubint, &award.Promotionlim, &award.ExpirationDate,
				&award.EffectiveDate, &award.Imgurl1, &award.Imgurl2, &award.Imgurl3, &award.Imgurl4, &award.Supported, &award.CreatedAt)
			if err != nil {
				log.Println(err)
				panic(err.Error()) // proper error handling instead of panic in your apps
			}
			awardStruct := Award{
				Id: award.Id, Name: award.Name, Institution: award.Institution, Outcome: award.Outcome, ServiceLine: award.ServiceLine,
				ExtSource: award.ExtSource, IntSource: award.IntSource, Messaging: award.Messaging, Comments: award.Comments, Frequency: award.Frequency,
				NotifDate: award.NotifDate, Cmcontact: award.Cmcontact, Sourceatr: award.Sourceatr, Wherepubint: award.Wherepubint, Promotionlim: award.Promotionlim,
				Supported: award.Supported, CreatedAt: award.CreatedAt, EffectiveDate: award.EffectiveDate, ExpirationDate: award.ExpirationDate,
				Imgurl1: award.Imgurl1, Imgurl2: award.Imgurl2, Imgurl3: award.Imgurl3, Imgurl4: award.Imgurl4,
			}
			awards = append(awards, awardStruct)
		}

		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(awards)
	}
}

func findAward(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
        

		request, err := io.ReadAll(r.Body)

        if err != nil {
            log.Fatal(err)
        }
// unmarshal the request into findId struct
        var findId FindId 
        
        err = json.Unmarshal(request, &findId)
        if err != nil {
            panic(err)
        }
        log.Println(findId.Id)
        query := findId.Id 
        log.Println(query)

		awards := []Award{}
		results, err := db.Query("SELECT id, name, institution, outcome, serviceLine, extSource, intSource, messaging, comments, frequency, notifDate, cmcontact, sourceatr, wherepubint, promotionlim, IFNULL(expirationDate,''), IFNULL(effectiveDate,''), IFNULL(imgurl1,''),IFNULL(imgurl2,''),IFNULL(imgurl3,''), IFNULL(imgurl4,''), supported, createdAt FROM accolade WHERE id = ?", query)
		if err != nil {
            log.Println("error in sql query")
			http.Error(w, err.Error(), http.StatusBadRequest)
			panic(err.Error())
		}
		for results.Next() {
			var award Award
			err = results.Scan(&award.Id, &award.Name, &award.Institution, &award.Outcome, &award.ServiceLine,
				&award.ExtSource, &award.IntSource, &award.Messaging, &award.Comments, &award.Frequency, &award.NotifDate,
				&award.Cmcontact, &award.Sourceatr, &award.Wherepubint, &award.Promotionlim, &award.ExpirationDate,
				&award.EffectiveDate, &award.Imgurl1, &award.Imgurl2, &award.Imgurl3, &award.Imgurl4, &award.Supported, &award.CreatedAt)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your apps
			}
			awardStruct := Award{
				Id: award.Id, Name: award.Name, Institution: award.Institution, Outcome: award.Outcome, ServiceLine: award.ServiceLine,
				ExtSource: award.ExtSource, IntSource: award.IntSource, Messaging: award.Messaging, Comments: award.Comments, Frequency: award.Frequency,
				NotifDate: award.NotifDate, Cmcontact: award.Cmcontact, Sourceatr: award.Sourceatr, Wherepubint: award.Wherepubint, Promotionlim: award.Promotionlim,
				Supported: award.Supported, CreatedAt: award.CreatedAt, EffectiveDate: award.EffectiveDate, ExpirationDate: award.ExpirationDate,
				Imgurl1: award.Imgurl1, Imgurl2: award.Imgurl2, Imgurl3: award.Imgurl3, Imgurl4: award.Imgurl4,
			}
			awards = append(awards, awardStruct)
		}

		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(awards)
	}
}

func searchAwards(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		search, err := io.ReadAll(r.Body)
		s := string(search)

		//sql query where name like %s%

		awards := []Award{}
        query := "%" + s + "%"
        results, err := db.Query("SELECT id, name, institution, outcome, serviceLine, extSource, intSource, messaging, comments, frequency, notifDate, cmcontact, sourceatr, wherepubint, promotionlim, IFNULL(expirationDate,''), IFNULL(effectiveDate,''), IFNULL(imgurl1,''),IFNULL(imgurl2,''),IFNULL(imgurl3,''), IFNULL(imgurl4,''), supported, createdAt FROM accolade WHERE name LIKE ?", query)
        if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			panic(err.Error())
		}
		for results.Next() {
			var award Award
			err = results.Scan(&award.Id, &award.Name, &award.Institution, &award.Outcome, &award.ServiceLine,
				&award.ExtSource, &award.IntSource, &award.Messaging, &award.Comments, &award.Frequency, &award.NotifDate,
				&award.Cmcontact, &award.Sourceatr, &award.Wherepubint, &award.Promotionlim, &award.ExpirationDate,
				&award.EffectiveDate, &award.Imgurl1, &award.Imgurl2, &award.Imgurl3, &award.Imgurl4, &award.Supported, &award.CreatedAt)
			if err != nil {
				log.Println(err)
				panic(err.Error()) // proper error handling instead of panic in your apps
			}
			awardStruct := Award{
				Id: award.Id, Name: award.Name, Institution: award.Institution, Outcome: award.Outcome, ServiceLine: award.ServiceLine,
				ExtSource: award.ExtSource, IntSource: award.IntSource, Messaging: award.Messaging, Comments: award.Comments, Frequency: award.Frequency,
				NotifDate: award.NotifDate, Cmcontact: award.Cmcontact, Sourceatr: award.Sourceatr, Wherepubint: award.Wherepubint, Promotionlim: award.Promotionlim,
				Supported: award.Supported, CreatedAt: award.CreatedAt, EffectiveDate: award.EffectiveDate, ExpirationDate: award.ExpirationDate,
				Imgurl1: award.Imgurl1, Imgurl2: award.Imgurl2, Imgurl3: award.Imgurl3, Imgurl4: award.Imgurl4,
			}
			awards = append(awards, awardStruct)
		}

		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Origin", "*")
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

		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// function that gets all deleted awards from accoladeBackup table
func getDeleted(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		awards := []BackupAward{}
		results, err := db.Query("SELECT id, name, institution, outcome, IFNULL(serviceLine,''), extSource, intSource, messaging, comments, frequency, notifDate, cmcontact, sourceatr, wherepubint, promotionlim, IFNULL(expirationDate,''), IFNULL(effectiveDate,''), IFNULL(imgurl1,''),IFNULL(imgurl2,''),IFNULL(imgurl3,''), IFNULL(imgurl4,''), supported, deletedAt FROM accoladeBackup")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			panic(err.Error())
		}
		for results.Next() {
			var award BackupAward
			err = results.Scan(&award.Id, &award.Name, &award.Institution, &award.Outcome, &award.ServiceLine,
				&award.ExtSource, &award.IntSource, &award.Messaging, &award.Comments, &award.Frequency, &award.NotifDate,
				&award.Cmcontact, &award.Sourceatr, &award.Wherepubint, &award.Promotionlim, &award.ExpirationDate,
				&award.EffectiveDate, &award.Imgurl1, &award.Imgurl2, &award.Imgurl3, &award.Imgurl4, &award.Supported, &award.DeletedAt)
			if err != nil {
				log.Println(err)
				panic(err.Error()) // proper error handling instead of panic in your apps
			}
			awardStruct := BackupAward{
				Id: award.Id, Name: award.Name, Institution: award.Institution, Outcome: award.Outcome, ServiceLine: award.ServiceLine,
				ExtSource: award.ExtSource, IntSource: award.IntSource, Messaging: award.Messaging, Comments: award.Comments, Frequency: award.Frequency,
				NotifDate: award.NotifDate, Cmcontact: award.Cmcontact, Sourceatr: award.Sourceatr, Wherepubint: award.Wherepubint, Promotionlim: award.Promotionlim,
				Supported: award.Supported, DeletedAt: award.DeletedAt, EffectiveDate: award.EffectiveDate, ExpirationDate: award.ExpirationDate,
				Imgurl1: award.Imgurl1, Imgurl2: award.Imgurl2, Imgurl3: award.Imgurl3, Imgurl4: award.Imgurl4,
			}
			awards = append(awards, awardStruct)
		}

		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(awards)
	}
}

func main() {

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}
    log.Println("DB connected and ready to serve🫡 ")

	http.HandleFunc("/getusers", getUsers(db))
	http.HandleFunc("/getdeleted", getDeleted(db))
	http.HandleFunc("/search", searchAwards(db))
	http.HandleFunc("/recentawards", recentAwards(db))
	http.HandleFunc("/findaward", findAward(db))
	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}

	http.ListenAndServe("0.0.0.0:"+port, nil)

	log.Println("listening and serving")

}
