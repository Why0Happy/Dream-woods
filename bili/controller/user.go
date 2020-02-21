package controller

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var db *sql.DB
func init() {
	db1, err := sql.Open("mysql", "root:@tcp(localhost:3306)/user?charset=utf8")
	db = db1
	//db.SetMaxOpenConns(1000)
	if err != nil {
		fmt.Println("fail to open\n")
		log.Fatal(err)
	}
}

// form表单提交
func PostForm(context *gin.Context) {
	println(">>>> bind form post params action start <<<<")

	type user struct {
		username string
		password string
/*		id string
		sex string
		age string*/
	}
	var u user

	// 绑定参数到结构体
	context.Bind(&u)
	res,err := db.Exec("insert into login (username,password) values ($1,$2)",
		&u.username,&u.password)
	var count int64
	count,err = res.RowsAffected()
	checkError(err)

	if count != 1 {
		context.JSON(200,gin.H{
			"success":false,
		})
	} else {
		// 重定向
		context.Redirect(http.StatusMovedPermanently,"/file/view")
	}

}

// 跳转html
func RenderForm(context *gin.Context) {
	println(">>>> render to html action start <<<<")

	context.Header("Content-Type", "text/html; charset=utf-8")
	context.HTML(200,"Untitled-1.html",gin.H{})
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
