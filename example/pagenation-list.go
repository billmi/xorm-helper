package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm-pagenation/pagenation"
	"fmt"
)

/**
	datetime : 2019-08-30 18:18:18
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

//relation po test => test1
type DeviceOauthLogPo1 struct {
	DeviceType         int    `json:"device_type"`
	DeviceTypeName     string `json:"device_type_name"`
	DeviceTypeTitle    string `json:"device_type_title"`
	DeviceScreenWidth  int    `json:"device_screen_width"`
	DeviceScreenHeight int    `json:"device_screen_height"`
	Title              string `json:"title"` //test1 .title
}

func main() {
	engine, err := xorm.NewEngine("mysql",
		"root"+":"+"root"+"@tcp("+"127.0.0.1"+":"+"3306"+")/"+"bill"+"?charset=utf8",
	)
	if err != nil {
		fmt.Print(err.Error())
	}
	var (
		PageHelper = pagenation.DaoBase{}

		//Use Data po struct
		po = []*DeviceOauthLogPo1{}

		//if not join  can be set empty string  => ""
		po1Join = [][]string{
			{"INNER", "test1 b", "b.id = a.r_id"},
		}

		//if no condition can be set empty string  => ""
		condition = ""

		page      = 1  //default

		listRow   = 15 //default

		group     = ""
		order     = "a.id DESC"
		alias     = "a"
		pk        = "a.id"
		tableName = "test"

		//field Select
		fields = "a.*,b.title"
	)

	//Set xorm engine
	PageHelper.SetDatasource(engine)

	pageListInfo := PageHelper.GetPageLists(&po, tableName, fields, pk, alias, PageHelper.ConditionJoin(po1Join), condition, order, group, page, listRow)

	//Can Get PageInfo
	fmt.Printf("\r\n total_page : %d \r\n", pageListInfo["total_page"])
	fmt.Printf("\r\n curr_page : %d \r\n", pageListInfo["curr_page"])
	fmt.Printf("\r\n page_rows : %d \r\n", pageListInfo["page_rows"])
	fmt.Printf("\r\n total_record : %d \r\n", pageListInfo["total_record"])

	//Data in po
	for _, row := range po {
		fmt.Printf("\r\n relation_title : %s \r\n", row.Title)
	}
}
