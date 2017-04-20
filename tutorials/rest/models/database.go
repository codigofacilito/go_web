package models

import(
  "fmt"
  "log"

  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)
var db *sql.DB

const username string = "root"
const password string = ""
const database string = "go_web"
const host string = "localhost"
const port int = 3306

func CreateConnection(){
  url := generateURL()
  if connection, err := sql.Open("mysql", url); err != nil{
    panic(err)
  }else{
    db = connection
  }
}

func CreateTables(){
  createTable(userTable, userSchema)
}

func createTable(tableName, schema string){
  if !existsTable(tableName){
    Exec(schema)
  }
}

func existsTable(tableName string) bool{
  sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
  row, _ := db.Query(sql)
  return row.Next()
}

func Ping(){
  if err := db.Ping(); err != nil{
    panic(err)
  }
}

func CloseConnection(){
  db.Close()
}

func generateURL() string{
  return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, database)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
  result, err := db.Exec(query, args...)
  if err != nil{
    log.Println(err)
  }
  return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
  rows, err := db.Query(query, args...)
  if err != nil{
    log.Println(err)
  }
  return rows, err
}
