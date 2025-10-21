// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-20 17:35:28
package study

import (
	"fmt"

	"github.com/gocolly/colly"
)

func MaxDepthTest() {

	//1.创建采集器
	c := colly.NewCollector(

		//2.设置爬取的最大深度为1,即只爬取当前页面的链接
		colly.MaxDepth(1),
	)

	//3.当响应为HTML，并且匹配a[href]标签时，执行回调函数
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		//4.获取连接
		link := e.Attr("href")

		//5.打印连接
		fmt.Println(link)

		//6.访问连接，实际上不会执行，因为MaxDepth为1，只会爬取当前页面的链接
		if err := e.Request.Visit(link); err != nil {
			fmt.Println(err)
		}
	})

	//7.访问wiki
	if err := c.Visit("https://en.wikipedia.org/"); err != nil {
		fmt.Println(err)
	}
}
