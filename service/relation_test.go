package service

import (
	"fmt"
	"testing"
)

//func TestRelationAction(t *testing.T) {
//	res := RelationAction("abc", 1,2,1)
//	fmt.Println(res)
//}
//
//func TestRelationAction2(t *testing.T) {
//	res := RelationAction("abc",2,3,1)
//	fmt.Println(res)
//}
//func TestRelationAction3(t *testing.T) {
//	res := RelationAction("abc",2,3,1)
//	fmt.Println(res)
//}
//func TestRelationAction4(t *testing.T) {
//	res := RelationAction("abc",3,2,1)
//	fmt.Println(res)
//}
//func TestRelationAction5(t *testing.T) {
//	res := RelationAction("abc",3,2,2)
//	fmt.Println(res)
//}

func TestFollowList(t *testing.T) {
	resp := FollowList(2)

	fmt.Println(resp)
}
