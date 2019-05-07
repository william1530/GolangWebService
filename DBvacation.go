package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

//Country object from chinook db
type Country struct {
	CountryName   string         `db:"CountryName"`
	Hyperlink     string         `db:"Hyperlink" json:",omitempty"`
	CurrencyName  sql.NullString `db:"CurrencyName" json:",omitempty"`
	ContinentName string         `db:"ContinentName"`
	LengthKM      sql.NullInt64  `db:"Length_KM" json:",omitempty"`
	BorderNumber  sql.NullInt64  `db:"cnt_border"`
}

//CountriesEnvelope ist ein Kontainer für die Country Entität, wird benötigt für XML marshalling
type CountriesEnvelope struct {
	XMLName xml.Name  `xml:"Countries"`
	List    []Country `xml:"Country"`
}

//Handler2 makes a select on the db
func tstVacation(w http.ResponseWriter, r *http.Request) {

	var err error
	var data []byte
	var fehlerMessage string

	contentTypeRequest := r.Header.Get("Content-type")

	Countries := []Country{}

	fmt.Print("executing query")

	//Select data and write it to variable
	err = vacationdb.Select(&Countries, `select country.name CountryName,ifnull(Hyperlink,'') Hyperlink,continent.Name ContinentName, currency.Name CurrencyName, sum(border.Length_KM) Length_KM, count(distinct border.country2) cnt_border
										from country
										left join currency on country.CurrencyCode = currency.Code
										left join continent on ContinentCode = continent.Code
										join border on country.ISO_A3_Code = border.Country1
										group by country.name ,Hyperlink ,continent.Name,currency.Name `)

	if err != nil {
		fehlerMessage = err.Error()
		http.Error(w, "Datenbankfehler: "+fehlerMessage, http.StatusInternalServerError)
		return
	}

	//Wenn XML angefordert Transformiere DBQuery zu XML
	if contentTypeRequest == "application/xml" || true {
		data, err = xml.Marshal(CountriesEnvelope{List: Countries})
		w.Header().Set("Content-Type", "application/xml")
	}

	//Wenn JSON angefordert oder nichts angegeben Transformiere DBQuery zu json
	if contentTypeRequest == "application/json" || contentTypeRequest != "application/xml" {
		data, err = json.Marshal(Countries)
		w.Header().Set("Content-Type", "application/json")
	}

	if err != nil {
		fehlerMessage = err.Error()
		http.Error(w, "Fehler beim Generieren der Nachricht", http.StatusInternalServerError)
		return
	}

	w.Write(data)

}

//Handler2 makes a select on the db
func tstvacation(w http.ResponseWriter, r *http.Request) {
	//db, err := sql.Open("mysql", "rdr:supradin@tcp(192.168.1.107:3307)/Vacation")

	var err error

	//fmt.Print("calling init db function")
	//InitDB("none")

	fmt.Print("executing query")
	rows, err := vacationdb.Query("SELECT name,hyperlink FROM country where name in ('Switzerland','Sweden')")

	if err != nil {
		panic(err)
		//os.Exit(1)
		// log.Fatal(err)

	}

	for rows.Next() {
		var hyperlin string
		var name string
		rows.Scan(&name, &hyperlin)

		fmt.Println(name, "\t", hyperlin)
	}

	fmt.Fprintf(w, "Hi there, I love %s!", "snippets")

}
