package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type FeedResponse struct {
//	Response
//	VideoList []Video `json:"video_list,omitempty"`
//	NextTime  int64   `json:"next_time,omitempty"`
//}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	resp := service.Feed()
	c.JSON(http.StatusOK, *resp)
	fmt.Println(resp)

	//c.JSON(http.StatusOK, FeedResponse{
	//	Response:  Response{StatusCode: 0},
	//	VideoList: DemoVideos,
	//	NextTime:  time.Now().Unix(),
	//})
}
