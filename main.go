package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
)

/**
	datetime : 2019-03-12 18:18:18
	author : Bill
 */


//po base database => xorm
type DeviceOauthLogPo struct {
	DeviceType         int    `json:"device_type"`
	DeviceTypeName     string `json:"device_type_name"`
	DeviceTypeTitle    string `json:"device_type_title"`
	DeviceScreenWidth  int    `json:"device_screen_width"`
	DeviceScreenHeight int    `json:"device_screen_height"`
}

func main() {
	engine, err := xorm.NewEngine( "mysql",
		"root" + ":" + "root" + "@tcp("+"127.0.0.1"+":"+ "3306" +")/"+ "bill"+"?charset=utf8",
	)
	if err != nil{
		fmt.Print(err.Error())
	}
	var ListHelper = pagenation.DaoBase{}
	ListHelper.SetDatasource(engine)
	var po = []*DeviceOauthLogPo{}
	ListHelper.GetLists(&po,"`test`","`id`","","`id` DESC")
	if len(po) > 0{
		for _,row := range po{
			fmt.Print(row.DeviceTypeName)
			fmt.Print("\r\n ========= \r\n")
			fmt.Print(row.DeviceTypeTitle)
			fmt.Print("\r\n ========= \r\n")
		}
	}

	var po1 = []*DeviceOauthLogPo{}
	var pageListInfo = ListHelper.GetPageLists(&po1,"`test`","`id`","","`id` DESC",0,2)
	if len(po) > 0{
		fmt.Print(pageListInfo)    //result : map[string]interface{} 请阅读源码
	}

	fmt.Print("\r\n ========== condition Build \r\n")

	//condition Build  后面可以自己扩展
	var condi = make(map[string]map[string]interface{})
	var inCodi = make(map[string]interface{})
	var like = make(map[string]interface{})
	like["name"] = "Bill"
	inCodi["type"] = 1
	inCodi["title"] = "Bill"
	condi["AND"] = inCodi
	condi["LIKE"] = like
	fmt.Print(ListHelper.ConditionBuild(condi))
}