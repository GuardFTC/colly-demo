// Package cryptocoins @Author:冯铁城 [17615007230@163.com] 2025-10-21 16:47:25
package cryptocoins

// Cryptocurrency 加密货币数据结构体，对应CoinMarketCap表格中的每一行数据
type Cryptocurrency struct {
	Rank              int    `json:"rank" db:"rank"`                             // 排名：加密货币在市场中的排名，按市值排序
	Name              string `json:"name" db:"name"`                             // 名称：加密货币的全称
	Symbol            string `json:"symbol" db:"symbol"`                         // 符号：加密货币的标准交易符号（如BTC、ETH）
	MarketCap         string `json:"market_cap" db:"market_cap"`                 // 市值：总市值 = 流通供应量 × 当前价格
	Price             string `json:"price" db:"price"`                           // 价格：当前市场价格（通常以USD计价）
	CirculatingSupply string `json:"circulating_supply" db:"circulating_supply"` // 流通供应量：当前在市场上流通的代币数量
	Volume24h         string `json:"volume_24h" db:"volume_24h"`                 // 24小时交易量：过去24小时内的总交易量
	PercentChange1h   string `json:"percent_change_1h" db:"percent_change_1h"`   // 1小时涨跌幅：过去1小时的价格变化百分比（正数上涨，负数下跌）
	PercentChange24h  string `json:"percent_change_24h" db:"percent_change_24h"` // 24小时涨跌幅：过去24小时的价格变化百分比
	PercentChange7d   string `json:"percent_change_7d" db:"percent_change_7d"`   // 7天涨跌幅：过去7天（168小时）的价格变化百分比
}
