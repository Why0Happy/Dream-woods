package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var db *sql.DB
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
	router := gin.Default()
	router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/bili2/template/*.html"))

	router.GET("bili",Homepage)
	router.POST("register", Register)
	router.POST("/login", Login)

	router.StaticFS("data",http.Dir("./data"))
	router.Run(":8080")
}

func Homepage(context *gin.Context){
	println("----homepage start----")

	context.Header("Content-Type", "text/html; charset=utf-8")
	context.HTML(200,"Untitled-1.html",gin.H{})
}

func Register(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	repassword := context.PostForm("password-repeat")

	if password!=repassword {
		println("错误")
	}
	fmt.Println("user:" + username + password)
	if UserRecord(username, password) {
		context.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
	} else {
		context.JSON(500, gin.H{"status": http.StatusInternalServerError, "message": "数据库Insert报错"})
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