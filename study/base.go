// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-20 17:05:19
package study

import (
	"fmt"

	"github.com/gocolly/colly"
)

// BaseDemoTest 基础示例
func BaseDemoTest() {

	//1.创建采集器
	c := colly.NewCollector(

		//2.设置采集器允许访问的域名
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	//3.当响应为HTML，并且匹配a[href]标签时，执行回调函数
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		//4.获取连接
		link := e.Attr("href")

		//5.打印
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		//6.访问连接
		err := c.Visit(e.Request.AbsoluteURL(link))
		if err != nil {
			return
		}
	})

	//7.在请求之前打印"正在访问..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	//8.访问hackerspaces.org
	if err := c.Visit("https://hackerspaces.org/"); err != nil {
		panic(err)
	}
}
