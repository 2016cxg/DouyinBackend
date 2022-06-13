package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed() *FeedResponse {
	videolst := []Video{}
	lst := dao.NewVideoDAO().QueryLast30Videos()
	fmt.Println("feed 222")
	for i := range lst {
		id := lst[i].UserId
		title := lst[i].Title
		fcnt := lst[i].FavoriteCount
		ccnt := lst[i].CommentCount
		isf := lst[i].IsFavorite

		finalName := fmt.Sprintf("%d_%s", id, title)
		//saveFile := filepath.Join("./public/", finalName)

		user := dao.NewUserDao().QueryUserByID(int(id))
		suser := User{
			Id:            user.ID,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}

		tmp := Video{
			Id:            int64(i),
			Author:        suser,
			PlayUrl:       http + finalName,
			CoverUrl:      "",
			FavoriteCount: fcnt,
			CommentCount:  ccnt,
			IsFavorite:    isf,
		}
		videolst = append(videolst, tmp)
		fmt.Println(http + finalName)
	}

	fmt.Println("feed 333")

	return &FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videolst,
		NextTime:  time.Now().Unix(),
	}
}
