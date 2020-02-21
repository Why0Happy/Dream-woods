package main

import (
	"github.com/gin-gonic/gin"
	"bili/controller"
)
func main()  {
	// Engin
	router := gin.Default()
	// 加载html文件，即template包下所有文件
	router.LoadHTMLGlob("template/*")
	//router.GET("/hello", hello)

	// 路由组
	user := router.Group("/user")
	{   // 请求参数在请求路径上
		user.GET("/form",controller.PostForm)// 跳转html页面
		user.POST("/form/post",controller.PostForm)
		//可以自己添加其他，一个请求的路径对应一个函数
	}

	file := router.Group("/file")
	{
		// 跳转上传文件页面
		file.GET("/view",controller.RenderView) // 跳转html页面
		// 根据表单上传
		file.POST("/insert",controller.FormUpload)
		file.POST("/multiUpload",controller.MultiUpload)
		// base64上传
		file.POST("/upload",controller.Base64Upload)
	}

	// 指定地址和端口号
	router.Run(":9090")
}


