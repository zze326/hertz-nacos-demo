package dao

import (
	"gorm.io/gorm"
	"hertz-demo/microservice/server1/repository/model"
)

/**
 * @Author: zze
 * @Date: 2022/6/10 10:15
 * @Desc: helm chart 数据访问层
 */

type HelmChartDao struct {
	*gorm.DB
}

func NewHelmChartDao(db *gorm.DB) *HelmChartDao {
	return &HelmChartDao{DB: db}
}

func (dao *HelmChartDao) GetByID(id int) (*model.HelmChart, error) {
	var chart model.HelmChart
	err := dao.Find(&chart, id).Error
	if err != nil {
		return nil, err
	}
	return &chart, nil
}
