package client

import (
	"strings"

	"github.com/yzlq99/eastmoneyapi/model"
)

// GetCanBuyNewStockList 查询可申请新股列表
func (e *EastMoneyClient) GetCanBuyNewStockList() (*model.StockList, error) {
	req, _ := createRequestWithBaseHeader(
		"POST",
		baseUrl+"/Trade/GetCanBuyNewStockListV3?validatekey="+e.validateKey,
		nil)
	resp, err := e.cli.Do(req)
	if err != nil {
		return nil, err
	}
	result := model.StockList{}
	if err := bindJson(resp.Body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetConvertibleBondList 查询新债列表
func (e *EastMoneyClient) GetNewConvertibleBondList() (*model.ConvertibleBondList, error) {
	req, _ := createRequestWithBaseHeader(
		"POST",
		baseUrl+"/Trade/GetConvertibleBondListV2?validatekey="+e.validateKey,
		nil)
	resp, err := e.cli.Do(req)
	if err != nil {
		return nil, err
	}
	result := model.ConvertibleBondList{}
	if err := bindJson(resp.Body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SubmitBatTrade 申购
func (e *EastMoneyClient) SubmitBatTrade(params model.SubmitBatTradeParams) (*model.SubmitBatTradeResult, error) {
	body, err := params.ToJson()
	if err != nil {
		return nil, err
	}
	req, _ := createRequestWithJson(
		"POST",
		baseUrl+"/Trade/SubmitBatTradeV2?validatekey="+e.validateKey,
		strings.NewReader(string(body)))
	resp, err := e.cli.Do(req)
	if err != nil {
		return nil, err
	}
	result := model.SubmitBatTradeResult{}
	if err := bindJson(resp.Body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
