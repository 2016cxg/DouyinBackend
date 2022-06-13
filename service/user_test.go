package service

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	user := Login("abc", "123")
	fmt.Println(user)
}

func TestLogin1(t *testing.T) {
	user := Login("abc", "abc")
	fmt.Println(user)
}
