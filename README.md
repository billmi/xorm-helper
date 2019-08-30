BaseFrom : go-xorm engine

## **Install**
`
go get -u -v github.com/billmi/xorm-pagenation
`

### 说明 : 

| 说明                                            | 链接                                                         | 备注                              |
| ----------------------------------------------- | ------------------------------------------------------------ | --------------------------------- |
| join生成器                                      | [Join](https://github.com/billmi/xorm-pagenation/blob/master/example/join.go) | 您也可以使用xorm提供的Builder生成 |
| 实现xorm常规操作DB,使用map结构化,使用时参考demo | [Curd](https://github.com/billmi/xorm-pagenation/blob/master/example/curd.go) | 您也可以使用xorm提供的Builder生成 |
| condition生成器                                 | [Condition](https://github.com/billmi/xorm-pagenation/blob/master/example/condition-build.go) | 使用时,请查看Demo                 |
| 连表分页查询                                    | [Pagenation](https://github.com/billmi/xorm-pagenation/blob/master/example/pagenation-list.go) | 结果必须关联Pojo-Struct           |
| 不分页List                                      | [List-NoPagenation](https://github.com/billmi/xorm-pagenation/blob/master/example/list.go) | 结果必须关联Pojo-Struct           |

ps : test.sql用来测试

此项目初衷，不想把一些DB简单操作复杂化,做为一个扩展工具使用.