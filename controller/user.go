package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	resp := service.Register(username, password)

	c.JSON(http.StatusOK, *resp)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	resp := service.Login(username, password)

	c.JSON(http.StatusOK, *resp)
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	resp := service.GetUserInfo(token)

	c.JSON(http.StatusOK, *resp)
}
