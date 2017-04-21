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
const host string = "localhost"
const port int = 3306
const database string = "project_go_web"

func CreateConnection(){
  if connection, err := sql.Open("mysql", generateURL() ); err != nil{
    panic(err)
  }else{
    db = connection
  }
}

func CreateTables(){
  createTable("users", userSchema)
}

func createTable(tableName, schema string){
  if !existsTable(tableName){
    Exec(schema)
  }else{
    truncateTable(tableName)
  }
}

func truncateTable(tableName string){
  sql := fmt.Sprintf("TRUNCATE %s", tableName)
  Exec(sql)
}

func existsTable(tableName string) bool{
  sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
  rows, _ := Query(sql)
  return rows.Next()
}

func Exec(query string, args ...interface{})(sql.Result, error) {
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


func Ping(){
  if err := db.Ping(); err != nil{
    panic(err)
  }
}

func CloseConnection(){
  db.Close()
}

//<username>:<password>@tcp(<host>:<port>)/<database>
func generateURL() string {
  return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, database)
}

