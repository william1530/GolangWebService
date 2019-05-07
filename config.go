package main

import "database/sql"

type databaseC struct {
	name             string
	connectionString string
	typ              string
	dbObject         *sql.DB
}

var databaseConf = []databaseC{
	databaseC{
		name:             "Vacation",
		connectionString: "rdr:supradin@tcp(192.168.1.107:3307)/Vacation",
		typ:              "mysql",
	},
}
