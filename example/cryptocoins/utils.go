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

// FormatPercent 百分比格式化(保留2位小数 + %符号)
// 示例: 0.35914323 → +0.36% | -3.14789939 → -3.15%
func FormatPercent(value float64) string {

	//1.保留2位小数
	formatted := fmt.Sprintf("%.2f", value)

	//2.添加符号(+/-)
	if value >= 0 {
		formatted = "+" + formatted
	}

	//3.返回
	return formatted + "%"
}

// FormatUSD 标准库实现
func FormatUSD(value float64) string {

	//1.小数保留8位，返回
	if value < 1 {
		return fmt.Sprintf("%.8f", value)
	}

	//2.分离整数和小数部分
	str := fmt.Sprintf("%.2f", value)
	parts := strings.Split(str, ".")
	integer := parts[0]
	decimal := parts[1]

	//3.整数部分添加千分位分隔符，拼接小数部分，返回
	return thousandSeparator(integer) + "." + decimal
}

// thousandSeparator 千分位分隔符
func thousandSeparator(s string) string {
	n := len(s)
	var b strings.Builder
	for i, r := range s {
		if i > 0 && (n-i)%3 == 0 {
			b.WriteRune(',')
		}
		b.WriteRune(r)
	}
	return b.String()
}
