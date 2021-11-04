package main

import (
	"fmt"
	"time"
)

func main() {

	want := `
[报警触发类型 :%s]
[报警名称 :%s]
[级别 :%d]
[机器ip列表 :%s]
[表达式 :%s]
[告警次数 :%d]
[触发时间：%s]
	`

	alarmContent := fmt.Sprintf(
		want,"promtheus",
		"api qps more then 1000",
		1,
		"10.10.10.20,20.10.11.4",
		`sum(rate(login_qps[1ms]))>100`,
		2,
		time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05"))

	fmt.Println(alarmContent)

}
