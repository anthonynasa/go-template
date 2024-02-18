// gen自定义sql查询
// Gen框架使用模板注释的方法支持自定义SQL查询，
// 按对应规则将SQL语句注释到interface的方法上 ,在generate.go中配置ApplyInterface
// Gen将对其进行解析，并为应用的结构生成查询API (通常建议将自定义查询方法添加到model模块下。)

package model

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE YHBH=@YHBH
	GetByID(YHBH string) (gen.T, error)
}
