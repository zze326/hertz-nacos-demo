package model

import "hertz-demo/common/model"

/**
 * @Author: zze
 * @Date: 2022/6/10 10:13
 * @Desc: Helm chart 模型
 */

type HelmChart struct {
	model.Model
	Name      string `gorm:"type:varchar(255);unique_index" json:"name"`
	Desc      string `gorm:"type:varchar(255)" json:"desc"`
	Type      uint8  `gorm:"type:tinyint" json:"type"`
	IsDefault bool   `gorm:"type:tinyint" json:"is_default"`
	Readonly  bool   `gorm:"type:tinyint" json:"readonly"`
}
