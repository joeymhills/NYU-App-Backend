package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
    "time"
    "sync"
    "github.com/joeymhills/go-sql-api/handlers"
	"github.com/patrickmn/go-cache"

	_ "github.com/go-sql-driver/mysql"
)


func main() {

    var wg sync.WaitGroup
    c := cache.New(10*time.Second, 1*time.Minute)

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}
    log.Println("DB connected and ready to serve🫡🫡🫡🫡🫡 ")

	http.HandleFunc("/changerole", handlers.ChangeRole(db, &wg))
	http.HandleFunc("/find", handlers.FindAward(db, c))
	http.HandleFunc("/getusers", handlers.GetUsers(db))
	http.HandleFunc("/getdeleted", handlers.GetDeleted(db))
	http.HandleFunc("/search", handlers.SearchAwards(db))
	http.HandleFunc("/recentawards", handlers.RecentAwards(db))
    http.HandleFunc("/create", handlers.CreateAward(db, c))
    http.HandleFunc("/update", handlers.UpdateAward(db, c))
	
    port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	http.ListenAndServe("0.0.0.0:"+port, nil)

	log.Println("listening and serving")

}
