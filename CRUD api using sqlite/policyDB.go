package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//var db *sql.DB
var schema = `
CREATE TABLE IF NOT EXISTS policy( 
id INTEGER  PRIMARY KEY AUTOINCREMENT,
apiversion varchar(20) ,
kind varchar(30) ,
metadata  varchar(200),
spec varchar(2000)
);
`

func createDatabase() *sql.DB {
	db, _ := sql.Open("sqlite3", "./policyDB.db")
	statement, err := db.Prepare(schema)

	if err != nil {
		log.Println("failed to create Database", err)
	}
	statement.Exec()
	return db

}

func insertDatabase(m Policy, db *sql.DB) {

	statement, err := db.Prepare("insert into policy (apiversion,kind,metadata,spec) values (?,?,?,?)")
	if err != nil {
		log.Println("Error in insertion :=>", err)

	}

	statement.Exec(m.APIVersion, m.Kind, m.Metadata, m.Spec)
}

func Readdatabase(db *sql.DB) []Policy {
	fmt.Println("INSIDE READ :=>")
	movies := []Policy{}
	statement, err := db.Query("Select * from policy")
	if err != nil {
		log.Fatal(err)
	}

	var apiversion string
	var kind string
	var metadata string
	var spec string
	for statement.Next() {
		statement.Scan(&apiversion, &kind, &metadata, &spec)
		movie := Policy{APIVersion: apiversion, Kind: kind, Metadata: metadata, Spec: spec}
		movies = append(movies, movie)
	}
	return movies
}
func DeleteDB(id string, db *sql.DB) {
	statement, err := db.Prepare("Delete from policy where id = ?")
	if err != nil {
		log.Println("Error in Deleting the database", err)
	}
	statement.Exec(id)
}

func UpdateDB(m Policy, id string, db *sql.DB) {
	statement, err := db.Prepare("Update policy set apiversion=?,kind=?,metadata=? , spec=? where id=?")
	if err != nil {
		log.Println("Error in Updation :", err)
	}
	statement.Exec(m.APIVersion, m.Kind, m.Metadata, m.Spec, id)

}
