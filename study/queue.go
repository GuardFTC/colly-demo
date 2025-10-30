// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-20 19:33:02
package study

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)

// QueueTest 队列测试
func QueueTest() {

	//1.创建队列
	q, _ := queue.New(
		2, //并发爬虫数量
		&queue.InMemoryQueueStorage{MaxSize: 10000}, //队列存储容量
	)

	//2.定义URL
	url := "https://httpbin.org/delay/1"

	//3.创建wg
	var wg sync.WaitGroup

	//4.向队列中添加5个URL
	for i := 0; i < 5; i++ {
		wg.Add(1)
		if err := q.AddURL(fmt.Sprintf("%s?n=%d", url, i)); err != nil {
			log.Printf("add url to queue error: %v", err)
		}
	}

	//5.创建采集器
	c := colly.NewCollector(
		colly.Async(true),
	)

	//6.请求之前触发函数
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	//7.响应回调函数
	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response Body: %s\n", string(r.Body))
		wg.Done()
	})

	//8.通过队列中的协程发起爬虫请求
	if err := q.Run(c); err != nil {
		log.Printf("queue run error: %v", err)
	}

	//9.阻塞等待全部响应回调完成
	wg.Wait()
}
