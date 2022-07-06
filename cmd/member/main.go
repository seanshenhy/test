package main

import (
	"flag"
	"fmt"
	"test/internal/model"
)

/**
 * 作用：查询某会员，在某月的消费总额
**/
func main() {
	// 获取解析参数会员id和月信息
	var (
		member_id int
		month     string
	)
	flag.IntVar(&member_id, "member_id", 1, "")
	flag.StringVar(&month, "month", "2022-07", "")
	flag.Parse()
	// 查询月统计表某会员某记录
	ms := GetMonthStaticRecord(member_id, month)
	m := GetMemberInfo(member_id)
	fmt.Printf("会员（%s） 在%s 消费总额为 %d元\n", m.Name, month, ms.Total)
}

// GetMonthStaticRecord 获取月统计记录表
func GetMonthStaticRecord(member_id int, month string) model.MonthStat {
	var ms model.MonthStat
	model.Db.Raw("SELECT member_id,total,mon_date FROM month_stat WHERE member_id = ? and mon_date = ?", member_id, month).Scan(&ms)
	return ms
}

// GetMemberInfo 获取会员信息
func GetMemberInfo(member_id int) model.Member {
	var m model.Member
	model.Db.Raw("SELECT id,name FROM member WHERE id = ? ", member_id).Scan(&m)
	return m
}
