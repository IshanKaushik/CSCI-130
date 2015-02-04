package main

import "fmt"

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {

	fmt.Printf("No Connection\n")

	db, err := sql.Open("mysql",
		"xxx:xxx@tcp(127.0.0.1:3306)/test")

	err = db.Ping()
	if err != nil {
		//do something here
		fmt.Printf("No Connection\n")
	}

	db.Exec("create table name(name char(205))")
	//	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
