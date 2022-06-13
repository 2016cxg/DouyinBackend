package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

var http = "http://192.168.186.130:8080/resource/"

func Publish(token string, data *multipart.FileHeader, c *gin.Context) *Response {

	if _, exist := tokenmap[token]; exist {

	} else {
		return &Response{
			StatusCode: 1, StatusMsg: "User doesn't login",
		}
	}

	// save file
	filename := filepath.Base(data.Filename)
	userid := tokenmap[token][1]
	finalName := fmt.Sprintf("%d_%s", userid, filename)
	saveFile := filepath.Join("./public/", finalName)

	fmt.Println("save upload")
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		return &Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		}
	}

	// save file info into mysql
	dao.NewVideoDAO().Insert(&dao.Video{
		UserId: int64(userid),
		Title:  filename,
	})
	fmt.Println("write dao video")

	fmt.Println("save upload end")
	return &Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	}
}

func PublishList(token string, uid int64) *VideoListResponse {
	fmt.Println("111")
	if _, exist := tokenmap[token]; exist {
	} else {
		return &VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "User doesn't login",
			},
			VideoList: nil,
		}
	}

	fmt.Println("222")
	// Get all videos that user publishes
	videolst := []Video{}
	lst := dao.NewVideoDAO().QueryVideoByUID(uid)
	for i := range lst {
		id := lst[i].UserId
		title := lst[i].Title
		fcnt := lst[i].FavoriteCount
		ccnt := lst[i].CommentCount
		isf := lst[i].IsFavorite

		finalName := fmt.Sprintf("%d_%s", id, title)
		//saveFile := filepath.Join("./public/", finalName)

		user := dao.NewUserDao().QueryUserByID(int(uid))
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

	fmt.Println("333")

	return &VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videolst,
	}
}
