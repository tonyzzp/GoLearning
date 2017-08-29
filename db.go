package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func main() {
	os.Remove("./test.db")
	fmt.Println(sql.Drivers())
	db, _ := sql.Open("sqlite3", "./test.db")

	sql := `
	create table user(id integer primary key autoincrement,name text,age integer,money float)
	`
	db.Exec(sql)

	st, _ := db.Prepare("INSERT INTO user (name,age,money) VALUES (?,?,?);")
	st.Exec("zzp", 18, 1000000.8)

	row, _ := db.Query("SELECT * FROM user")
	for row.Next() {
		var id int32
		var name string
		var age int32
		var money float32
		row.Scan(&id, &name, &age, &money)
		fmt.Println(id, name, age, money)
	}
	row.Close()
	db.Close()
}
