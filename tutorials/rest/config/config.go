package config

import "github.com/eduardogpg/gonv"
import "fmt"

type DatabaseConfig struct{
  Username string
  Password string
  Host string
  Port int
  Database string
  Debug bool
}

var database *DatabaseConfig

func init(){
  database = &DatabaseConfig{}
  database.Username = gonv.GetStringEnv("USERNAME", "root")
  database.Password = gonv.GetStringEnv("PASSWORD", "")
  database.Host = gonv.GetStringEnv("HOST", "localhost")
  database.Port = gonv.GetIntEnv("PORT", 3306)
  database.Database = gonv.GetStringEnv("DATABASE", "project_go_web")
  database.Debug = gonv.GetBoolEnv("DEBUG", true)
}

func GetDebug() bool{
  return database.Debug
}

func GetUrlDatabase() string {
  return database.url()
}

//&parseTime=true
func (this *DatabaseConfig) url() string{
  return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", this.Username, this.Password, this.Host, this.Port, this.Database)
}








