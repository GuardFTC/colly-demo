// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-20 17:07:11
package study

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ErrorTest() {

	//1.创建采集器
	c := colly.NewCollector()

	//2.当响应为HTML，执行回调函数
	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	//3.设置异常处理器
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//4.访问错误页面
	if err := c.Visit("https://definitely-not-a.website/"); err != nil {
		fmt.Println(err)
	}
}
