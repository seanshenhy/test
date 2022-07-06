package model

// Member 会员
type Member struct {
	ID       int    `json:"id" comment:"ID"`
	Name     string `json:"name" comment:"会员名称"`
	ParentId int    `json:"parent_id" comment:"父级ID"`
	LeafNode int    `json:"leaf_node" comment:"是否是叶子节点:1是,0不是"`
}

// Record 消费记录
type Record struct {
	ID        int    `json:"id" comment:"ID"`
	MemberId  int    `json:"member_id" comment:"会员ID"`
	Cost      int    `json:"cost" comment:"消费单位元"`
	CreatedAt int    `json:"created_at" comment:"创建时间"`
	MonDate   string `json:"mon_date" comment:"年月用来查询"`
}

// MonthStat 月统计
type MonthStat struct {
	ID        int    `json:"id" comment:"ID"`
	MemberId  int    `json:"member_id" comment:"会员ID"`
	Total     int    `json:"total" comment:"总消费"`
	CreatedAt int    `json:"created_at" comment:"创建时间"`
	MonDate   string `json:"mon_date" comment:"年月用来查询"`
}
