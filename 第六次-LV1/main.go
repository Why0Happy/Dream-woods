package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func insertDB(db *sql.DB) {
	stmt, err := db.Prepare("insert into first values (?);")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec("0001,\"woods\",1")
}

func deleteDB(db *sql.DB) {
	stmt, err := db.Prepare("delete from first where aa = 01;")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
}

func updateDB(db *sql.DB) {
	stmt, err := db.Prepare("update first set cc = 1212 where aa = 1;")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(err)
}

func selectDB(db *sql.DB) {
	stmt, err := db.Query("select * from first;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for stmt.Next() {
		var aa int
		var bb string
		var cc int
		err := stmt.Scan(&aa, &bb, &cc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(aa, bb, cc)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/task1?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return
	}
	insertDB(db)
	deleteDB(db)
	updateDB(db)
	selectDB(db)

	db.Close()
}
