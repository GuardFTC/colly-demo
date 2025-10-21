// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-21 11:06:29
package study

import (
	"fmt"

	"github.com/gocolly/colly"
)

// RequestCtxTest 请求上下文测试
func RequestCtxTest() {

	//1.创建采集器
	c := colly.NewCollector()

	//2.在请求之前，将url放入上下文，key为"url"
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	//3.在响应回调中，通过key="url",在上下文中获取url
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Ctx.Get("url"))
	})

	//4.开启访问
	c.Visit("https://en.wikipedia.org/")
}
