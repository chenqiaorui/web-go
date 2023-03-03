package main

import (
	"fmt"

        "database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func main() {

        db, err := sql.Open("mysql", "root:PXDN93VRKUm8TeE7@tcp(192.168.1.167:33069)/andongni?charset=utf8")
        checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
