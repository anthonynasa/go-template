// gorm gen生成器配置文件

package main

import (
	"gorm.io/gen"
	"gorm.io/gorm"
	"multi-tenant-b/dal"
	"multi-tenant-b/dal/model"
)

func main() {
	// 初始化一个*gorm.DB 实例 dal.DB
	dal.InitDB("ChangSha")
	// 使用gen.Config初始化生成器
	g := gen.NewGenerator(gen.Config{
		OutPath:        "./dal/query",
		Mode:           gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:  true,
		FieldCoverable: true,
	})

	// 使用上面的 `*gorm.DB` 实例来初始化生成器，
	// 使用 `GenerateModel/GenerateModelAs` 时需要从数据库生成结构
	g.UseDB(dal.DB)
	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf和thrift
	// 指定model属性类型和 db 字段类型之间的映射关系。
	dataMap := map[string]func(gorm.ColumnType) (dataType string){
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// Generate default DAO interface for those specified structs:
	// 为那些指定的结构生成默认的 DAO 接口
	// 参考: https://gorm.io/zh_CN/gen/database_to_structs.html
	// 设置要生成的表
	g.ApplyBasic(
		g.GenerateModel("sys_customer"),
		// g.GenerateModel("Base_Devices"),
		// g.GenerateModel("Sys_StateDic"),
		/*
			g.GenerateModel("Base_DataRecord",
				// 忽略字段，被忽略的字段不参与gorm的读写操作
				generate.FieldIgnore("")),
		*/
	)

	// Gen框架使用模板注释的方法支持自定义SQL查询，
	// 我们只需要按对应规则将SQL语句注释到interface的方法上即可。
	// Gen将对其进行解析，并为应用的结构生成查询API。(通常建议将自定义查询方法添加到model模块下。)
	g.ApplyInterface(func(querier model.Querier) {}, g.GenerateModel("sys_customer"))
	// 生成代码
	g.Execute()
}
