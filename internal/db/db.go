package db

import (
	"database/sql"
	"fmt"
	"log"
 _"github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
  var err error
  dbHost := "localhost"
  dbPort := "5432"
  dbUser := "postgres"
  dbPassword := "postgres"
  dbName := "twitterClone"

  connStr:=fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

  DB, err = sql.Open("postgres", connStr)
  if err != nil {
    log.Fatalf("failed to connect to db %v", err);
  }
  err = DB.Ping()
  if err != nil {
log.Fatalf("failed to ping db %v", err);
  }
    log.Println("connected to db");
  createUserTable()
} 

func createUserTable(){
    query:=`CREATE TABLE IF NOT EXISTS users(
      id SERIAL PRIMARY KEY,
      username VARCHAR(50) UNIQUE NOT NULL,
      password VARCHAR(50) NOT NULL
      );`


  _,err:=DB.Exec(query)
  if err!=nil{
    log.Fatalf("failed to create table %v",err)
  }

    

}
