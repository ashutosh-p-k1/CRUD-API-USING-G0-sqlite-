package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//var db *sql.DB
var schema = `
CREATE TABLE IF NOT EXISTS policy(
id INTEGER PRIMARY KEY AUTOINCREMENT, 
apiversion varchar(20), 
kind varchar(20), 
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

	statement.Exec(m.APIVersion, m.Kind, m.GetMetadata(), m.GetSpec())
}

type Policyresult struct {
	ID         string `json:"id"`
	APIVersion string `json:"apiversion"`
	Kind       string `json:"kind"`
	Metadata   string `json:"metadata"`
	Spec       string `json:"spec"`
}

func Readdatabase(db *sql.DB) []Policyresult {
	fmt.Println("INSIDE READ :=>")
	newpolicy := []Policyresult{}
	statement, err := db.Query("Select * from policy")
	if err != nil {
		log.Fatal(err)
	}

	for statement.Next() {
		temp := Policyresult{}
		statement.Scan(&temp.ID, &temp.APIVersion, &temp.Kind, &temp.Metadata, &temp.Spec)

		newpolicy = append(newpolicy, temp)
	}
	return newpolicy
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
	statement.Exec(m.APIVersion, m.Kind, m.GetMetadata(), m.GetSpec(), id)

}

/*get Policy Spec in String format*/
func (p *Policy) GetSpec() string {
	data, _ := json.Marshal(p.Spec)

	return string(data)
}

/*get Policy Metadata in String format*/
func (p *Policy) GetMetadata() string {
	data, _ := json.Marshal(p.Metadata)
	return string(data)
}
