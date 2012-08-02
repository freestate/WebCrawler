package main

import (
  "github.com/ziutek/mymysql/mysql"
  _ "github.com/ziutek/mymysql/thrsafe"
//  "os"
//  "fmt"
)


func NewMysqlConnection() (mysql.Conn) {
  db := mysql.New( "tcp", "", "home.ndirt.com:3306", "crawl", "passu", "crawl" )

  err := db.Connect()
  if err != nil {
    panic( err )
  }

  return db
}

