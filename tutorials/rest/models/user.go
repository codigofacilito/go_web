package models

type User struct{
  Id        int64   `json:"id"`
  Username  string  `json:"username"`
  Password  string  `json:"password"`
  Email     string  `json:"email"`
}

const userSchema string = `CREATE TABLE users(
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(30) NOT NULL UNIQUE, 
        password VARCHAR(64) NOT NULL,
        email VARCHAR(40),
        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

type Users []User

func NewUser(username, password, email string) *User{
  user := &User{Username: username, Password: password, Email: email}
  return user
}

func CreateUser(username, password, email string) (*User, error){
  user := NewUser(username, password, email)
  err := user.Save()
  return user, err
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
