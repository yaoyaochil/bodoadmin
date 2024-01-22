package passport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/util"
	"mime/multipart"
)

type GagService struct{}

type GetShopEntityCheckResponse struct {
	Status string `json:"status"`
	Data   struct {
		Data []struct {
			ContractType     string `json:"contractType"`
			IsTerminal       int    `json:"isTerminal"`
			EntityId         string `json:"entityId"`
			UpdateTime       string `json:"updateTime"`
			Operator         string `json:"operator"`
			EntityStatus     string `json:"entityStatus"`
			CreateTime       string `json:"createTime"`
			EntityName       string `json:"entityName"`
			TerminalStatus   string `json:"terminalStatus"`
			IsClaimDeviation int    `json:"isClaimDeviation"`
			IsAppendPrint    int    `json:"isAppendPrint"`
			ShopType         string `json:"shopType"`
			IsClaim          int    `json:"isClaim"`
		} `json:"data"`
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		PageTotal int `json:"pageTotal"`
		Total     int `json:"total"`
	} `json:"data"`
	Msg     string `json:"msg"`
	LogUUID string `json:"logUUID"`
}

// GetShopEntityCheck 获取门店考核列表
func (g *GagService) GetShopEntityCheck() (data GetShopEntityCheckResponse, err error) {
	url := "http://databiapi.web.pri.xincheng.com/shopEntityCheck/listShopEntityCheck?token=misUser1H8GT4I9LVJ0I1002VT05VFO88M84JR9"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("id", "1DQGAF26VB203C0AB2M10187U00017VS")
	_ = writer.WriteField("pageIndex", "1")
	_ = writer.WriteField("pageSize", "999")
	_ = writer.WriteField("entityStatus", "U")
	_ = writer.WriteField("contractType", "Z001")
	_ = writer.WriteField("terminalStatus", "0")
	_ = writer.WriteField("orderBy", "desc")
	_ = writer.WriteField("orderKey", "createTime")
	_ = writer.WriteField("shopEntityId", "")
	_ = writer.WriteField("shopTypeId", "")
	err = writer.Close()
	if err != nil {
		return
	}
	request, err := util.SeazenGagRequest(url, method, payload, nil, writer)
	if err != nil {
		return
	}
	err = json.Unmarshal(request, &data)
	if err != nil {
		return
	}
	return
}

type GetShopEntitySaleInfoResponse struct {
	Status string `json:"status"`
	Data   struct {
		Data []struct {
			RefundSaleCount       float64 `json:"refundSaleCount"`
			ConfirmAmount         float64 `json:"confirmAmount"`
			PBrand                string  `json:"pBrand"`
			ChargeType            string  `json:"chargeType"`
			ClaimBillTotal        float64 `json:"claimBillTotal"`
			ChannelCount          float64 `json:"channelCount"`
			ClaimDeviationRate    float64 `json:"claimDeviationRate"`
			PEfficiency           float64 `json:"pEfficiency"`
			ShowId                string  `json:"showId"`
			NetSale               float64 `json:"netSale"`
			EntityName            string  `json:"entityName"`
			TerminalStatus        string  `json:"terminalStatus"`
			DeviationCheck        string  `json:"deviationCheck"`
			RecommendClaimDiffAbs float64 `json:"recommendClaimDiffAbs"`
			SaleAmount            float64 `json:"saleAmount"`
			SaleCount             float64 `json:"saleCount"`
			RecommendAmount       float64 `json:"recommendAmount"`
			EntityId              string  `json:"entityId"`
			LeasingResource       string  `json:"leasingResource"`
			DailyBillCount2       float64 `json:"dailyBillCount2"`
			Storey                string  `json:"storey"`
			ChannelAmount         float64 `json:"channelAmount"`
			RefundSaleAmount      float64 `json:"refundSaleAmount"`
			ClaimMoneyTotal       float64 `json:"claimMoneyTotal"`
			CheckAmount           float64 `json:"checkAmount"`
			EntityStatus          string  `json:"entityStatus"`
			BillAvg               float64 `json:"billAvg"`
			EntityTypeRoot        string  `json:"entityTypeRoot"`
			IsDel                 string  `json:"isDel"`
			DailyBillAmount       float64 `json:"dailyBillAmount"`
			PEfficiencyAvg        float64 `json:"pEfficiencyAvg"`
			MerchantCode          string  `json:"merchantCode,omitempty"`
		} `json:"data"`
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		PageTotal int `json:"pageTotal"`
		Total     int `json:"total"`
	} `json:"data"`
	Msg     string `json:"msg"`
	LogUUID string `json:"logUUID"`
}

// GetShopEntitySaleInfo 获取考核店铺差异金额
func (g *GagService) GetShopEntitySaleInfo() (data GetShopEntitySaleInfoResponse, err error) {
	url := "http://databiapi.web.pri.xincheng.com/shop/shopAnalyze/getShopEntitySaleInfo?token=misUser1EH019V4E6OSNK7Q2OVBQR9N950012JE"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("from", util.GetYearFirstDay())
	_ = writer.WriteField("to", util.GetYesterday())
	_ = writer.WriteField("pageIndex", "1")
	_ = writer.WriteField("pageSize", "999")
	_ = writer.WriteField("type", "0")
	_ = writer.WriteField("id", "1DQGAF26VB203C0AB2M10187U00017VS")
	_ = writer.WriteField("orderKey", "recommendClaimDiffAbs")
	_ = writer.WriteField("effectiveTerminalFlag", "1")
	_ = writer.WriteField("deviationCheck", "1")
	_ = writer.WriteField("terminalStatus", "")
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	request, err := util.SeazenGagRequest(url, method, payload, nil, writer)
	if err != nil {
		return
	}
	err = json.Unmarshal(request, &data)
	if err != nil {
		return
	}
	return
}
