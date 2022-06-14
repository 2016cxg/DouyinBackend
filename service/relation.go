package service

import (
	"github.com/RaymondCode/simple-demo/dao"
)

//follow action
//query record (fuid, tuid) or (tuid, fuid)													--select from relation
//if (fuid, tuid) is not null
//	if record is (fuid, tuid, relation=0), revise relation=1								--update relation
// 	if record is (fuid, tuid, relation=1), return "Already follow" error
//	if record is (fuid, tuid, relation=2), revise relation=3, get them mutual follow
//	if record is (fuid, tuid, relation=3), return "Already follow" error
//if (tuid, fuid) is not null
//	if record is (tuid, fuid, relation=0), revise relation=2
//	if record is (tuid, fuid, relation=1), revise relation=3
//	if record is (tuid, fuid, relation=2), return "Aready follow" error
//	if record is (tuid, fuid, realtion=3), return "Aready follow" error
//if record (fuid, tuid) and (tuid, fuid) are all null
//	insert into relation table (fuid, tuid, relation=1)					--insert into relation

var NONFOLLOW = int64(0)
var FFOLLOWT = int64(1) //fuid follow tuid
var TFOLLOWF = int64(2) //tuid follow fuid
var MUTUALFOLLOW = int64(3)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

func RelationAction(token string, tuid int64, action int64) *Response {
	if _, exist := tokenmap[token]; exist {
	} else {
		return &Response{
			StatusCode: 1, StatusMsg: "User doesn't login",
		}
	}
	fuid := int64(tokenmap[token][1])

	if action == 1 {
		r1 := dao.NewRelationDao().SelFuidTuid(fuid, tuid)
		r2 := dao.NewRelationDao().SelFuidTuid(tuid, fuid)
		if r1 != nil {
			if r1.Relation == NONFOLLOW || r1.Relation == TFOLLOWF {
				dao.NewRelationDao().UpRelation(fuid, tuid, r1.Relation+1)
			} else if r1.Relation == FFOLLOWT || r1.Relation == MUTUALFOLLOW {
				return &Response{StatusCode: 1, StatusMsg: "Already follow"}
			}
		} else if r2 != nil {
			if r2.Relation == NONFOLLOW || r2.Relation == FFOLLOWT {
				dao.NewRelationDao().UpRelation(tuid, fuid, r2.Relation+2)
			} else if r2.Relation == TFOLLOWF || r2.Relation == MUTUALFOLLOW {
				return &Response{StatusCode: 1, StatusMsg: "Already follow"}
			}
		} else {
			//no record in relation table, either (fuid, tuid) or (tuid, fuid)
			dao.NewRelationDao().Insert(dao.Relation{
				Fuid:     fuid,
				Tuid:     tuid,
				Relation: FFOLLOWT,
			})
		}
		return &Response{StatusCode: 0, StatusMsg: "Fullow succeed"}
	} else if action == 2 {
		r1 := dao.NewRelationDao().SelFuidTuid(fuid, tuid)
		r2 := dao.NewRelationDao().SelFuidTuid(tuid, fuid)
		if r1 != nil {
			if r1.Relation == NONFOLLOW || r1.Relation == TFOLLOWF {
				return &Response{StatusCode: 1, StatusMsg: "Not even follow error"}
			} else if r1.Relation == FFOLLOWT || r1.Relation == MUTUALFOLLOW {
				dao.NewRelationDao().UpRelation(fuid, tuid, r1.Relation-1)
				return &Response{StatusCode: 0, StatusMsg: "Unfollow succeed"}
			}
		} else if r2 != nil {
			if r2.Relation == NONFOLLOW || r2.Relation == FFOLLOWT {
				return &Response{StatusCode: 1, StatusMsg: "Not even follow error"}
			} else if r2.Relation == TFOLLOWF || r2.Relation == MUTUALFOLLOW {
				dao.NewRelationDao().UpRelation(tuid, fuid, r2.Relation-2)
				return &Response{StatusCode: 0, StatusMsg: "Unfollow succeed"}
			}
		} else {
			return &Response{StatusCode: 1, StatusMsg: "Not even follow error"}
		}
	} else {
		return &Response{StatusCode: 1, StatusMsg: "Invalid action Error"}
	}
	return nil
}

func FollowList(uid int64) *UserListResponse {

	followLst := dao.NewRelationDao().SelFollowLst(uid)
	userLst := []User{}
	for i := range followLst {
		uid := followLst[i]
		user := dao.NewUserDao().QueryUserByID(int(uid))
		sUser := User{
			Id:            user.ID,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
		}
		userLst = append(userLst, sUser)
	}
	return &UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userLst,
	}
}

func FollowerList(uid int64) *UserListResponse {

	followerLst := dao.NewRelationDao().SelFollowerLst(uid)
	userLst := []User{}
	for i := range followerLst {
		uid := followerLst[i]
		user := dao.NewUserDao().QueryUserByID(int(uid))
		sUser := User{
			Id:            user.ID,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
		}
		userLst = append(userLst, sUser)
	}
	return &UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userLst,
	}
}
