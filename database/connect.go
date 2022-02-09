package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type AlarmDB struct {
	Id        uint       `json: id`
	DataAlarm string     `json: dataAlarm`
	TimeAlarm *time.Time `json: timeAlarm`
}

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "987123"
	dbname   = "postgres"
)

type Quer struct {
	DB *sql.DB
}

var DB Quer

func Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	DB.DB = db
	err = db.Ping()
	CheckError(err)

}

var DBT Quer

func ConnectTg(conDB chan Quer) {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	DBT.DB = db
	err = db.Ping()
	CheckError(err)

	conDB <- DBT

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
