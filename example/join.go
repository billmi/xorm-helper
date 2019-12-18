package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/billmi/xorm-helper"
	"fmt"
)

//参考 pagenation-list.go

func main() {

	engine, err := xorm.NewEngine("mysql",
		"root"+":"+"root"+"@tcp("+"127.0.0.1"+":"+"3306"+")/"+"bill"+"?charset=utf8",
	)
	if err != nil {
		fmt.Print(err.Error())
	}
	var XormHelper = xormhelper.XormHelper{}

	//Set xorm engine
	XormHelper.SetDatasource(engine)

}
