package orm

import "time"

type User struct{ 
	Id			     int64	 `json:"id"`
	Username 	   string	 `json:"username"`
	Password 	   string	 `json:"password"`
  Email        string  `json:"email"`
  CreatedAt    time.Time
}

type Users []User

func CreateUser(username, password, email string) *User{
  user := NewUser(username, password, email)
  user.Save()
  return user
}

func NewUser(username, password, email string) *User{
  user := &User{Username: username, Password: password, Email: email}
  return user
}

func GetUser(id int) *User{
  user := &User{}
  db.Where("id=?", id).First(user)
  return user
}

func GetUsers() Users {
  users := Users{}
  db.Find(users)
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
  db.Create(&this)
}

func (this *User) update(){
  user := User{ Username: this.Username, Password: this.Password, Email : this.Email}
  db.Model(&this).UpdateColumns(user) 
}

func (this *User) Delete(){
  db.Delete(&this)
}


