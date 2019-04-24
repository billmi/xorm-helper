package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm-pagenation/src/pagenation"
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

type DeviceOauthLogPo1 struct {
	DeviceType         int    `json:"device_type"`
	DeviceTypeName     string `json:"device_type_name"`
	DeviceTypeTitle    string `json:"device_type_title"`
	DeviceScreenWidth  int    `json:"device_screen_width"`
	DeviceScreenHeight int    `json:"device_screen_height"`
	Title              string `json:"title"`
}

func main() {
	engine, err := xorm.NewEngine( "mysql",
		"root" + ":" + "root" + "@tcp("+"127.0.0.1"+":"+ "3306" +")/"+ "bill"+"?charset=utf8",
	)
	if err != nil{
		fmt.Print(err.Error())
	}
	var PageHelper = pagenation.DaoBase{}
	PageHelper.SetDatasource(engine)
	var po = []*DeviceOauthLogPo{}
	PageHelper.GetLists(&po,"`test`","`id`","","`id` DESC")
	if len(po) > 0{
		for _,row := range po{
			fmt.Print(row.DeviceTypeName)
			fmt.Print("\r\n ========= \r\n")
			fmt.Print(row.DeviceTypeTitle)
			fmt.Print("\r\n ========= \r\n")
		}
	}

	var po1 = []*DeviceOauthLogPo1{}
	var po1Join = [][]string{
		{"INNER","test1 b","b.id = a.r_id"},
	}
	var pageListInfo = PageHelper.GetPageLists(&po1,"test","a.*,b.title","a.id","a",PageHelper.ConditionJoin(po1Join),"","a.id DESC",0,2)
	fmt.Printf("\r\n total_page : %d \r\n",pageListInfo["total_page"])
	fmt.Printf("\r\n curr_page : %d \r\n",pageListInfo["curr_page"])
	fmt.Printf("\r\n page_rows : %d \r\n",pageListInfo["page_rows"])
	fmt.Printf("\r\n total_record : %d \r\n",pageListInfo["total_record"])
	for _,row := range po1{
		fmt.Printf("\r\n relation_title : %s \r\n",row.Title)
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
	fmt.Print(PageHelper.ConditionBuild(condi))

	//join build , -- join生成器
	var data = [][]string{
		{"INNER","b b","b.id = a.id"},
		{"INNER","c c","c.id = b.id"},
	}
	fmt.Print("\r\n ========== Join Conditon Build \r\n")
	fmt.Print(PageHelper.ConditionJoin(data))
}