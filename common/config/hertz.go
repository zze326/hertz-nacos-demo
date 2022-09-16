package config

import "fmt"

/**
 * @Author: zze
 * @Date: 2022/9/14 15:06
 * @Desc:
 */

type HertzServer struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (c *HertzServer) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
