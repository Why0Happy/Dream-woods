package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	u := gin.Default()
	u.GET("/hello", func(a *gin.Context) {
		a.JSON(200, gin.H{
			"hello": "guest",
		})
	})
	u.POST("/username/password", Register)
	u.POST("/login/username/password", Login)

	u.Run(":8080")
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("user:" + username + password)
	c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
}

func Login(u *gin.Context) {
	username := u.PostForm("username")
	password := u.PostForm("password")
	u.SetCookie("username", username, 10, "/", "localhost", false, true)
	u.JSON(200, gin.H{
		"status":  http.StatusOK,
		"hello":   username,
		"your":    password,
		"message": "登陆成功",
	})
}
