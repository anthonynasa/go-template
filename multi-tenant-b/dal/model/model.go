// 自定义结构体

package model

import "time"

type Customer struct {
	TCH        *string    `gorm:"column:TCH" json:"TCH"`
	YHBH       string     `gorm:"column:YHBH;primaryKey" json:"YHBH"`
	YHDZ       string     `gorm:"column:YHDZ" json:"YHDZ"`
	CMNAME     *string    `gorm:"column:CMNAME" json:"CMNAME"`
	BARCODE    *string    `gorm:"column:BARCODE" json:"BARCODE"`
	ChLastTime *time.Time `gorm:"column:chLastTime" json:"chLastTime"`
	ChLastData *int64     `gorm:"column:chLastData" json:"chLastData"`
	ChAccData  *int64     `gorm:"column:chAccData" json:"chAccData"`

	// ChReadType *string `gorm:"column:chReadType" json:"chReadType"`
	SdName string `gorm:"column:sd_Name;not null" json:"sd_Name"`

	ChBattery *int64  `gorm:"column:chBattery" json:"chBattery"`
	Chvalve   *int64  `gorm:"column:chvalve" json:"chvalve"`
	DeVersion *string `gorm:"column:de_Version;not null;default:-" json:"de_Version"`
}
