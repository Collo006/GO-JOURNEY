package main

import (
	"database/sql"
	"log"
	"fmt"
	"os"
	"github.com/go-sql-driver/mysql"
)

//this is used to hold row data
type Album struct {
	ID int64
	Title string
	Artist string
	Price float32
}

//this is the db handle
var db *sql.DB

//func to give you db access 
func main() {
	//capture connection properties from MYSQL driver's Config and the type's FormatDSN.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	//get db handle.
	var err error
	//call sql.Open to initialize the db variable, passing the return value of FormatDSN
	//check for an error from sql.Open incase it fails, eg your db connection specifics weren't well-formed
	db, err = sql.Open("mysql",cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	//call DB.ping to confirm that connecting to the databse works
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Albums found: %v\n", albums)
}

//albumsByArtist queries for albums that have the specified artists name.
func albumsByArtist(name string) ([]Album, error) {
	//An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist =?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	//loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err !=nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err !=nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}