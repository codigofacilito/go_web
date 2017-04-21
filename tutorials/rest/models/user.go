package models

// import "errors"

type User struct{
	Id			  int64		`json:"id"`
	Username 	string	`json:"username"`
	Password 	string	`json:"password"`
  Email     string  `json:"email"`
}

const userSchema string = `CREATE TABLE users(
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(30) NOT NULL, 
        password VARCHAR(64) NOT NULL,
        email VARCHAR(40),
        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

type Users []User


func NewUser(username, password, email string) *User{
  user := &User{Username: username, Password: password, Email: email}
  return user
}

func CreateUser(username, password, email string) *User{
  user := NewUser(username, password, email)
  user.Save()
  return user
}

func GetUser(id int) *User{
  user := NewUser("", "", "")
  sql := "SELECT id, username, password, email FROM users WHERE id=?"
  rows, _ := Query(sql, id)

  for rows.Next(){
    rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
  }
  return user
}

func GetUsers() Users {
  sql := "SELECT id, username, password, email FROM users"
  users := Users{}
  rows, _ := Query(sql)

  for rows.Next(){
    user := User{}
    rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
    users = append(users, user)
  }
  
  return users
}

func (this *User) Save(){
  if this.Id == 0{
    this.insert()
  }else{
    this.update()
  }
}

func (this *User) insert(){
  sql := "INSERT users SET username=?, password=?, email=?"
  result, _ := Exec(sql, this.Username, this.Password, this.Email)
  this.Id, _ = result.LastInsertId() //int64
}

func (this *User) update(){
  sql := "UPDATE users SET username=?, password=?, email=?"
  Exec(sql, this.Username, this.Password, this.Email )
}

func (this *User) Delete(){
  sql := "DELETE FROM users WHERE id=?"
  Exec(sql, this.Id)
}















