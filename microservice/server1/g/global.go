package g

import (
	"gorm.io/gorm"
	"hertz-demo/common/hertz"
	"hertz-demo/microservice/server1/etc"
)

/**
 * @Author: zze
 * @Date: 2022/9/14 15:23
 * @Desc: ćšć±ćé
 */

var (
	DB          *gorm.DB
	HertzClient *hertz.Client
	Config      = new(etc.Config)
)
