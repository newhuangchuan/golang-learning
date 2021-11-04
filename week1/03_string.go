package main

import "fmt"

//反引号
//- 不支持转义
//- 支持换行
//- 主要用来创建原生的字符串
//- 举例
func main() {

	jsonstr := `
	{
	"region":"bj",
	"ids":[1,2,3,4]
	}
`
	promql := `sum(rate(api_qps{code=~"2xx"}{1m}))*100`
	fmt.Println(jsonstr,promql)

}
