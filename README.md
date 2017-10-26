
# 聚合支付SDK使用说明 #

## 版本变更记录 ##


- 1.0.0 : 初稿

## 1. 概述 ##

### 1.1 简介 ###

- 聚合支付SDK。

### 1.2 如何获取 ###

[获取源码](https://github.com/ipaynowORG/ipaynow_pay_go)

[demo源码](https://github.com/ipaynowORG/ipaynow_pay_go)



## 2. API ##

使用"go get github.com/ipaynowORG/ipaynow_pay_go/ipaynow_pay" 命令clone 并 install ipaynow_pay模块

代码中 import (git "github.com/ipaynowORG/ipaynow_pay_go/ipaynow_pay")使用

- 微信被扫支付

        /**
         * 微信被扫支付
         * @param app appId(应用ID)和appKey ,
         * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
         * @param mhtSubAppId 微信子号对应多个公众号的时候必填,如果只对应一个公众号则不传
         * @param notifyUr 后台通知地址
         * @param channelAuthCode 支付码
         */
        func wx_scan_05(app *App, orderDetail *OrderDetail, mhtSubAppId string, notifyUrl string, channelAuthCode string) string

- 支付宝被扫支付

        /**
         * 支付宝被扫支付
         * @param app appId(应用ID)和appKey ,
         * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
         * @param notifyUr 后台通知地址
         * @param channelAuthCode 支付码
         */
        func ali_scan_05(app *App, orderDetail *OrderDetail, notifyUrl string, channelAuthCode string) string



- 微信主扫支付

        /**
         * 微信主扫支付
         * @param app appId(应用ID)和appKey ,
         * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
         * @param mhtSubAppId 微信子号对应多个公众号的时候必填,如果只对应一个公众号则不传
         * @param notifyUrl 后台通知地址
         */
        func wx_scan_08(app *App, orderDetail *OrderDetail, mhtSubAppId string, notifyUrl string) string


- 支付宝主扫支付

        /**
         * 支付宝主扫支付
         * @param app appId(应用ID)和appKey ,
         * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
         * @param notifyUrl 后台通知地址
         */
        func ali_scan_08(app *App, orderDetail *OrderDetail, notifyUrl string) string



- 微信公众号支付

        /**
         * 微信公众号支付
         * @param app appId(应用ID)和appKey ,
         * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
         * @param notifyUrl 后台通知地址
         * @param frontNotifyUrl 前台页面跳转地址
         */
        func wx_p_account(app *App, orderDetail *OrderDetail, notifyUrl string, frontNotifyUrl string) string



- 支付宝公众号支付

        /**
         * 支付宝公众号支付
         * @param app appId(应用ID)和appKey ,
         * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
         * @param notifyUrl 后台通知地址
         * @param frontNotifyUrl 前台页面跳转地址
         */
        func ali_p_account(app *App, orderDetail *OrderDetail, notifyUrl string, frontNotifyUrl string) string



- 微信H5

        /**
         * 微信H5
         * @param app appId(应用ID)和appKey ,
         * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
         * @param consumerCreateIp 用户支付IP
         * @param notifyUrl 后台通知地址
         * @param frontNotifyUrl 前台页面跳转地址
         */
        func wx_h5(app *App, orderDetail *OrderDetail, consumerCreateIp string, notifyUrl string, frontNotifyUrl string) string


- 支付宝H5

        /**
         * 支付宝H5
         * @param app appId(应用ID)和appKey ,
         * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
         * @param notifyUrl 后台通知地址
         * @param frontNotifyUrl 前台页面跳转地址
         */
        func ali_h5(app *App, orderDetail *OrderDetail, notifyUrl string, frontNotifyUrl string) string 


- 支付宝网页web

        /**
         * 支付宝网页web
         * @param app appId(应用ID)和appKey ,
         * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
         * @param notifyUrl 后台通知地址
         */
        func ali_web(app *App, orderDetail *OrderDetail, notifyUrl string) string 



- 微信小程序支付

        /**
         * 微信小程序支付
         * @param app appId(应用ID)和appKey ,
         * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param orderDetail   商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动),
         * @param consumerId  用户openId
         * @param notifyUrl 后台通知地址
         */
        func wx_app(app *App, orderDetail *OrderDetail, consumerId string, notifyUrl string) string



- 商户支付订单查询

        /**
         * 商户支付订单查询
         * @param mhtOrderNo    商户订单号
         * @param appId 商户的AppId,https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param appKey 商户的AppKey,https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param deviceType    被扫05，主扫08，公众号传0600，h5传0601，网页04
         */
        func queryOrder(mhtOrderNo string, app *App, deviceType string) string 



- 退款

        /**
         * 退款
         * @param appId 商户的AppId,https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param appKey 商户的AppKey,https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param mhtOrderNo    商户订单号
         * @param amount    退款金额
         * @param reason    退款原因
         */
        func refundOrder(app *App, mhtOrderNo string, amount int, reason string) string


- 退款查询

        /**
         * 退款查询
         * @param appId 商户的AppId,https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param appKey 商户的AppKey,https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param mhtRefundNo   商户退款单号
         */
        func refundQuery(app *App, mhtRefundNo string) string

- 撤销

        /**
         * 撤销(只能撤销当天的交易,且无论成功失败(逻辑包含退款))
         * @param appId 商户的AppId,https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param appKey 商户的AppKey,https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param mhtOrderNo    商户订单号
         * @param reason    退款原因
         */
        func backOrder(app *App, mhtOrderNo string, reason string) string

- 撤销查询

        /**
         * 撤销查询
         * @param appId 商户的AppId,https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param appKey 商户的AppKey,https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
         * @param mhtRefundNo   商户退款单号
         */
        func backQuery(app *App, mhtRefundNo string) string 


## 3. DEMO说明 ##

	package main
	
	import (
		"fmt"
		//使用go get github.com/ipaynowORG/ipaynow_pay_go/ipaynow_pay 命令clone 并 install ipaynow_pay模块
		git "github.com/ipaynowORG/ipaynow_pay_go/ipaynow_pay"
	)
	
	func main() {
	
		app := git.App{
			AppId:  "xxxxxxxxxxxxxx",
			AppKey: "xxxxxxxxxxxxxxxxxxxxxxxxxx",
		}
		orderDetail := git.OrderDetail{
			MhtOrderName:   "测试商品",
			MhtOrderDetail: "的描述",
			MhtOrderAmt:    1,
			MhtGoodsTag:    "",
		}
	
		//主扫测试
		content := git.Wx_scan_05(&app, &orderDetail, "", "https://xxxxxxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxx")
		fmt.Println(content)
	
		//商户字符订单查询
		//	resp := git.QueryOrder("xxxxxxxxxxxxxx", &app, "05")
		//	fmt.Println(resp)
	}            