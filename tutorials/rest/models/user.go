package models

import (
  "time"
  "regexp"
  "golang.org/x/crypto/bcrypt"
)

type User struct{
  Id            int64   `json:"id"`
  Username      string  `json:"username"`
  Password      string  `json:"password"`
  Email         string  `json:"email"`
  createdDate   time.Time
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const userSchema string = `CREATE TABLE users(
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(30) NOT NULL UNIQUE, 
        password VARCHAR(64) NOT NULL,
        email VARCHAR(40),
        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

type Users []User

func NewUser(username, password, email string) *User{
  user := &User{Username: username, Email: email}
  user.SetPassword(password)
  return user
}

func CreateUser(username, password, email string) (*User, error){
  user := NewUser(username, password, email)
  err := user.Save()
  return user, err
}

func GetUser(id int) *User{
  user := NewUser("", "", "")
  sql := "SELECT id, username, password, email, created_date FROM users WHERE id=?"
  rows, _ := Query(sql, id)

  for rows.Next(){
    rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.createdDate)
  }
  return user
}

func GetUserByUsername(username string) *User{
  user := NewUser("", "", "")
  sql := "SELECT id, username, password, email, created_date FROM users WHERE username=?"
  rows, _ := Query(sql, username)

  for rows.Next(){
    rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.createdDate)
  }
  return user
}

func GetUsers() Users {
  sql := "SELECT id, username, password, email, created_date FROM users"
  users := Users{}
  rows, _ := Query(sql)

  for rows.Next(){
    user := User{}
    rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.createdDate)
    users = append(users, user)
  }
  
  return users
}

func Login(username, password string) bool{
  user := GetUserByUsername(username)
  err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
  return err == nil
}

func ValidEmail(email string) bool{
  return emailRegexp.MatchString(email)
}

func (this *User) Save() error {
  if this.Id == 0{
    return this.insert()
  }else{
    return this.update()
  }
}

func (this *User) insert() error {
  sql := "INSERT users SET username=?, password=?, email=?"
  id, err := InsertData(sql, this.Username, this.Password, this.Email)
  this.Id = id
  return err
}

func (this *User) update() error {
  sql := "UPDATE users SET username=?, password=?, email=?"
  _, err := Exec(sql, this.Username, this.Password, this.Email )
  return err
}

func (this *User) Delete() error {
  sql := "DELETE FROM users WHERE id=?"
  _, err := Exec(sql, this.Id)
  return err
}

func (this *User) SetPassword(password string){
  hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  this.Password = string(hash)
}

func (this *User) GetCreatedDate() time.Time{
  return this.createdDate
}

