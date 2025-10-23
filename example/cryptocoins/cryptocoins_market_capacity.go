// Package cryptocoins @Author:冯铁城 [17615007230@163.com] 2025-10-21 11:39:40
package cryptocoins

import (
	"colly-demo/example/mongo/client"
	"encoding/json"
	"log"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
)

// rank20DataUrl 前20名数据接口
const rank20DataUrl = "https://coinmarketcap.com/all/views/all/"

// rankAfter20DataUrl 后续排行数据接口
const rankAfter20DataUrl = "https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing?sortBy=market_cap&sortType=desc&convert=USD&cryptoType=all&tagType=all&audited=false"

// TestGetCryptocoinsData 获取数字货币市场容量
func TestGetCryptocoinsData() {

	//1.创建采集器
	c := colly.NewCollector()

	//2.设置错误回调
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//3.设置采集器速率限制
	if err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",             //域名匹配规则
		RandomDelay: 5 * time.Second, //随机延迟最大seconds
	}); err != nil {
		log.Printf("error setting limit: %v", err)
	}

	//4.获取前20名数据
	rank20Data := get20RankData(c)

	//5.获取后续排行数据
	rankAfter20Data := getAfter20RankData(c)

	//6.合并数据
	rankData := append(rank20Data, rankAfter20Data...)
	log.Printf("get rank data count:%v", len(rankData))

	//7.保存数据到MongoDB
	if err := saveToMongo(rankData); err != nil {
		log.Printf("save to mongo error:%v", err)
	}
}

// get20RankData 获取前20数据
func get20RankData(c *colly.Collector) []Cryptocurrency {

	//1.定义结构体切片
	var cryptocurrencies []Cryptocurrency

	//2.设置HTML响应回调
	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {

		//3.解析属性，封装为结构体
		cryptocurrency := Cryptocurrency{
			Rank:              cast.ToInt(e.ChildText("td:nth-child(1)")),
			Name:              e.ChildText("td:nth-child(2)"),
			Symbol:            e.ChildText("td:nth-child(3)"),
			MarketCap:         ExtractFullAmount(e.ChildText("td:nth-child(4)")),
			Price:             e.ChildText("td:nth-child(5)"),
			CirculatingSupply: e.ChildText("td:nth-child(6)"),
			Volume24h:         e.ChildText("td:nth-child(7)"),
			PercentChange1h:   e.ChildText("td:nth-child(8)"),
			PercentChange24h:  e.ChildText("td:nth-child(9)"),
			PercentChange7d:   e.ChildText("td:nth-child(10)"),
		}

		//4.存入切片
		if cryptocurrency.Rank == 0 {
			return
		} else {
			cryptocurrencies = append(cryptocurrencies, cryptocurrency)
		}
	})

	//5.请求地址
	if err := c.Visit(rank20DataUrl); err != nil {
		log.Printf("visit error:%v", err)
	}

	//6.等待请求结束
	c.Wait()

	//7.返回结构体切片
	return cryptocurrencies
}

// getAfter20RankData 获取后续排行数据
func getAfter20RankData(c *colly.Collector) []Cryptocurrency {

	//1.获取URL集合
	urls := getUrls(c)

	//2.创建队列
	q, _ := queue.New(
		5, //并发爬虫数量
		&queue.InMemoryQueueStorage{MaxSize: 500}, //队列存储容量
	)

	//3.url写入队列
	for _, url := range urls {
		if err := q.AddURL(url); err != nil {
			log.Printf("add url error:%v", err)
			continue
		}
	}

	//4.定义结构体切片
	var cryptocurrencies []Cryptocurrency

	//5.设置响应回调
	c.OnResponse(func(r *colly.Response) {

		//6.解析JSON到结构体
		var response Response
		err := json.Unmarshal(r.Body, &response)
		if err != nil {
			log.Printf("parse json error:%v", err)
			return
		}

		//7.如果响应异常，打印异常
		if response.Status.ErrorCode != "0" {
			log.Printf("response error:%s", response.Status.ErrorMessage)
			return
		}

		//6.循环封装结构体
		for _, item := range response.Data.CryptoCurrencyList {

			//7.获取报价信息
			quote := item.Quotes[0]

			//8.初始化结构体
			cryptocurrency := Cryptocurrency{
				Rank:              item.CMCRank,
				Name:              item.Name,
				Symbol:            item.Symbol,
				MarketCap:         "$" + FormatUSD(quote.MarketCap),
				Price:             "$" + FormatUSD(quote.Price),
				CirculatingSupply: FormatUSD(item.CirculatingSupply) + " " + item.Symbol,
				Volume24h:         "$" + FormatUSD(quote.Volume24h),
				PercentChange1h:   FormatPercent(quote.PercentChange1h),
				PercentChange24h:  FormatPercent(quote.PercentChange24h),
				PercentChange7d:   FormatPercent(quote.PercentChange7d),
			}

			//9.存入集合
			cryptocurrencies = append(cryptocurrencies, cryptocurrency)
		}
	})

	//10.基于队列爬取数据
	if err := q.Run(c); err != nil {
		log.Printf("queue run error:%v", err)
	}

	//11.返回集合
	return cryptocurrencies
}

// getUrls 获取数据总量
func getUrls(c *colly.Collector) []string {

	//1.定义参数
	url := rankAfter20DataUrl + "&start=1&limit=1"

	//2.定义响应结构体
	var response Response

	//3.设置响应回调
	c.OnResponse(func(r *colly.Response) {

		//4.解析JSON到结构体
		err := json.Unmarshal(r.Body, &response)
		if err != nil {
			log.Printf("parse json error:%v", err)
			return
		}

		//5.如果响应异常，打印异常
		if response.Status.ErrorCode != "0" {
			log.Printf("response error:%s", response.Status.ErrorMessage)
			return
		}
	})

	//6.发起请求
	if err := c.Visit(url); err != nil {
		log.Printf("visit error:%v", err)
	}

	//7.获取数据总量
	totalCount := cast.ToInt(response.Data.TotalCount)

	//8.定义URL切片,首先查询21-100的数据
	urls := []string{rankAfter20DataUrl + "&start=21&limit=100"}

	//9.减去前100名
	totalCount = totalCount - 100

	//10.步长为200，循环拼接后续URL
	for i := 101; i <= totalCount; i += 200 {
		if i+200 > totalCount {
			end := totalCount - i + 1
			urls = append(urls, rankAfter20DataUrl+"&start="+cast.ToString(i)+"&limit="+cast.ToString(end))
		} else {
			urls = append(urls, rankAfter20DataUrl+"&start="+cast.ToString(i)+"&limit=200")

		}
	}

	//11.返回
	return urls
}

// saveToMongo 保存数据到MongoDB
func saveToMongo(cryptocurrencies []Cryptocurrency) error {

	//1.获取客户端
	mongoClient := client.CreateMongoClient()
	defer client.CloseMongoClient(mongoClient)

	//2.声明数据库以及集合
	db := mongoClient.GetClient().Database("testDb")
	collection := db.Collection("cryptocurrencies")

	//3.将 []Cryptocurrency 转换为 []interface{}
	var saveDataList []interface{}
	for _, cryptocurrency := range cryptocurrencies {
		saveDataList = append(saveDataList, cryptocurrency)
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
