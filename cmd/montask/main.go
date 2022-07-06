package main

import (
	"test/internal/model"
	"time"
)

/**
 * 作用：月统计需要定时任务执行或手动执行一次
**/
func main() {
	// 查询parent_id=0的会员节点
	var mems []model.Member
	model.Db.Raw("SELECT id,parent_id FROM member WHERE parent_id = ?", 0).Scan(&mems)
	month := time.Now().Format("2006-01")
	StaticChildMemberCost(mems, month)
}

// StaticChildMemberCost 统计member每月消费
func StaticChildMemberCost(mems []model.Member, month string) int {
	if len(mems) == 0 {
		return 0
	}
	totalCost := 0
	for i := 0; i < len(mems); i++ {
		var cmems []model.Member
		model.Db.Raw("SELECT id FROM member WHERE parent_id = ?", mems[i].ID).Scan(&cmems)
		// 孩子节点消费额度
		childCost := StaticChildMemberCost(cmems, month)
		totalCost += CaculateMemberMonthStat(month, mems[i].ID, childCost)
	}
	return totalCost
}

// CaculateMemberMonthStat 计算会员月统计
func CaculateMemberMonthStat(month string, member_id int, childCost int) int {
	var r model.Record
	// 统计会员月消耗
	model.Db.Raw("SELECT member_id,sum(cost) as cost FROM record WHERE mon_date = ? and member_id=? GROUP by mon_date,member_id", month, member_id).Scan(&r)
	totalCost := childCost + r.Cost
	if totalCost > 0 {
		// 插入记录
		model.Db.Exec("INSERT INTO month_stat (member_id,total,created_at,mon_date) VALUES (?,?,?,?)", member_id, totalCost, int(time.Now().Unix()), month)
	}
	return totalCost
}
