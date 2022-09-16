package util

import "strings"

/**
 * @Author: zze
 * @Date: 2022/9/16 18:12
 * @Desc: 字符串相关
 */

const BlankStr string = ""

func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
