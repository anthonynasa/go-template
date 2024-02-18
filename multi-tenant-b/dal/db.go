// 初始化数据连接

package dal

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"strings"
)

var (
	DB *gorm.DB
)

// InitDB
//
//	@Description: 根据参数cityName初始化全局变量 DB *gorm.DB
//	@param cityName
func InitDB(cityName string) {
	var err error
	// 转换为小写
	cityName = strings.ToLower(cityName)
	// 设置viper读取配置文件
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %s", err)
	}
	// 读取数据库配置
	host := viper.GetString(fmt.Sprintf("%s.host", cityName))
	port := viper.GetInt(fmt.Sprintf("%s.port", cityName))
	user := viper.GetString(fmt.Sprintf("%s.user", cityName))
	password := viper.GetString(fmt.Sprintf("%s.password", cityName))
	dbname := viper.GetString(fmt.Sprintf("%s.dbname", cityName))
	// 动态拼接 Gorm 的 DSN
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&connection+timeout=50&encrypt=disable",
		user, password, host, port, dbname)
	// 初始化一个*gorm.DB 实例
	// 实际上通过gorm.Open函数并没有和数据库建立连接，而只是返回了一个全局的gorm.DB对象。
	// 真正的数据库连接是在具体执行sql语句时才建立的。
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取数据库连接对象失败: " + err.Error())
	}
}

/*
// 弃用
func SwitchDB(cityName string) {
	// 关闭之前的数据库连接
	preDB, _ := DB.DB()
	preDB.Close()

	// 创建新的数据库连接
	InitDB(cityName)
}

*/
