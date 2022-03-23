package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main() {
	initDb()
	defer db.Close()
	http.HandleFunc("/", testFunc)
	http.HandleFunc("/api/events", getEvents)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is home")
}

type EventSummary struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Organizer string `json:"organizer"`
	Date      string `json:"date"`
}

type Events struct {
	events []EventSummary
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	eves := Events{}

	err := queryRepos(&eves)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(eves)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))

}

func queryRepos(eves *Events) error {
	rows, err := db.Query(`select * from Events`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		eve := EventSummary{}
		err = rows.Scan(
			&eve.ID,
			&eve.Name,
			&eve.Location,
			&eve.Organizer,
			&eve.Date)

		if err != nil {
			return err
		}

		eves.events = append(eves.events, eve)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func initDb() {
	config := dbConfig()

	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to DB")
}

func dbConfig() map[string]string {

	conf := make(map[string]string)

	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}

	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}

	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}

	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}

	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}

	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name

	return conf

}
