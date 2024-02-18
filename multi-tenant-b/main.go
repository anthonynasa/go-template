// 主程序入口文件
package main

import (
	"multi-tenant-b/route"
)

func main() {
	// 初始化一个*gorm.DB 实例 dal.DB
	// dal.InitDB("WuXi")
	// // 为query设置*gorm.DB实例
	// query.SetDefault(dal.DB)

	// 配置路由
	r := route.SetupRoute()
	// 配置端口
	// 只允许本地,localhost: port
	r.Run("127.0.0.1:8080")
	// 允许任意,localhost: port
	// r.Run("0.0.0.0:9090")
}
