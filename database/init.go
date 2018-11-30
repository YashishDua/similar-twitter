package database

import (
  "log"
  "time"
  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() {
  var err error
  DB, err = sqlx.Open("postgres", "user=yashishdua dbname=postman-twitter sslmode=disable")
  if err != nil {
      log.Fatalln(err)
  }
  log.Print("Successfully connected to DB!")

	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(0)
	DB.SetConnMaxLifetime(time.Nanosecond)
}
