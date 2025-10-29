// Package douban @Author:冯铁城 [17615007230@163.com] 2025-10-23 16:04:59
package douban

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gocolly/colly"
)

// GetRandomSeconds 生成指定范围内的随机秒数
// min: 最小值（秒）
// max: 最大值（秒）
func GetRandomSeconds(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// setRequestHeaders 设置请求头
func setRequestHeaders(r *colly.Request, uaPool *userAgentPool, pageNum int) *colly.Request {

	//1.设置动态UserAgent
	userAgent := uaPool.getUserAgent()
	r.Headers.Set("User-Agent", userAgent)

	//2.设置Referer
	if pageNum != 0 {
		if pageNum == 1 {
			r.Headers.Set("Referer", host+path)
		} else {
			r.Headers.Set("Referer", fmt.Sprintf(host+path+payloadTemplate, (pageNum-1)*20))
		}
	}

	//3.根据User-Agent设置相应的安全头部
	setSecurityHeaders(r, userAgent)

	//4.设置通用请求头
	r.Headers.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	r.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	r.Headers.Set("cache-control", "no-cache")
	r.Headers.Set("pragma", "no-cache")
	r.Headers.Set("upgrade-insecure-requests", "1")

	//5.返回
	return r
}

// setSecurityHeaders 根据User-Agent设置相应的安全头部
func setSecurityHeaders(r *colly.Request, userAgent string) {

	//1.根据User-Agent设置相应的安全头部
	if strings.Contains(userAgent, "Edg") || strings.Contains(userAgent, "Edge") {

		//2.Edge浏览器
		r.Headers.Set("sec-ch-ua", `"Microsoft Edge";v="141", "Not?A_Brand";v="8", "Chromium";v="141"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
	} else if strings.Contains(userAgent, "Chrome") && !strings.Contains(userAgent, "Edg") {

		//3.Chrome浏览器
		r.Headers.Set("sec-ch-ua", `"Google Chrome";v="141", "Not?A_Brand";v="8", "Chromium";v="141"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
	} else if strings.Contains(userAgent, "Firefox") {

		//4.Firefox浏览器
		r.Headers.Set("sec-ch-ua", `"Firefox";v="125", "Not?A_Brand";v="8", "Chromium";v="141"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
	} else if strings.Contains(userAgent, "Safari") && !strings.Contains(userAgent, "Chrome") {

		//5.Safari浏览器
		r.Headers.Set("sec-ch-ua", `"Safari";v="17", "Not?A_Brand";v="8", "Chromium";v="141"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"macOS"`)
	} else {

		//6.默认设置
		r.Headers.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="141", "Microsoft Edge";v="141"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
	}

	//7.设置通用的安全头部
	r.Headers.Set("sec-fetch-dest", "document")
	r.Headers.Set("sec-fetch-mode", "navigate")
	r.Headers.Set("sec-fetch-site", "same-origin")
	r.Headers.Set("sec-fetch-user", "?1")
}
