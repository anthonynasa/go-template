// 路由配置

package route

import (
	"github.com/gin-gonic/gin"
	"multi-tenant-b/handler"
)

func SetupRoute() *gin.Engine {
	// 创建gin引擎
	r := gin.Default()

	// 映射static目录
	r.Static("/assets", "./web/assets")
	r.Static("/v1/assets", "./web/assets")
	r.StaticFile("/favicon.ico", "./web/assets/images/favicon.png")

	// 加载html文件
	r.LoadHTMLGlob("./web/*.html")

	// 路由注册

	// index路由
	r.GET("/index", handler.Index)

	// v1: customer CRUD
	v1 := r.Group("/v1")
	{
		v1.GET("/:city-name/:id", handler.GetCustomerById)
		v1.POST("/:city-name/:id", nil)
		v1.PUT("/:city-name/:id", nil)
		v1.DELETE("/:city-name/:id", nil)
	}

	// v2
	v2 := r.Group("/v2")
	{
		v2.GET("/customer", nil)
	}

	// 注册 处理未匹配路由的中间件
	r.NoRoute(nil)

	return r
}
