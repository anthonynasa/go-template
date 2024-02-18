// 处理器,中间件

package handler

import (
	"github.com/gin-gonic/gin"
	"multi-tenant-b/dal"
	"multi-tenant-b/dal/query"
	"net/http"
)

// Index
//
//	@Description: 主页
//	@param c
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// GetCustomerById
//
//	@Description: gen根据接口注释自动生成的接口方法,根据用户编号查询单个用户
//	@param c
func GetCustomerById(c *gin.Context) {
	// 获取参数
	cityName := c.Param("city-name")
	id := c.Param("id")
	// 切换数据库,设置*gorm.DB实例为 dal.DB
	dal.InitDB(cityName)
	// 为query设置*gorm.DB实例
	query.SetDefault(dal.DB)
	// gorm查询
	s := query.SysCustomer
	customer, err := s.GetByID(id)
	if err != nil {
		return
	}
	// 返回查询结果
	c.JSON(http.StatusOK, customer)
}

/*

// GetCustomerById2
//
//	@Description: 多表join查询用户信息
//	@param c
func GetCustomerById2(c *gin.Context) {
	// 获取表单数据
	cityName := c.Param("city-name")
	id := c.Param("id")

	fmt.Println(cityName, id)

	// 切换数据库,设置*gorm.DB实例为 dal.DB
	dal.InitDB(cityName)
	// 为query设置*gorm.DB实例
	query.SetDefault(dal.DB)

	s := query.SysCustomer
	d := query.BaseDevice
	t := query.SysStateDic
	var customer []model.Customer
	// select sd_name from Sys_StateDic where sd_Table = 'ch_ReadType'  and sd_KeyWord = chReadType
	err := s.
		Select(s.TCH, s.YHBH, s.YHDZ, s.CMNAME, s.BARCODE, s.ChLastTime, s.ChLastData, s.ChAccData,
			t.SdName, s.ChBattery, s.Chvalve, d.DeVersion).
		LeftJoin(d, s.BARCODE.EqCol(d.DeBarCode)).
		LeftJoin(t, s.ChReadType.EqCol(t.SdKeyWord)).
		Where(s.YHBH.Eq(id)).
		Scan(&customer)

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, customer)
}


*/
