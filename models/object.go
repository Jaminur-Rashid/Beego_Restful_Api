package models

import (
	"errors"
	"fmt"
)

var (
	Objects map[string]*Object
)

type Object struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Phone       string `json:"Phone"`
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	DateOfBirth string `json:"DoB"`
}

func init() {
	Objects = make(map[string]*Object)

}

func AddNewUser(object Object) (ObjectId string) {
	Objects[object.Email] = &object
	return object.Email
}

func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAll() map[string]*Object {
	return Objects
}

/*
func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}
*/

func CheckError(err error) {

	if err != nil {
		fmt.Println("Database error is ", err)

		panic(err)

	}

}
