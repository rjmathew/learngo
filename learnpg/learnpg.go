package main

import (
"database/sql"
_"github.com/lib/pq"
	"log"
	"fmt"
)
const (
	DB_USER     = "user"
	DB_PASSWORD = "pwd"
	DB_NAME     = "ensemble"
	DB_HOST		= "localhost"
	DB_SCHEMA 	= "showbiz_xform"
)
type tableInfo struct {
	name string
	hasRows bool
}

func NewTableInfo(name string, hasRows bool) *tableInfo {
m := new(tableInfo)
m.name = name
m.hasRows = hasRows
return m
}

func main() {

	ch := make(chan tableInfo)

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s",
		DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
	db,err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	rows, err := db.Query("SELECT table_name FROM information_schema.tables where table_schema='" + DB_SCHEMA + "'")

	counter := 0
	for rows.Next() {
		counter ++
		var table_name string
		err = rows.Scan(&table_name)
		go checkTable(table_name, ch, db)
	}
	iter := 1
	for iter <= counter {
		tableValue := <- ch
		if (!tableValue.hasRows) {
			fmt.Println(tableValue)
		}
		iter ++

	}

}

func checkTable(tableName string, ch chan tableInfo, db *sql.DB) {
	rows, err := db.Query("SELECT count(*) rowcount FROM (select 1 from " + DB_SCHEMA + "." + tableName + "  limit 1) as t")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var rowcount int
		err = rows.Scan(&rowcount)
		if rowcount == 0 {
			ch <- *NewTableInfo(tableName, false)
		} else {
			ch <- *NewTableInfo(tableName, true)
		}
	}
}
