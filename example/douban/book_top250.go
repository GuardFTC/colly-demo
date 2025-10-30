// Package douban @Author:冯铁城 [17615007230@163.com] 2025-10-30 15:42:56
package douban

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

// url 请求路由
const urlTemplate = "https://book.douban.com/top250?start=%v"

// userAgent 模拟浏览器UA
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0"

func TestGetTop250BookData() {

	//1.获取top250图书访问URL
	urls := getTop250Urls()
	if len(urls) == 0 {
		log.Println("get urls error")
		return
	}

	//2.获取top250图书数据
	books, err := getTop250BookData(urls)
	if err != nil {
		log.Println("get book data error:", err)
		return
	}
	if len(books) == 0 {
		log.Println("get book data error")
		return
	}

	//3.保存到MongoDB
	if err = saveToMongo(books, "top250_books"); err != nil {
		log.Println("save to mongo error:", err)
	}
}

// 获取top250图书访问URL
func getTop250Urls() []string {

	//1.定义url切片
	var urls []string

	//2.循环总页数，获取所有URL，并加入url切片
	for i := 0; i < 10; i++ {

		//3.格式化url
		url := fmt.Sprintf(urlTemplate, i*25)

		//4.存入切片
		urls = append(urls, url)
	}

	//5.返回切片
	return urls
}

// 获取top250图书数据
func getTop250BookData(urls []string) ([]Book, error) {

	//1.定义图书切片
	var books []Book

	//2.创建采集器
	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	//3.设置请求回调
	c.OnRequest(func(r *colly.Request) {
		log.Printf("spider request url: %v", r.URL)
	})

	//4.设置异常回调
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("spider request error: %s", err)
	})

	//5.设置响应回调
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Printf("spider get book data error: %s", r.StatusCode)
		}
	})

	//6.设置HTML解析回调
	c.OnHTML("div[class='indent']", func(e *colly.HTMLElement) {
		e.ForEach("table", func(i int, el *colly.HTMLElement) {

			//7.解析图书数据
			book := newTop250Book(el)

			//8.写入集合
			books = append(books, book)
		})

		//9.打印Book的
		log.Printf("books add success. book len is: %d", len(books))
	})

	//10.设置请求限制
	if err := c.Limit(&colly.LimitRule{
		DomainGlob:  `book.douban.com`,
		Delay:       3 * time.Second,
		RandomDelay: 500 * time.Millisecond,
	}); err != nil {
		log.Printf("error setting limit: %v", err)
	}

	//11.循环访问URL
	for _, url := range urls {
		if err := c.Visit(url); err != nil {
			log.Printf("error visiting %s: %v", url, err)
		}
	}

	//12.爬取完成返回集合
	return books, nil
}
