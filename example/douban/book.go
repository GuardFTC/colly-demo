// Package douban @Author:冯铁城 [17615007230@163.com] 2025-10-22 15:11:14
package douban

import (
	"colly-demo/example/mongo/client"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
)

// baseUrl 基础URL
const baseUrl = "https://book.douban.com/tag/小说?start=%v&type=T"

// TestGetDouBanBookData 测试获取豆瓣图书数据
func TestGetDouBanBookData() {

	//1.创建采集器
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0"), //设置浏览器UA，模拟浏览器行为
	)

	//2.设置请求之前触发
	c.OnRequest(func(r *colly.Request) {
		log.Printf("spider is visiting: %s", r.URL)
	})

	//3.设置请求异常回调
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("spider request error: %s", err)
	})

	//4.获取爬取全部数据URL集合
	urls := getUrls(c.Clone())
	log.Printf("spider get url success: %d", len(urls))

	//5.随机睡几秒，模拟正常行为
	time.Sleep(time.Duration(GetRandomSeconds(3, 8)) * time.Second)

	//6.爬取图书数据
	books, err := getBookData(c.Clone(), urls)
	if err != nil {
		log.Println("spider get book data error:", err)
	} else {
		log.Printf("spider get book data success: %d", len(books))
	}

	//7.保存到MongoDB
	if err = saveToMongo(books); err != nil {
		log.Println("save to mongo error:", err)
	}
}

// 获取爬取全部数据的URL
func getUrls(c *colly.Collector) []string {

	//1.定义总页数
	var lastPage int

	//2.设置HTML解析回调，获取总页数
	c.OnHTML("div[class='paginator']", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, el *colly.HTMLElement) {
			if page, err := strconv.Atoi(el.Text); err == nil {
				lastPage = page
			}
		})
	})

	//3.访问链接
	if err := c.Visit("https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4"); err != nil {
		log.Println("spider visit error:", err)
	}

	//4.定义url切片
	var urls []string

	//5.循环总页数，获取所有URL，并加入url切片
	for i := 0; i < lastPage; i++ {

		//6.格式化url
		url := fmt.Sprintf(baseUrl, i*20)

		//7.存入切片
		urls = append(urls, url)
	}

	//8.返回切片
	return urls
}

// getBookData 爬取图书数据
func getBookData(c *colly.Collector, urls []string) ([]Book, error) {

	//1.定义图书切片
	var books []Book

	//2.设置响应回调
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Printf("spider get book data error: %s", r.StatusCode)
		}
	})

	//3.设置HTML解析回调
	c.OnHTML("ul[class='subject-list']", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, el *colly.HTMLElement) {

			//4.解析图书数据
			book := newBook(el)

			//5.写入集合
			books = append(books, book)
		})
	})

	//6.循环URL,按照页数爬取数据
	for _, url := range urls {

		//7.访问URL
		if err := c.Visit(url); err != nil {
			log.Printf("spider visit error: %v", err)
		}

		//8.随机睡几秒，模拟正常行为
		time.Sleep(time.Duration(GetRandomSeconds(4, 8)) * time.Second)
	}

	//9.爬取完成返回集合
	return books, nil
}

// saveToMongo 保存数据到MongoDB
func saveToMongo(books []Book) error {

	//1.获取客户端
	mongoClient := client.CreateMongoClient()
	defer client.CloseMongoClient(mongoClient)

	//2.声明数据库以及集合
	db := mongoClient.GetClient().Database("testDb")
	collection := db.Collection("books")

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
