package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	/*
		1. user@unix(/path/to/socket)/dbname?charset=utf8
		2. user:password@tcp(localhost:5555)/dbname?charset=utf8
		3. user:password@/dbname
		4. user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	*/
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.191.6:3306)/test?charset=utf8")
	checkError(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	//db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	checkError(err)

	res, err := stmt.Exec("root", "devops", "2021-05-27")
	//stmt.Exec()函数用来执行stmt准备好的SQL语句

	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)

	log.Println(id)
	//更新数据

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkError(err)

	affect, err := res.RowsAffected()
	checkError(err)
	log.Println(affect)

	//查询数据
	log.Println("开始查询数据")
	rows, err  := db.Query("SELECT * FROM userinfo")
	//db.Query()函数用来直接执行Sql返回Rows结果。
	checkError(err)
	log.Println(rows)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkError(err)
		log.Println(uid)
		log.Println(username)
		log.Println(department)
		log.Println(created)
	}

	//删除数据
	log.Println("开始删除数据")
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkError(err)

	res, err = stmt.Exec(id)
	checkError(err)

	affect, err = res.RowsAffected()
	checkError(err)

	log.Println(affect)
	db.Close()


}
