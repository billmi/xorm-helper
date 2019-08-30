package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/billmi/xorm-pagenation/pagenation"
	"fmt"
)

/**
	datetime : 2019-08-30 18:18:18
	author : Bill
 */

func main() {
	engine, err := xorm.NewEngine("mysql",
		"root"+":"+"root"+"@tcp("+"127.0.0.1"+":"+"3306"+")/"+"bill"+"?charset=utf8",
	)
	if err != nil {
		fmt.Print(err.Error())
	}
	var PageHelper = pagenation.DaoBase{}

	//Set xorm engine
	PageHelper.SetDatasource(engine)

	//Use Data po struct
	var po = []DeviceOauthLogPo{}
	PageHelper.GetLists(&po, "`test`", "*", "id", "", "", "", "`id` DESC", "")
	if len(po) > 0 {
		for _, row := range po {
			fmt.Print(row.DeviceTypeName)
			fmt.Print(row.DeviceTypeTitle)
		}
	}
}
