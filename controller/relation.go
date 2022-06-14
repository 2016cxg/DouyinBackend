package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")     //token
	tuid := c.Query("to_user_id") //to user
	tuid_, _ := strconv.Atoi(tuid)
	action := c.Query("action_type") // action type
	action_, _ := strconv.Atoi(action)

	resp := service.RelationAction(token, int64(tuid_), int64(action_))
	c.JSON(http.StatusOK, resp)
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {

	uid := c.Query("user_id")
	uid_, _ := strconv.Atoi(uid)

	resp := service.FollowList(int64(uid_))
	c.JSON(http.StatusOK, resp)
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	uid := c.Query("user_id")
	uid_, _ := strconv.Atoi(uid)

	resp := service.FollowerList(int64(uid_))
	c.JSON(http.StatusOK, resp)
}
