package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

var tokenmap = make(map[string][]int)

func Register(username string, password string) *UserLoginResponse {

	fmt.Println(username, password)
	user := dao.NewUserDao().QueryUserByName(username)

	fmt.Println(user)
	if user != nil {
		return &UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		}
	}

	dao.NewUserDao().Insert(&dao.User{
		Username: username,
		Password: password,
	})

	user = dao.NewUserDao().QueryUserByName(username)

	token := username + password

	// Add token in tokenmap, this means that user is active
	tokenmap[token] = []int{10, int(user.ID)}

	return &UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   user.ID,
		Token:    token,
	}

}

func Login(username string, password string) *UserLoginResponse {

	fmt.Println(username, password)
	user := dao.NewUserDao().QueryUserByName(username)

	fmt.Println(user)

	//if user is nil, then user doen,t exist
	if user != nil {
		if user.Password == password {
			token := username + password

			// Repeatablly login, return false
			if lst, exist := tokenmap[token]; exist {
				if lst[0] > 0 {
					return &UserLoginResponse{
						Response: Response{StatusCode: 1, StatusMsg: "Already in, repeatable login"},
					}
				}

			}

			// Add token in tokenmap, this means that user is active
			tokenmap[token] = []int{10, int(user.ID)}
			return &UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.ID,
				Token:    token,
			}
		}
	}

	// User doesn't exist
	return &UserLoginResponse{
		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	}

}

func GetUserInfo(token string) *UserResponse {
	// Check user permission, whether user is logined
	if lst, exist := tokenmap[token]; exist {
		user := dao.NewUserDao().QueryUserByID(lst[1])
		return &UserResponse{
			Response: Response{StatusCode: 0},
			User: User{
				Id:            user.ID,
				Name:          user.Username,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
			},
		}
	} else {
		return &UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't login"},
		}
	}
}
