// Package douban @Author:冯铁城 [17615007230@163.com] 2025-10-22 15:11:14
package douban

import (
	"colly-demo/example/mongo/client"
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
)

// host 服务器地址
const host = "https://book.douban.com"

// path 接口地址
const path = "/tag/小说"

// payloadTemplate 请求参数模版
const payloadTemplate = "?start=%v&type=T"

// TestGetDouBanBookData 测试获取豆瓣图书数据
func TestGetDouBanBookData() {

	//1.获取爬取全部数据URL集合
	urls := getUrls()
	log.Printf("spider get url success: %d", len(urls))

	//2.URL集合为空，直接返回
	if len(urls) == 0 {
		return
	}

	//3.随机睡几秒，模拟正常行为
	time.Sleep(time.Duration(GetRandomSeconds(5, 10)) * time.Second)

	//4.爬取图书数据
	books, err := getBookData(urls)
	if err != nil {
		log.Println("spider get book data error:", err)
	} else {
		log.Printf("spider get book data success: %d", len(books))
	}

	//5.保存到MongoDB
	if err = saveToMongo(books, "books"); err != nil {
		log.Println("save to mongo error:", err)
	}
}

// 获取爬取全部数据的URL
func getUrls() []string {

	//1.定义url切片
	var urls []string

	//2.循环总页数，获取所有URL，并加入url切片
	for i := 0; i < 50; i++ {

		//3.格式化url
		url := fmt.Sprintf(host+path+payloadTemplate, i*20)

		//4.存入切片
		urls = append(urls, url)
	}

	//5.返回切片
	return urls
}

// getBookData 爬取图书数据
func getBookData(urls []string) ([]Book, error) {

	//1.定义图书切片
	var books []Book

	//3.创建采集器
	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	//4.设置真实UserAgent
	c.OnRequest(func(r *colly.Request) {
		log.Printf("spider request url: %v", r.URL)
	})

	//5.设置请求异常回调
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("spider request error: %s", err)
	})

	//6.设置响应回调
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Printf("spider get book data error: %s", r.StatusCode)
		}
	})

	//7.设置HTML解析回调
	c.OnHTML("ul[class='subject-list']", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, el *colly.HTMLElement) {

			//8.解析图书数据
			book := newBook(el)

			//9.写入集合
			books = append(books, book)
		})

		//10.打印Book的
		log.Printf("books add success. book len is: %d", len(books))
	})

	//10.设置请求限制
	if err := c.Limit(&colly.LimitRule{
		DomainGlob:  `book.douban.com`,
		Delay:       3 * time.Second,
		RandomDelay: 600 * time.Millisecond,
	}); err != nil {
		log.Printf("error setting limit: %v", err)
	}

	//12.循环URL,按照页数爬取数据
	for _, url := range urls {

		//13.访问URL
		if err := c.Visit(url); err != nil {
			log.Printf("spider visit error: %v", err)
		}
	}

	//17.爬取完成返回集合
	return books, nil
}

// saveToMongo 保存数据到MongoDB
func saveToMongo(books []Book, collectionName string) error {

	//1.获取客户端
	mongoClient := client.CreateMongoClient()
	defer client.CloseMongoClient(mongoClient)

	//2.声明数据库以及集合
	db := mongoClient.GetClient().Database("testDb")
	collection := db.Collection(collectionName)

	//3.将 []Book 转换为 []interface{}
	var saveDataList []interface{}
	for _, book := range books {
		saveDataList = append(saveDataList, book)
	}

	//4.全量删除数据
	if _, err := collection.DeleteMany(mongoClient.GetCtx(), bson.M{}); err != nil {
		return err
	} else {
		log.Println("delete mongo data success")
	}

	//5.保存
	if saveRes, err := collection.InsertMany(mongoClient.GetCtx(), saveDataList); err != nil {
		return err
	} else {
		log.Printf("save to mongo success, data count:%v", len(saveRes.InsertedIDs))
		return nil
	}
}
