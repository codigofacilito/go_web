package models

import "errors"

type User struct{
	Id			int		`json:"id"`
	Username 	string	`json:"username"`
	Password 	string	`json:"password"`
}

var users = make(map[int]User)

func SetDefaultUser(){
	user := User{ Id: 1, Username: "eduardo_gpg", Password: "password123" }
	users[user.Id] = user
}

func GetUser(userId int) (User, error){
	if user, ok := users[userId]; ok{
		return user, nil
	}
	return User{}, errors.New("El usuario no se encuentra dentro del map")
}

