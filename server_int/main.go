package main

import (
	"database/sql"
	"fmt"
	"log"
	"maxd935/server_int/clock"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	http.HandleFunc("/", clock.ClockHandler)
	http.HandleFunc("/test", toto("max"))

	// myEntries := make(map[string]string)

	// http.HandleFunc("/add", entry.EntryHandler)
	// http.HandleFunc("/add", test(myEntries, entry.EntryHandler))

	// http.HandleFunc("/entries", entry.EntryHandler)

	database()
	http.ListenAndServe(":4567", nil)
}

func toto(str string) func(w http.ResponseWriter, req *http.Request) {
	testHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(str)
		switch req.Method {
		case http.MethodGet:
			fmt.Fprintf(w, str)
		}
	}
	return testHandler
}

// func test(myEntries map[string]string) func() {
// 	EntryHandler := func (w http.ResponseWriter, req *http.Request) {

// 		fmt.Println("EntryHandler")
// 		fmt.Println(myEntries)

// 		switch req.Method {
// 		case http.MethodGet:
// 			if len(myEntries) > 0 {
// 				for i, mess := range myEntries {
// 					fmt.Fprintf(w, "index: %s\n", i)
// 					fmt.Fprintf(w, "mess: %s\n", mess)
// 					fmt.Fprintf(w, "\n")
// 				}
// 			} else {
// 				fmt.Fprintf(w, "No Message")
// 			}
// 		case http.MethodPost:
// 			if err := req.ParseForm(); err != nil {
// 				fmt.Println("Something went bad")
// 				fmt.Fprintln(w, "Something went bad")
// 				return
// 			}
// 			for key, value := range req.PostForm {
// 				fmt.Println(key, "=>", value[0])
// 				myEntries[key] = value[0]
// 			}
// 			fmt.Println(myEntries)

// 			fmt.Fprintf(w, "Information received: %v\n", req.PostForm)
// 		}
// 	}
// 	return EntryHandler
// }

func database() {
	// The `sql.Open` function opens a new `*sql.DB` instance. We specify the driver name
	// and the URI for our database. Here, we're using a Postgres URI
	db, err := sql.Open("pgx", "postgresql://localhost:5432/db_go?user=user&password=pwd")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")
}
