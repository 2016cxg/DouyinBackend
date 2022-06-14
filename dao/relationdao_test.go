package dao

import (
	"fmt"
	"testing"
)

//func TestRelationDAO_Insert(t *testing.T) {
//	NewRelationDao().Insert(Relation{Fuid: 1, Tuid: 2, Relation: 1})
//}
//
//func TestRelationDAO_SelFuidTuid_1(t *testing.T) {
//	relation := NewRelationDao().SelFuidTuid(1,2)
//
//	fmt.Println(relation)
//}
//
//func TestRelationDAO_UpRelation(t *testing.T) {
//	NewRelationDao().UpRelation(1,2,3)
//}
//
//func TestRelationDAO_SelFuidTuid_2(t *testing.T) {
//	relation := NewRelationDao().SelFuidTuid(1,2)
//
//	fmt.Println(relation)
//}

func TestSelFollowLst(t *testing.T) {
	u := NewRelationDao().SelFollowLst(2)
	fmt.Println(u)
}
