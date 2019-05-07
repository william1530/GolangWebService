package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	fmt.Print("calling init db function")
	InitDB("none")

	http.HandleFunc("/", handler)
	http.HandleFunc("/tstvacation", tstvacation)
	http.HandleFunc("/tstVacation", tstVacation)
	http.HandleFunc("/tstSQLite", tstSQLite)

	/*Ohne spezifischen Server:
	log.Fatal(http.ListenAndServe(":8080", nil))
	*/

	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    300 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())

}
