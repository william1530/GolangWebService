package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//database configuration
/*type database struct {
	name: string,
	db_type: string,
	conn_str: string
}

vacation := database{name:"vacation",db_type: "mysql",conn_str: "rdr:supradin@tcp(192.168.1.107:3307)/Vacation"}
*/

//username:password@protocol(address)/dbname?param=value
//db, err := sql.Open("mysql", "db_user:password@tcp(localhost:3306)/my_db")

//init database function

/*
gets input dbconnection string, Database name

connects to database and creates object
object can be accessed from handlers

*/

var vacationdb *sqlx.DB

var chinookdb *sqlx.DB

//InitDB opens Database connection for queries
func InitDB(dataSourceName string) {
	var err error

	fmt.Print("setting vacation db. (initdb function)")

	vacationdb, err = sqlx.Open("mysql", "rdr:supradin@tcp(192.168.1.107:3307)/Vacation")

	if err != nil {
		log.Panic(err)
	}

	if err = vacationdb.Ping(); err != nil {
		log.Panic(err)
	}

	fmt.Print("setting chinook db. (initdb function)")

	chinookdb, err = sqlx.Open("sqlite3", "C:\\Users\\ckw-mewi\\Documents\\golang\\WebService\\chinook\\chinook.db")
	if err != nil {
		log.Panic(err)
	}

	if err = chinookdb.Ping(); err != nil {
		log.Panic(err)
	}
}
