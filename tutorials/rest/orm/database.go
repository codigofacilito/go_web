package orm

import(
  "../config"

  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)
var db *gorm.DB

func CreateConnection(){
  url := config.GetUrlDatabase()
  if connection, err := gorm.Open("mysql", url ); err != nil{
    panic(err)
  }else{
    db = connection
  }
}

func CreateTables(){
  db.DropTableIfExists(&User{})
  db.CreateTable(&User{})
}

func CloseConnection(){
  db.Close()
}