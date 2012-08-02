package main

import (
//  "github.com/Karethoth/NDParser/parser"
//  "github.com/ziutek/mymysql/mysql"
  _ "github.com/ziutek/mymysql/thrsafe"
//  "os"
    "fmt"
)


func main() {
  settings := NewSettings()
  fmt.Printf( "interval=%d\n", settings.settingsUpdateInterval )
}

