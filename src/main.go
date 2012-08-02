package main

import (
//  "github.com/Karethoth/NDParser/parser"
//  "github.com/ziutek/mymysql/mysql"
  _ "github.com/ziutek/mymysql/thrsafe"
//  "os"
    "fmt"
    "time"
)



func main() {
  settings := NewSettings()
  go SettingsUpdater( &settings )
  go Queuer( &settings )
  for true {
    fmt.Printf( "interval    =%d\n", settings.settingsUpdateInterval )
    fmt.Printf( "loopInterval=%d\n", settings.settingsUpdaterLoopInterval )
    fmt.Printf( "maxCrawlers =%d\n", settings.maxCrawlerCount )
    time.Sleep( 1000 * time.Millisecond )
  }
}

