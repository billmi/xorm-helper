baseFrom : go-xorm engine

##### 1. get xorm
`
go get -u -v github.com/go-xorm/xorm
`

#####  2. get mysql device
`
go get -u -v github.com/go-sql-driver/mysql
`

#####  3. import test.sql to database

##### MAIN

```
package main

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"bill/src/pagenation"
)

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
		fmt.Print(pageListInfo)    //result : map[string]interface{}【 You can edit in it ! 】
	}

	fmt.Print("\r\n ========== condition Build \r\n")

	//condition Build  【Can add by yourself !】
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
```