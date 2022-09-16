package etc

import "hertz-demo/common/config"

/**
 * @Author: zze
 * @Date: 2022/9/14 13:50
 * @Desc:
 */

type Config struct {
	Server config.HertzServer `yaml:"server"`
}
