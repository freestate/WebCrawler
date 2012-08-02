package main

import (
//  "github.com/Karethoth/NDParser/parser"
  "github.com/ziutek/mymysql/mysql"
  _ "github.com/ziutek/mymysql/thrsafe"
//  "os"
//  "fmt"
)


func main() {
  db := mysql.New( "tcp", "", "home.ndirt.com:3306", "crawl", "passu", "crawl" )

  err := db.Connect()
  if err != nil {
    panic( err )
  }
  defer db.Close()

  _, _, err = db.Query( "select * from settings" )
  if err != nil {
    panic( err )
  }

}

