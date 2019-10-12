package xormhelper

import (
	"math"
	"fmt"
)

const (
	_pkField = "id"
	countField = "count"
	ListRow = 15
	LIST_ROWS = 15
)

var (
	_innerFlag = "INNER"
	_leftFlag  = "LEFT"
	_rightFlag = "RIGHT"


	innerTpl = " INNER JOIN %s ON %s "
	leftTpl = " LEFT JOIN %s ON %s "
	rightTpl = " RIGHT JOIN %s ON %s "


	_andFlag     = "AND"
	_likeFlag    = "LIKE"
	_gtFlag      = "GT"
	_ltFlag      = "LT"
	_InFlag      = "IN"
	_nullFlag    = "NULL"
	_orFlag      = "OR"
	_betweenFlag = "BETWEEN"
	_expFlag     = "EXP"

	_condi        = ""
	str           = " AND ( %s = '%v' )"
	intCondi      = " AND ( %s = %v )"
	_gtCondi      = " AND ( %s > %v )"
	_ltCondi      = " AND ( %s < %v )"
	_InCondi      = " AND ( %s IN (%s) )"
	_NullCondi    = " AND ( ISNULL(%s) )"
	_orCondi      = " OR ( %v )"
	_betweenCondi = " AND ( %s BETWEEN %s )"

	_expCondi = " AND ( %s %s )"
	_currRel  = ""
)

/**
	ResPageInfo [same]
	@author Bill
 */
func CommaPaginator(page int, listRow int, total int64) map[string]int {
	totalpages := int(math.Ceil(float64(total) / float64(listRow)))
	if page <= 0 {
		page = 1
	}
	paginator := make(map[string]int)
	paginator["total_page"] = totalpages
	paginator["curr_page"] = page
	paginator["page_rows"] = listRow
	return paginator
}

/**
	Paginator [same]
 */
func Paginator(page int, listRow int, total int64) map[string]interface{} {
	totalpages := int(math.Ceil(float64(total) / float64(listRow)))
	if page <= 0 {
		page = 1
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["totalPage"] = totalpages
	paginatorMap["currpage"] = page
	paginatorMap["listRow"] = listRow
	return paginatorMap
}

/**
	PagenationParse Temp
 */
func PagenationParse(page int, limitRow int) string {
	pageInfo := map[string]int{
		"starRow":  0,
		"limitRow": limitRow,
	}
	var _page = page
	if _page <= 0 {
		_page = 1
	}
	switch _page {
	case 1:
		return fmt.Sprintf("LIMIT %d,%d", pageInfo["starRow"], pageInfo["limitRow"])
	default:
		pageInfo["starRow"] = (page - 1) * limitRow
	}
	return fmt.Sprintf("LIMIT %d,%d", pageInfo["starRow"], pageInfo["limitRow"])
}

/**
	PagenationStart
 */
func PagenationStart(page int, limitRow int) int {
	var start = 0
	var _page = page
	if _page <= 0 {
		_page = 1
	}
	switch _page {
	case 1:
		start = 0
	default:
		start = (page - 1) * limitRow
	}
	return start
}


func QueryBuild(querySql string,currPage int,listRows int,usePage bool) string{
	var sql string
	var _listRows = ListRow
	if usePage == true{
		if listRows > 0{
			_listRows = listRows
		}
		sql = querySql + PagenationParse(currPage,_listRows)
	}
	return sql
}

func QueryTotalBuild(tableName string,otherCondi string,alias string) string {
	totalQuery := "SELECT count(" + _pkField + ") AS " + countField + " FROM " + tableName + " "
	if alias != ""{
		totalQuery += alias + " "
	}
	if otherCondi != ""{
		totalQuery += otherCondi
	}
	return totalQuery
}