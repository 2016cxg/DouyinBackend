package dao

import (
	"fmt"
	"testing"
)

//func TestUserDAO_insert(t *testing.T) {
//	u := User{
//		Username:"tizi365",
//		Password:"123456",
//	}
//
//	NewUserDao().insert(&u)
//}

//func TestUserDAO_queryUserByName(t *testing.T) {
//
//	username := "abc"
//	user := NewUserDao().QueryUserByName(username)
//	fmt.Println(user)
//}

func TestUserDAO_queryUserByID(t *testing.T) {
	fmt.Println("asdf")
	id := 4
	user := NewUserDao().QueryUserByID(id)
	fmt.Println(user)
}
