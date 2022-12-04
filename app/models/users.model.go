package models

import (
	"encoding/json"
	"errors"

	"github.com/ilies-a/go-googleauth-example/app/utils"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
}

var user1Id = utils.GenerateStringUUID()
var user2Id = utils.GenerateStringUUID()

var jsonUsers string = `{
	"` + user1Id + `":
	{
		"id":"` + user1Id + `",
		"name":"john",
		"email":"john@domain.com"
	},
	"` + user2Id + `":
	{
		"id":"` + user2Id + `",
		"name":"doe",
		"email":"doe@domain.com"
	}
}`

var Users = make(map[string]User)

func GetAllUsers() map[string]User {
	return Users
}

func GetUser(userId string) (User, error) {
	if result, ok := Users[userId]; !ok {
		return result, errors.New("404")
	}

	return Users[userId], nil
}

func SaveUser(user *User) {
	Users[user.Id] = *user
}

func DeleteUser(userId string) error {
	if _, ok := Users[userId]; !ok {
		return errors.New("404")
	}

	delete(Users, userId)
	return nil
}

func InitFakeDb() {
	err := json.Unmarshal([]byte(jsonUsers), &Users)
	if err != nil {
		panic("FillFakeDb unmarshal error")
	}
}
