package model

import (
	"encoding/json"
	"strconv"
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
	NewQuota     []NewQuota     `json:"NewQuota"`
	NewStockList []NewStockList `json:"NewStockList"`
}

type NewQuota struct {
	Gddm    string `json:"Gddm"`
	Kcbsged string `json:"Kcbsged"`
	Ksgsz   string `json:"Ksgsz"`
	Market  string `json:"Market"`
}

type NewStockList struct {
	Market     string      `json:"Market"`
	Sgrq       string      `json:"Sgrq"` // 申购日期 20231121
	Zqdm       string      `json:"Zqdm"` // 证券代码
	Zqmc       string      `json:"Zqmc"`
	Sgdm       string      `json:"Sgdm"` // 申购代码
	Fxzs       string      `json:"Fxzs"`
	Wsfxs      string      `json:"Wsfxs"`
	Fxj        string      `json:"Fxj"` // 申购价格
	YcFxj      string      `json:"Yc_Fxj"`
	Sgsx       string      `json:"Sgsx"` // 申购最大数量
	YcSgsx     string      `json:"Yc_Sgsx"`
	Sgzjsx     string      `json:"Sgzjsx"`
	YcSgzjsx   string      `json:"Yc_Sgzjsx"`
	Ksgsx      string      `json:"Ksgsx"` // 可申购数量
	SgState    string      `json:"SgState"`
	MinStep    string      `json:"Min_Step"`
	CDRFlag    string      `json:"CDR_Flag"`
	YLFlag     interface{} `json:"YL_Flag"`
	TPQCYFLag  string      `json:"TPQCY_FLag"`
	Cybbz      string      `json:"Cybbz"`
	SFZCZ      string      `json:"SFZCZ"`
	JYXYJG     string      `json:"JYXYJG"`
	APPLYPRICE string      `json:"APPLYPRICE"`
	Zqzwqc     string      `json:"zqzwqc"`
}

func (s StockList) GetSubmitBatTradeParams() SubmitBatTradeParams {
	res := SubmitBatTradeParams{}
	for i := range s.NewStockList {
		stock := s.NewStockList[i]
		amount, _ := strconv.Atoi(stock.Ksgsx)
		res = append(res, SubmitBatTradeParam{
			StockCode: stock.Sgdm,
			StockName: stock.Zqmc,
			Price:     stock.Fxj,
			Amount:    amount,
			TradeType: "B",
			Market:    stock.Market,
		})
	}
	return res
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
