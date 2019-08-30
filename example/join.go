package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm-pagenation/pagenation"
	"fmt"
)

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


}
