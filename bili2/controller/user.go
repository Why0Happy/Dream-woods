package controller

import (
	"github.com/gin-gonic/gin"
)


//主页
func homepage(context *gin.Context){
	println("----homepage start----")

	context.Header("Content-Type", "text/html; charset=utf-8")
	context.HTML(200,"Untitled-1.html",gin.H{})
}

//注册
/*func register(context *gin.Context) {

	println("----register start----")

	username:=context.PostForm("username")
	password:=context.PostForm("password")
	repassword:=context.PostForm("password-repeat")

	if password!=repassword {
		context.JSON(400, gin.H{
			"success":  false,
			"message": "两次密码输入不同，请再次输入",
		})
	}
	println("----mysql start----")
	stmt, err := db.Prepare(
		"insert into login(username,password)values(?,?) ")
	if err != nil {
		fmt.Println("fail to Prepare\n")
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(username,password)
	if err != nil {
		fmt.Println("fail to Exec\n")
		log.Fatal(err)
	}

}*/
