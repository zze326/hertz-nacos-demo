package config

/**
 * @Author: zze
 * @Date: 2022/5/25 10:56
 * @Desc: Nacos 配置
 */
import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"hertz-demo/common/util"
	"sync"
)

var (
	configClient config_client.IConfigClient
	nacosOnce    sync.Once
)

type Nacos struct {
	Addr        string
	Port        uint64
	Group       string
	DataID      string
	ExtDataIDs  []string `json:",optional"`
	NamespaceID string
}

func (conf *Nacos) InitConfigClient() (err error) {
	nacosOnce.Do(func() {
		configClient, err = clients.NewConfigClient(
			vo.NacosClientParam{
				ClientConfig: &constant.ClientConfig{TimeoutMs: 5000, NamespaceId: conf.NamespaceID},
				ServerConfigs: []constant.ServerConfig{
					{IpAddr: conf.Addr, Port: conf.Port},
				},
			},
		)
	})
	return
}

func (conf *Nacos) GetConfig() (string, error) {
	var configMap = make(map[interface{}]interface{})
	mainConfig, err := configClient.GetConfig(vo.ConfigParam{DataId: conf.DataID, Group: conf.Group})
	if err != nil {
		return "", err
	}

	mainMap, err := util.UnmarshalYamlToMap(mainConfig)
	if err != nil {
		return "", err
	}

	var extMap = make(map[interface{}]interface{})
	for _, dataID := range conf.ExtDataIDs {
		extConfig, err := configClient.GetConfig(vo.ConfigParam{DataId: dataID, Group: conf.Group})
		if err != nil {
			return "", err
		}

		tmpExtMap, err := util.UnmarshalYamlToMap(extConfig)
		if err != nil {
			return "", err
		}

		extMap = util.MergeMap(extMap, tmpExtMap)
	}

	configMap = util.MergeMap(configMap, extMap)
	configMap = util.MergeMap(configMap, mainMap)

	yamlString, err := util.MarshalObjectToYamlString(configMap)
	if err != nil {
		return "", err
	}

	return yamlString, nil
}

func MustLoad(nacosConfigFilePath string, v interface{}) *Nacos {
	var (
		err    error
		config string
	)

	viper.SetConfigFile(nacosConfigFilePath)
	viper.AutomaticEnv()

	var nacosConfig Nacos
	if err = viper.ReadInConfig(); err != nil {
		hlog.Fatalf("read config error: %v", err)
	}
	if err = viper.Unmarshal(&nacosConfig); err != nil {
		hlog.Fatalf("unmarshal config error: %v", err)
	}
	err = nacosConfig.InitConfigClient()
	if err != nil {
		hlog.Fatalf("init config client error: %v", err)
	}

	config, err = nacosConfig.GetConfig()
	if err != nil {
		hlog.Fatalf("get config error: %v", err)
	}

	if err = yaml.Unmarshal([]byte(config), v); err != nil {
		hlog.Fatalf("load config error: %v", err)
	}

	return &nacosConfig
}

func (conf *Nacos) GetNamingClient() (naming_client.INamingClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Addr, conf.Port),
	}

	cc := constant.ClientConfig{
		NamespaceId:         conf.NamespaceID,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              ".nacos/log",
		CacheDir:            ".nacos/cache",
		LogLevel:            "info",
	}

	return clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
}
