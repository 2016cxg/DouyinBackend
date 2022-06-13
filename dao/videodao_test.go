package dao

import (
	"fmt"
	"testing"
)

//func TestVideoDAO_Insert(t *testing.T) {
//	NewVideoDAO().Insert(&Video{
//		UserId: 1,
//		Title: "cef",
//	})
//}

func TestVideoDAO_QueryVideoByUID(t *testing.T) {
	lst := NewVideoDAO().QueryVideoByUID(1)

	//fmt.Printf((*lst)[0].UserId, (*lst)[0].Title)
	for i := range lst {
		fmt.Println(i)
	}
	fmt.Println(lst[0].UserId)

	fmt.Println(lst)
}
