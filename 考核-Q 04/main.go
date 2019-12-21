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

type user struct {
	name string
	id int
	sums int
}

type admi struct {
	name string
}

func check (u *gin.Context,){
	stmt, err := db.Query("select sums from user where id=?", id)
	defer stmt.Close()
	if err != nil {
		println("fail to Select\n")
		log.Fatal(err)
	}
	var sums int
	for stmt.Next() {
		stmt.Scan(&sums)
		if err != nil {
			println("fail to Scan\n")
			log.Fatal(err)
		}
	}

		u.JSON(200, gin.H{
			"sums": id,


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
	u.POST("/check", check)

	u.Run(":8080")
}
