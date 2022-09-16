package g

import (
	"hertz-demo/common/hertz"
	"hertz-demo/microservice/server1/etc"
)

/**
 * @Author: zze
 * @Date: 2022/9/14 15:23
 * @Desc: 服务全局变量 & 服务上下文初始化
 */

var (
	Config      = new(etc.Config)
	HertzClient *hertz.Client
)

type ServiceContext struct {
}

func InitServiceContext() {
}
