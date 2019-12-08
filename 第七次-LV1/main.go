package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("user:" + username + password)
	if UserRecord(username, password) {
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
	} else {
		c.JSON(500, gin.H{"status": http.StatusInternalServerError, "message": "数据库Insert报错"})
	}
}

func UserRecord(username string, password string) bool {
	stmt, err := db.Prepare(
		"insert into login(username,password)values(?,?) ")
	if err != nil {
		fmt.Println("fail to Prepare\n")
		log.Fatal(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, password)
	if err != nil {
		fmt.Println("fail to Exec\n")
		log.Fatal(err)
		return false
	}
	return true
}

func Login(u *gin.Context) {
	username := u.PostForm("username")
	password := u.PostForm("password")
	if UserCheck(username, password) {
		u.SetCookie("username", username, 10, "/", "localhost", false, true)
		u.JSON(200, gin.H{
			"status":  http.StatusOK,
			"hello":   username,
			"message": "登陆成功",
		})
	} else {
		u.JSON(500, gin.H{
			"status":  http.StatusForbidden,
			"message": "登陆失败，请检查用户名或密码是否错误",
		})
	}
}

func UserCheck(username string, password string) bool {
	stmt, err := db.Query("select password from login where username=?", username)
	defer stmt.Close()
	if err != nil {
		println("fail to Select\n")
		log.Fatal(err)
		return false
	}
	var password0 string
	for stmt.Next() {
		stmt.Scan(&password0)
		if err != nil {
			println("fail to Scan\n")
			log.Fatal(err)
		}
	}
	if password != password0 {
		return false
	} else {
		return true
	}
}

func init() {
	db1, err := sql.Open("mysql", "root:@tcp(localhost:3306)/user?charset=utf8")
	db = db1
	db.SetMaxOpenConns(1000)
	if err != nil {
		fmt.Println("fail to open\n")
		log.Fatal(err)
	}
}

func main() {
	u := gin.Default()
	u.GET("/hello", func(a *gin.Context) {
		a.JSON(200, gin.H{
			"hello": "guest",
		})
	})
	u.POST("/register", Register)
	u.POST("/login", Login)

	u.Run(":8080")
}
