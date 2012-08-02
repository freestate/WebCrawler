package main

import (
//  "github.com/ziutek/mymysql/mysql"
  _ "github.com/ziutek/mymysql/thrsafe"
)


type Settings struct {
  settingsUpdateInterval int
}


func NewSettings() (settings Settings) {
  settings.settingsUpdateInterval = 30 // Update settings every 30 seconds
  settings.Load()
  return settings  
}



func (s *Settings) Load() {
  db := NewMysqlConnection()
  defer db.Close()

  rows, res, err := db.Query( "select * from settings" )
  if err != nil {
    panic( err )
  }

  keyCol   := res.Map( "key" )
  valueCol := res.Map( "value" )

  for _, row := range rows {
    key := row.Str( keyCol )
    switch( key ) {
      case "settingsUpdateInterval":
        s.settingsUpdateInterval = row.Int( valueCol )
        break
    }
  }
}


