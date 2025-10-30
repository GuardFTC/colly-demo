// Package douban @Author:冯铁城 [17615007230@163.com] 2025-10-23 16:04:59
package douban

import (
	"math/rand"
)

// GetRandomSeconds 生成指定范围内的随机秒数
// min: 最小值（秒）
// max: 最大值（秒）
func GetRandomSeconds(min, max int) int {
	return rand.Intn(max-min+1) + min
}
