package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	dblayer "tiktok/db"
)

// token = username+password
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// RegisterHandler 注册账号
func RegisterHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 3, StatusMsg: "User already exist"},
		})
	} else {
		exist := dblayer.UserExist(username)
		if exist {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 4, StatusMsg: "UserName already exist"},
			})
		}
		userId := dblayer.AddUser(username, password)
		if userId == -100 {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 5, StatusMsg: "faile to register"},
			})
		} else {
			newUser := User{
				Id:   userId,
				Name: username,
			}
			usersLoginInfo[token] = newUser
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   userId,
				Token:    username + password,
			})
		}
	}
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

// UserInfoHandler 用户信息获取
func UserInfoHandler(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
