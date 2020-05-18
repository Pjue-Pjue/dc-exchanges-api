package okex

import (
	dc_exchanges_api "dc-exchanges-api"
	"strconv"
)

type AccountOKEXClient struct {
	client *dc_exchanges_api.Client
}

func NewAccountOKEXClient(apiKey, secretKey, passphrase string, endPoint string) *AccountOKEXClient {
	return &AccountOKEXClient{client: newOKExClient(apiKey, secretKey, passphrase, endPoint)}
}

// 获取账户资产估值
func (c *AccountOKEXClient) GetAssetValuation(accountType int, valuationCurrency string) (AssetValuation, error) {
	r := AssetValuation{}
	params := map[string]string{}
	if accountType >= 0 {
		params["account_type"] = strconv.Itoa(accountType)
	}
	if valuationCurrency != "" {
		params["valuation_currency"] = valuationCurrency
	}

	uri := BuildParams(TOTAL_ASSEST,params)
	if _, _, err := c.client.Request("GET", uri, nil, &r); err != nil {
		return AssetValuation{}, err
	}
	return r,nil
}

// 转移资产
// type [0:母/子账户中各个账户之间划转] [1:母账户转子账号] [2:子账户转母账号] 默认为0
// 转入、转出 1-币币账户 3-交割合约 4-法币账户 5-币币杠杆账户 6-资金账户 8-余币宝账户 9-永续合约账户 12-期权合约 14-挖矿账户 17-借贷账户
// sub_account 子账号登录名，type为1、2、时，sub_account为必填项
//当from为3或5或9时，instrument_id为必填参数
//当to为3或5或9时，to_instrument_id为必填参数
func (c *AccountOKEXClient)Transfer(currency string, from, to int32, amount float32, optionalParams map[string]string) ([]byte, TransferResult, error) {
	var result TransferResult
	transferInfo := map[string]interface{}{}
	transferInfo["currency"] = currency
	transferInfo["from"] = from
	transferInfo["to"] = to
	transferInfo["amount"] = amount

	if optionalParams != nil && len(optionalParams) > 0 {
		if from == 0 || to == 0 {
			transferInfo["sub_account"] = optionalParams["sub_account"]
		}
		transferInfo["instrument_id"] = optionalParams["instrument_id"]
		if to == 5 || to == 3 { // 杠杆转入币对 / usdt保证金合约underlying
			transferInfo["to_instrument_id"] = optionalParams["to_instrument_id"]
		}
	}
	respBody, _, err := c.client.Request("POST",TRANSFER, transferInfo, &result)
	return respBody, result, err
}