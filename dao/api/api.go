package api
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterAPIRouter() {
	gin.SetMode(gin.DebugMode) //调试模式
	router := gin.Default()     //获得路由实例

	routerDatasource := router.Group("/data/source")
	// 监听/data/source/connect的get和post请求，对应方法：ConnTest
	routerDatasource.GET("/connect", ConnTest)
	routerDatasource.POST("/connect", ConnTest)

	//监听端口
	http.ListenAndServe(":9000", router)

}

func ConnTest(c *gin.Context) {
	var(
		status int
		desc string
	)
	dbtype := c.Query("dbtype")
	dbname := c.Query("dbname")
	user   := c.Query("user")
	password := c.Query("password")
	host := c.Query("host")
	port := c.Query("port")

	constr := user+":"+password+"@tcp("+host+":"+port+")/"+dbname

	db, err := sql.Open(dbtype, constr)
	err = db.Ping()  //sql.Open无法断定数据库是否正常连接，所以调用db.Ping()进行判定
	if err != nil {
		status = 300
		desc = "数据库连接失败"
	}else{
		status = 200
		desc = "数据库连接成功"
	}
	c.JSON(200, gin.H{"status": status,"msg": desc})
}