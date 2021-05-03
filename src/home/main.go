package main

import (
	"database/sql"
	"fmt"
	"redertemplate"
	"time"

	_ "github.com/go-sql-driver/mysql"
	// "postapi"
)

type TestMysql struct {
	db *sql.DB
}

/* 初始化数据库引擎 */
func Init() (*TestMysql, error) {
	test := new(TestMysql)
	db, err := sql.Open("mysql", "root:qaz123456@tcp(127.0.0.1:3306)/userpromission")
	if err != nil {
		fmt.Println("database initialize error : ", err.Error())
		return nil, err
	}
	test.db = db
	return test, nil
}

/* 测试数据库数据添加 */
func (test *TestMysql) Create() {
	if test.db == nil {
		return
	}
	stmt, err := test.db.Prepare("insert into user(id,name ,age,create_time)values(?,?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	if result, err := stmt.Exec(1, "张三", 18, time.Now()); err == nil {
		if id, err := result.LastInsertId(); err == nil {
			fmt.Println("insert id : ", id)
		}
	}
}

/* 数据查询 */
func (test *TestMysql) Query() {
	if test.db == nil {
		return
	}
	stmt, err := test.db.Query("select * from user")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s\n", stmt)
	defer stmt.Close()

}

func (test *TestMysql) Close() {
	if test.db != nil {
		test.db.Close()
	}
}

func main() {
	//if test,err := Init();err==nil {
	//	test.Create();
	//	test.Query();
	//	test.Close();
	//}
	//postapi.RegisterAPIRouter()
	redertemplate.RederTemplate()
}
