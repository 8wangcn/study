package main

import (
	"fmt"
	"time"
)

/*
interface
语法

	type 接口名 interface {
		方法1(参数列表) 返回值列表
		方法2(参数列表) 返回值列表
	}
*/
type Payment interface {
	Pay(amount float64) (string, error)
	Refund(transactionID string, amount float64) (string, error)
	Query(transactionID string) (string, error)
}

type Alipay struct {
	AppID     string
	AppSecret string
	MerchID   string
}

func (a Alipay) Pay(amount float64) (string, error) {
	return fmt.Sprintf("ALIPAY-%d", time.Now().Unix()), nil
}

func (a Alipay) Refund(transactionID string, amount float64) (string, error) {
	return fmt.Sprintf("ALIPAY-REFUND-%d", transactionID), nil
}

func (a Alipay) Query(transactionID string) (string, error) {
	return "交易完成", nil
}

type WechatPay struct {
	AppID     string
	AppSecret string
	MerchID   string
}

func (a WechatPay) Pay(amount float64) (string, error) {
	return fmt.Sprintf("WECHAT-%d", time.Now().Unix()), nil
}

func (a WechatPay) Refund(transactionID string, amount float64) (string, error) {
	return fmt.Sprintf("WECHAT-REFUND-%d", transactionID), nil
}

func (a WechatPay) Query(transactionID string) (string, error) {
	return "交易完成", nil
}

func ProcessPayment(p Payment, amount float64) (string, error) {
	fmt.Println("开始处理支付请求....")
	transactionID, err := p.Pay(amount)
	if err != nil {
		return "", err
	}
	fmt.Printf("支付成功 交易号:%s\n", transactionID)
	return transactionID, nil
}

func main() {
	alipay := Alipay{
		AppID:     "ali_id_12245",
		AppSecret: "salkjdf",
		MerchID:   "123",
	}

	wechat := WechatPay{
		AppID:     "wechat_id 35406",
		AppSecret: "123456",
		MerchID:   "123",
	}

	transactions := []struct {
		Payment Payment
		Amount  float64
		Name    string
	}{
		{alipay, 100, "支付宝"},
		{wechat, 200, "微信"},
	}

	for _, pp := range transactions {
		transactionID, err := ProcessPayment(pp.Payment, pp.Amount)
		if err != nil {
			fmt.Printf("支付失败 错误信息:%S \n", err.Error())
			continue
		}
		fmt.Printf("支付成功 交易号:%s \n", transactionID)

	}
}
