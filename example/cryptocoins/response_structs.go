// Package cryptocoins @Author:冯铁城 [17615007230@163.com] 2025-10-21 16:44:56
package cryptocoins

// Response 完整的API响应结构体
type Response struct {
	Data   Data   `json:"data"`   // 数据部分
	Status Status `json:"status"` // 状态部分
}

// Data API响应数据部分
type Data struct {
	CryptoCurrencyList []CryptoCurrencyListItem `json:"cryptoCurrencyList"` // 加密货币列表
	TotalCount         string                   `json:"totalCount"`         // 总记录数(字符串格式)
}

// Status API响应状态部分
type Status struct {
	Timestamp    string `json:"timestamp"`     // API响应时间戳
	ErrorCode    string `json:"error_code"`    // 错误代码(0=成功)
	ErrorMessage string `json:"error_message"` // 错误消息(SUCCESS)
	Elapsed      string `json:"elapsed"`       // 响应耗时(毫秒)
	CreditCount  int    `json:"credit_count"`  // 消耗的API积分
}

// CryptoCurrencyListItem 单个加密货币项
type CryptoCurrencyListItem struct {
	ID                            int64         `json:"id"`                            // 加密货币唯一ID
	Name                          string        `json:"name"`                          // 加密货币名称
	Symbol                        string        `json:"symbol"`                        // 加密货币符号（如SHIB）
	Slug                          string        `json:"slug"`                          // URL友好的标识符
	Tags                          []string      `json:"tags"`                          // 加密货币标签数组
	CMCRank                       int           `json:"cmcRank"`                       // CoinMarketCap排名
	MarketPairCount               int           `json:"marketPairCount"`               // 交易对数量
	CirculatingSupply             float64       `json:"circulatingSupply"`             // 流通供应量
	SelfReportedCirculatingSupply float64       `json:"selfReportedCirculatingSupply"` // 自我报告的流通供应量
	TotalSupply                   float64       `json:"totalSupply"`                   // 总供应量
	MaxSupply                     float64       `json:"maxSupply"`                     // 最大供应量
	IsActive                      int           `json:"isActive"`                      // 是否活跃(1=活跃, 0=不活跃)
	LastUpdated                   string        `json:"lastUpdated"`                   // 最后更新时间
	DateAdded                     string        `json:"dateAdded"`                     // 添加日期
	Quotes                        []Quote       `json:"quotes"`                        // 报价信息数组(不同货币)
	Platform                      *Platform     `json:"platform"`                      // 平台信息(指针，可选)
	IsAudited                     bool          `json:"isAudited"`                     // 是否已审计
	AuditInfoList                 []AuditInfo   `json:"auditInfoList"`                 // 审计信息列表
	Badges                        []interface{} `json:"badges"`                        // 徽章信息(空数组)
}

// Quote 报价信息(不同法币的价格数据)
type Quote struct {
	Name                     string  `json:"name"`                     // 报价货币名称(USD)
	Price                    float64 `json:"price"`                    // 当前价格
	Volume24h                float64 `json:"volume24h"`                // 24小时交易量
	MarketCap                float64 `json:"marketCap"`                // 市值
	PercentChange1h          float64 `json:"percentChange1h"`          // 1小时价格变化百分比
	PercentChange24h         float64 `json:"percentChange24h"`         // 24小时价格变化百分比
	PercentChange7d          float64 `json:"percentChange7d"`          // 7天价格变化百分比
	LastUpdated              string  `json:"lastUpdated"`              // 报价最后更新时间
	PercentChange30d         float64 `json:"percentChange30d"`         // 30天价格变化百分比
	PercentChange60d         float64 `json:"percentChange60d"`         // 60天价格变化百分比
	PercentChange90d         float64 `json:"percentChange90d"`         // 90天价格变化百分比
	FullyDilutedMarketCap    float64 `json:"fullyDilluttedMarketCap"`  // 完全稀释市值
	MarketCapByTotalSupply   float64 `json:"marketCapByTotalSupply"`   // 按总供应量计算的市值
	Dominance                float64 `json:"dominance"`                // 市场主导率
	Turnover                 float64 `json:"turnover"`                 // 周转率
	YTDPriceChangePercentage float64 `json:"ytdPriceChangePercentage"` // 年初至今价格变化百分比
	PercentChange1y          float64 `json:"percentChange1y"`          // 1年价格变化百分比
}

// Platform 区块链平台信息
type Platform struct {
	ID           int    `json:"id"`            // 平台ID
	Name         string `json:"name"`          // 平台名称(Ethereum)
	Symbol       string `json:"symbol"`        // 平台符号(ETH)
	Slug         string `json:"slug"`          // 平台URL标识
	TokenAddress string `json:"token_address"` // 代币合约地址
}

// AuditInfo 审计信息项
type AuditInfo struct {
	CoinID      string `json:"coinId"`      // 币种ID
	Auditor     string `json:"auditor"`     // 审计机构(CertiK)
	AuditStatus int    `json:"auditStatus"` // 审计状态(2=已完成)
	AuditTime   string `json:"auditTime"`   // 审计时间
	ReportURL   string `json:"reportUrl"`   // 审计报告URL
}
