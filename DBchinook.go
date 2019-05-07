package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

//Customer object from chinook db
type Customer struct {
	CustomerID int    `db:"CustomerId"`
	FirstName  string `db:"FirstName"`
	LastName   string `db:"LastName"`
	Company    string `db:"Company" json:",omitempty"`
	Country    string `db:"Country"`
	Email      string `db:"Email"`
}

//CustomersEnvelope ist ein Kontainer für die Customer Entität, wird benötigt für XML marshalling
type CustomersEnvelope struct {
	XMLName xml.Name   `xml:"Customers"`
	List    []Customer `xml:"Customer"`
}

//Handler2 makes a select on the db
func tstSQLite(w http.ResponseWriter, r *http.Request) {

	var err error
	var data []byte
	var fehlerMessage string

	contentTypeRequest := r.Header.Get("Content-type")

	Customers := []Customer{}

	fmt.Print("executing query")
	//Select data and write it to variable
	err = chinookdb.Select(&Customers, "SELECT CustomerId, FirstName,LastName, ifnull(Company,'') Company, Country,Email FROM customers")

	if err != nil {
		fehlerMessage = err.Error()
		http.Error(w, "Datenbankfehler: "+fehlerMessage, http.StatusInternalServerError)
		return
	}

	//Wenn XML angefordert Transformiere DBQuery zu XML
	if contentTypeRequest == "application/xml" || true {
		data, err = xml.Marshal(CustomersEnvelope{List: Customers})
		w.Header().Set("Content-Type", "application/xml")
	}

	//Wenn JSON angefordert oder nichts angegeben Transformiere DBQuery zu json
	if contentTypeRequest == "application/json" || contentTypeRequest != "application/xml" {
		data, err = json.Marshal(Customers)
		w.Header().Set("Content-Type", "application/json")
	}

	if err != nil {
		fehlerMessage = err.Error()
		http.Error(w, "Fehler beim Generieren der Nachricht", http.StatusInternalServerError)
		return
	}

	w.Write(data)

}
