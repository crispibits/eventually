package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"eventually/model"

	_ "github.com/glebarez/go-sqlite"
)

var db *sql.DB

func main() {
	// connect
	var err error
	db, err = sql.Open("sqlite", "objects/events.db")
	if err != nil {
		log.Fatal(err)
	}

	// get SQLite version
	row := db.QueryRow("select sqlite_version()")
	var version string
	row.Scan(&version)
	fmt.Println(version)
	createDB()
	row = db.QueryRow("select display_name from cost_type where name = 'STANDARD'")
	var cost string
	row.Scan(&cost)
	fmt.Println(cost)
	createEvents()
	createLocations()
}

func createDB() {
	fmt.Println("Creating DB")
	schema, err := os.ReadFile("schema.sql")
	fmt.Println(string(schema))
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal(err)
	}
}

func createEvents() {
	events, err := model.LoadEvents("objects/events.json")
	if err != nil {
		log.Fatal(err)
	}
	eventStmt, err := db.Prepare("INSERT INTO event(title,content,costs,times,location,owner,tags,ticketURLs,restrictions) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	var result sql.Result
	for _, event := range events {
		fmt.Println(event.Title)
		costs, _ := json.Marshal(event.Costs)
		times, _ := json.Marshal(event.Times)
		tags, _ := json.Marshal(event.Tags)
		ticketURLs, _ := json.Marshal(event.TicketURLs)
		restrictions, _ := json.Marshal(event.Restrictions)
		result, err = eventStmt.Exec(event.Title, event.Content, string(costs), string(times), event.Location, event.Owner, string(tags), string(ticketURLs), string(restrictions))
		if err != nil {
			log.Fatal(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
	}
}

func createLocations() {
	locations, err := model.LoadLocations("objects/locations.json")
	if err != nil {
		log.Fatal(err)
	}
	eventStmt, err := db.Prepare("INSERT INTO location(name,display_name,address,longitude,latitude,url) values(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	var result sql.Result
	for _, location := range locations {
		fmt.Println(location.Name)
		result, err = eventStmt.Exec(location.Name, location.DisplayName, location.Address, location.Longitude, location.Latitude, location.URL)
		if err != nil {
			log.Fatal(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
	}
}

func getEvents() {
	rows, err := db.Query("select costs from event")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
		var costs []model.Cost
		json.Unmarshal([]byte(name), &costs)
		for _, cost := range costs {
			fmt.Println(cost.TicketType)
		}
	}
}
