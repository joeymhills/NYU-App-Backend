package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
    "time"
    "github.com/joeymhills/go-sql-api/handlers"
	"github.com/patrickmn/go-cache"

	_ "github.com/go-sql-driver/mysql"
)


func main() {

    c := cache.New(1*time.Minute, 10*time.Minute)

	db, err := sql.Open("mysql", "3sujnpq09n7n64wfzw0o:pscale_pw_kSSIWrx1QEXPwy2J1dCJj0uPiHDx0uvr8AJ4qyN9AOv@tcp(aws.connect.psdb.cloud)/nyu-db?tls=true&interpolateParams=true")
	if err != nil {
		log.Fatal(err)
	}
    log.Println("DB connected and ready to serve🫡🫡🫡🫡🫡 ")

	http.HandleFunc("/find", handlers.FindAward(db, c))
	http.HandleFunc("/getusers", handlers.GetUsers(db))
	http.HandleFunc("/getdeleted", handlers.GetDeleted(db))
	http.HandleFunc("/search", handlers.SearchAwards(db))
	http.HandleFunc("/recentawards", handlers.RecentAwards(db))
	
    port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	http.ListenAndServe("0.0.0.0:"+port, nil)

	log.Println("listening and serving")

}
