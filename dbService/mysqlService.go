package dbService
import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"MANGAGO/service"
)

//"github.com/satori/go.uuid"

//https://blog.51cto.com/9291927/2344802
func setupConnect() *sql.DB {
   db, err := sql.Open("mysql", "root:xxxxxx@tcp(" + service.GetConfig("db_ip") + ":" + service.GetConfig("db_port") + ")/" + service.GetConfig("db_name") + "?charset=utf8")
   //errorHandler(err)
   return db
}

func init() {
   db := setupConnect()
}

func Disconnect(){
   // 关闭数据库连接
   db.Close()
}