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

    c := cache.New(10*time.Second, 1*time.Minute)

	db, err := sql.Open("mysql", os.Getenv("DSN"))
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
