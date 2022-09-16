package hertz

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	commonConfig "hertz-demo/common/config"
	"log"
)

/**
 * @Author: zze
 * @Date: 2022/9/15 13:34
 * @Desc:
 */

func NewServer(nacosConfig *commonConfig.Nacos, hertzServerConfig *commonConfig.HertzServer) (*server.Hertz, error) {
	cli, err := nacosConfig.GetNamingClient()
	if err != nil {
		return nil, err
	}
	addr := hertzServerConfig.GetAddr()
	r := nacos.NewNacosRegistry(cli)
	h := server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: hertzServerConfig.Name,
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}))

	h.Use(AccessLog())
	h.Use(NewRequestID())

	url := swagger.URL("/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
	return h, nil
}

func MustNewServer(nacosConfig *commonConfig.Nacos, hertzServerConfig *commonConfig.HertzServer) *server.Hertz {
	h, err := NewServer(nacosConfig, hertzServerConfig)
	if err != nil {
		log.Fatal(err)
	}
	return h
}
