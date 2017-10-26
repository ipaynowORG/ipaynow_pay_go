package ipaynow_pay

import (
	"fmt"

	git "github.com/ipaynowORG/ipaynow_pay_go/ipaynow_pay"
)

func main() {

	app := git.App{
		appId:  "150753082825470",
		appKey: "8jTST7ywIBY0QQ3RlcxWEl08Xj9gaYyQ",
	}
	orderDetail := git.OrderDetail{
		mhtOrderName:   "测试商品",
		mhtOrderDetail: "的描述",
		mhtOrderAmt:    1,
		mhtGoodsTag:    "",
	}

	content := git.wx_scan_05(&app, &orderDetail, "", "https://op-tester.ipaynow.cn/paytest/notify", "135604611220313639")
	fmt.Println(content)

	//	resp := git.queryOrder("YEIehmWjkRzrR", &app, "05")
	//	fmt.Println(resp)
}
