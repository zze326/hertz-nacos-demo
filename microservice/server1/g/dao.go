package g

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz-demo/microservice/server1/repository/dao"
)

/**
 * @Author: zze
 * @Date: 2022/9/15 13:50
 * @Desc: 数据访问对象初始化
 */

var (
	HelmChartDao *dao.HelmChartDao
)

func initDao() {
	var err error
	DB, err = Config.Mysql.InitGorm()
	if err != nil {
		hlog.Fatal("init gorm error: %v", err)
	}

	HelmChartDao = dao.NewHelmChartDao(DB)
}
