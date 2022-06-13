package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	fmt.Printf("abcdef")
	token := c.PostForm("token")

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	resp := service.Publish(token, data, c)

	c.JSON(http.StatusOK, *resp)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	uid := c.Query("user_id")
	token := c.Query("token")

	uid_, _ := strconv.Atoi(uid)
	resp := service.PublishList(token, int64(uid_))

	c.JSON(http.StatusOK, *resp)
}
