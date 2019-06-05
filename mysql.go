package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int64
	Email sql.NullString
}

func queryOne(db *sql.DB) User {
	user := new(User)
	//查询
	row := db.QueryRow(`select id, email from users where id = 109`)
	if err := row.Scan(&user.ID, &user.Email); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}
	return *user
}

func queryRows(db *sql.DB) {
	rows, err := db.Query(`select id, email from users limit 3`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		user := new(User)
		rows.Scan(&user.ID, &user.Email)
		fmt.Println(user)
		fmt.Println("email is:", user.Email.String)
	}
}

func main() {
	//链接数据库
	dbw, err := sql.Open("mysql", "root:jfinfo88@tcp(192.168.102.6:3306)/jugupiao")
	if err != nil {
		panic(err)
	}
	user := queryOne(dbw)
	fmt.Println(user)
	queryRows(dbw)
	//var a = 2
	//var b = a
	//fmt.Println(&a)
	//fmt.Println(a)
	//a = 3
	//fmt.Println(b)
}
