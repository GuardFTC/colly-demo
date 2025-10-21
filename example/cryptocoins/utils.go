// Package cryptocoins @Author:冯铁城 [17615007230@163.com] 2025-10-21 16:48:17
package cryptocoins

import (
	"fmt"
	"regexp"
	"strings"
)

// ExtractFullAmount 提取完整的金额
func ExtractFullAmount(input string) string {

	//1.匹配完整的美元金额格式：$ + 数字 + 逗号（可选）
	re := regexp.MustCompile(`\$[\d,]+`)

	//2.找到所有匹配
	matches := re.FindAllString(input, -1)

	//3.返回最后一个匹配（完整的金额）
	if len(matches) > 0 {
		return matches[len(matches)-1]
	}

	//4.否则返回原始输入
	return input
}

// FormatUSD 通用USD格式化函数
// 自动判断: 大数→$X,XXX,XXX | 小数→$0.XXXXXXX
func FormatUSD(value float64) string {
	// 转换为字符串(8位小数)
	str := fmt.Sprintf("%.8f", value)

	// 分离整数和小数部分
	parts := strings.Split(str, ".")
	integer := parts[0]
	decimal := parts[1]

	// 格式化整数部分(千分位逗号)
	formattedInt := formatInteger(integer)

	// 规则: 如果是小数价格(0.xxxx)，显示8位小数；否则只显示整数
	if strings.HasPrefix(formattedInt, "0") && value < 1 {
		return formattedInt + "." + decimal[:8] // $0.00000995
	}

	return formattedInt // $5,854,478,863
}

// FormatPercent 百分比格式化(保留2位小数 + %符号)
// 示例: 0.35914323 → +0.36% | -3.14789939 → -3.15%
func FormatPercent(value float64) string {
	// 保留2位小数
	formatted := fmt.Sprintf("%.2f", value)

	// 添加符号(+/-)
	if value >= 0 {
		formatted = "+" + formatted
	}

	return formatted + "%"
}

// formatInteger 格式化整数部分(添加千分位逗号)
func formatInteger(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		if (len(s)-1-i)%3 == 0 && i != 0 {
			result.WriteByte(',')
		}
		result.WriteByte(s[i])
	}
	return reverseString(result.String())
}

// reverseString 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
