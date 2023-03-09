package model

import (
	"encoding/json"
	"strconv"
	"strings"
)

type SubmitBatTradeParam struct {
	StockCode string `json:"StockCode"`
	StockName string `json:"StockName"`
	Price     string `json:"Price"`
	Amount    int    `json:"Amount"`
	TradeType string `json:"TradeType"`
	Market    string `json:"Market"`
}

type SubmitBatTradeParams []SubmitBatTradeParam

func (s SubmitBatTradeParams) ToJson() ([]byte, error) {
	return json.Marshal(s)
}

type SubmitBatTradeResult struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
	// TODO: 完善模型
	Data []interface{} `json:"Data"`
}

/************** 新股 *******************/
type StockList struct {
	Kyzj         float64    `json:"Kyzj"`
	Zjye         string     `json:"Zjye"`
	NewQuota     []NewQuota `json:"NewQuota"`
	NewStockList []string   `json:"NewStockList"`
}

func (s StockList) GetSubmitBatTradeParams() SubmitBatTradeParams {
	res := SubmitBatTradeParams{}
	for i := range s.NewStockList {
		// TODO: 修改参数
		arr := strings.Split(s.NewStockList[i], ",")
		// stockCode := arr[1]
		stockName := arr[2]
		submitCode := arr[3]
		res = append(res, SubmitBatTradeParam{
			StockCode: submitCode,
			StockName: stockName,
			// Price:     d.PARVALUE,
			// Amount:    amount,
			TradeType: "B",
			// Market:    d.Market,
		})
	}
	return res
}

type NewQuota struct {
	Gddm    string `json:"Gddm"`
	Kcbsged string `json:"Kcbsged"`
	Ksgsz   string `json:"Ksgsz"`
	Market  string `json:"Market"`
}

/************** 新债 *******************/
type ConvertibleBondList struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
	Errcode int    `json:"Errcode"`
	Data    []Data `json:"Data"`
}

func (c ConvertibleBondList) GetSubmitBatTradeParams() SubmitBatTradeParams {
	res := SubmitBatTradeParams{}
	for i := range c.Data {
		d := c.Data[i]
		amount, _ := strconv.Atoi(d.LIMITBUYVOL)
		res = append(res, SubmitBatTradeParam{
			StockCode: d.SUBCODE,
			StockName: d.SUBNAME,
			Price:     d.PARVALUE,
			Amount:    amount,
			TradeType: "B",
			Market:    d.Market,
		})
	}
	return res
}

type Data struct {
	ExStatus      int         `json:"ExStatus"`
	ExIsToday     bool        `json:"ExIsToday"`
	BONDCODE      string      `json:"BONDCODE"`
	BONDNAME      string      `json:"BONDNAME"`
	BUYREDEMCODE  string      `json:"BUYREDEMCODE"`
	BUYREDEMNAME  string      `json:"BUYREDEMNAME"`
	BUYREDEMPRICE string      `json:"BUYREDEMPRICE"`
	CREDITRATING  string      `json:"CREDITRATING"`
	Cybbz         string      `json:"Cybbz"`
	FLOORBUYVOL   string      `json:"FLOORBUYVOL"`
	ISSUESDATE    string      `json:"ISSUESDATE"`
	ISSUEVOL      string      `json:"ISSUEVOL"`
	LIMITBUYVOL   string      `json:"LIMITBUYVOL"`
	LWRANDATE     string      `json:"LWRANDATE"`
	Market        string      `json:"Market"`
	PARVALUE      string      `json:"PARVALUE"`
	PLACINGCODE   string      `json:"PLACINGCODE"`
	PLACINGNAME   string      `json:"PLACINGNAME"`
	PLACINGRIGHT  interface{} `json:"PLACINGRIGHT"`
	PURCHASEDATE  string      `json:"PURCHASEDATE"`
	RATING        string      `json:"RATING"`
	SFZCZ         string      `json:"SFZCZ"`
	SUBCODE       string      `json:"SUBCODE"`
	SUBNAME       string      `json:"SUBNAME"`
	SWAPPRICE     string      `json:"SWAPPRICE"`
	SWAPSCODE     string      `json:"SWAPSCODE"`
	SWAPSNAME     interface{} `json:"SWAPSNAME"`
	SWAPVALUE     string      `json:"SWAPVALUE"`
	BONDTYPECODE  string      `json:"BONDTYPECODE"`
	Zqzwqc        string      `json:"zqzwqc"`
}
