// Package example @Author:冯铁城 [17615007230@163.com] 2025-10-21 11:39:40
package cryptocoins

import (
	"encoding/json"
	"log"

	"github.com/gocolly/colly"
	"github.com/spf13/cast"
)

// CryptocoinsMarketCapacity 获取数字货币市场容量
func CryptocoinsMarketCapacity() {

	//1.创建采集器
	c := colly.NewCollector()

	//2.设置错误回调
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//3.获取前20名数据
	cryptocurrencies := get20RankData(c)

	//4.获取后续排行数据
	cryptocurrencies = getAfter20RankData(c, cryptocurrencies)

	//5.控制台打印集合
	for _, cryptocurrency := range cryptocurrencies {
		log.Printf("Rank: %d, Name: %s, Symbol: %s, Market Cap: %s, Price: %s, Circulating Supply: %s, Volume 24h: %s, Percent Change 1h: %s, Percent Change 24h: %s, Percent Change 7d: %s",
			cryptocurrency.Rank,
			cryptocurrency.Name,
			cryptocurrency.Symbol,
			cryptocurrency.MarketCap,
			cryptocurrency.Price,
			cryptocurrency.CirculatingSupply,
			cryptocurrency.Volume24h,
			cryptocurrency.PercentChange1h,
			cryptocurrency.PercentChange24h,
			cryptocurrency.PercentChange7d,
		)
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
	if err := c.Visit("https://coinmarketcap.com/all/views/all/"); err != nil {
		log.Printf("visit error:%v", err)
	}

	//6.等待请求结束
	c.Wait()

	//7.返回结构体切片
	return cryptocurrencies
}

// getAfter20RankData 获取后续排行数据
func getAfter20RankData(c *colly.Collector, cryptocurrencies []Cryptocurrency) []Cryptocurrency {

	//1.设置响应回调
	c.OnResponse(func(r *colly.Response) {

		//2.解析JSON到结构体
		var response Response
		err := json.Unmarshal(r.Body, &response)
		if err != nil {
			log.Printf("parse json error:%v", err)
			return
		}

		//3.如果响应异常，打印异常
		if response.Status.ErrorCode != "0" {
			log.Printf("response error:%s", response.Status.ErrorMessage)
			return
		}

		//4.循环封装结构体
		for _, item := range response.Data.CryptoCurrencyList {

			//5.获取报价信息
			quote := item.Quotes[0]

			//6.初始化结构体
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

			//7.存入集合
			cryptocurrencies = append(cryptocurrencies, cryptocurrency)
		}
	})

	//8.开启访问
	if err := c.Visit("https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing?start=21&limit=10&sortBy=market_cap&sortType=desc&convert=USD&cryptoType=all&tagType=all&audited=false"); err != nil {
		log.Printf("visit error:%v", err)
	}

	//9.返回集合
	return cryptocurrencies
}
