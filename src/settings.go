package main

import (
//  "github.com/ziutek/mymysql/mysql"
  _ "github.com/ziutek/mymysql/thrsafe"
  "time"
)



type Settings struct {
  settingsUpdateInterval      int
  settingsUpdaterLoopInterval int
  maxCrawlerCount             int
  

  lastUpdate                  int64
}



func SettingsUpdater( s *Settings ){
  for true {
    if s.IsOld() {
      s.Load()
    }
    time.Sleep( time.Duration(s.settingsUpdaterLoopInterval) * time.Millisecond )
  }
}



func NewSettings() (settings Settings) {
  settings.settingsUpdateInterval = 30      // Update settings every 30 seconds
  settings.settingsUpdaterLoopInterval = 30 // Loop settings old check every 10 seconds
  settings.maxCrawlerCount = 1              // Max crawlers running at the same time
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

      case "settingsUpdaterLoopInterval":
        s.settingsUpdaterLoopInterval = row.Int( valueCol )
        break

      case "maxCrawlerCount":
        s.maxCrawlerCount = row.Int( valueCol )
        break
    }
  }

  s.lastUpdate = time.Now().Unix()
}



func (s *Settings) IsOld() bool {
  if time.Now().Unix() >= s.lastUpdate+int64(s.settingsUpdateInterval) {
    return true
  }
  return false
}

