package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
)

var db *sql.DB

func Instance() *sql.DB {
	return db
}
func Create() *sql.DB {
	db, err := sql.Open("sqlite3", "./emcs.db")
	checkErr(err)
	createTab(db)
	db.Begin()
	installData(db)
	return db
}

func createTab(db *sql.DB) {
	file, _ := os.ReadFile("./sql/table.sql")
	sqlString := strings.Split(strings.ReplaceAll(string(file), "\n", ""), ";")
	for index, sql := range sqlString {
		fmt.Println(index)
		if sql != "" {
			db.Exec(sql)
			fmt.Println(sql)
		}
	}

}
func installData(db *sql.DB) {
	file, _ := os.ReadFile("./sql/data.sql")
	sqlString := strings.Split(strings.ReplaceAll(string(file), "\n", ""), ";")
	for index, sql := range sqlString {
		fmt.Println(index)
		if sql != "" {
			db.Exec(sql)
			fmt.Println(sql)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
