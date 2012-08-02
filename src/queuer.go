package main

import (
//  "github.com/Karethoth/NDParser/parser"
    "github.com/ziutek/mymysql/mysql"
  _ "github.com/ziutek/mymysql/thrsafe"
//  "os"
    "fmt"
    "time"
)



func Queuer( settings *Settings ) {
  db := NewMysqlConnection()
  defer db.Close()

  for true {
    time.Sleep( 10000 * time.Millisecond )

    if !IsQueueSmall( &db, settings ) {
      continue;
    }

    fmt.Println( "Queue is small! Updating!" )
    GenerateQueue( &db, settings )
  }
}



func IsQueueSmall( db *mysql.Conn, settings *Settings ) bool {
  rows, _, err := (*db).Query( "select id from queue" )
  if err != nil {
    panic( err )
  }

  if len( rows ) <= settings.smallQueueLimit {
    return true
  }
  return false
}



func GenerateQueue( db *mysql.Conn, settings *Settings ) {
  rows, res, err := (*db).Query( "select id from page where inQueue = '0' and timestamp <= 'DATE_SUB(CURDATE(), INTERVAL 10 HOUR)' limit 50" )
  if err != nil {
    panic( err )
  }

  if len( rows ) <= 0 {
    fmt.Println( "No pages to queue!" )
    return
  }

  idCol := res.Map( "id" )

  for _, row := range rows {
    id := row.Str(idCol)

    _, _, err = (*db).Query( "update page set inQueue = 1 where id = '"+id+"' limit 1" )
    if err != nil {
      panic( err )
    }

    _, _, err = (*db).Query( "insert into queue (pageId) values ('"+id+"')" )
    if err != nil {
      panic( err )
    }
  }
}
