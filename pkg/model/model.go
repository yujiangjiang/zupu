package model

type Time struct {
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	PersonId int64  `json:"personId"`
	Time     `json:",inline"`
}

type Person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	IDCard   string `json:"idCard"`
	Age      int    `json:"age"`
	Birth    string `json:"birth"`
	Death    string `json:"death"`
	Sex      int    `json:"sex"`
	Phone    string `json:"phone"`
	FatherId int64  `json:"fatherId"`
	MotherId int64  `json:"motherId"`
	// 1：未婚、2：已婚、3：离异
	Marital int `json:"marital"`
	// 家庭住址
	Address string `json:"address"`
	//学历 小学|初中|高中|本科|硕士|博士
	Education string `json:"education"`
	// 职业
	Profession string `json:"profession"`
	// 介绍
	Introduction string `json:"introduction"`
	PhotoId      int64  `json:"photoId"`
	Time         `json:",inline"`
}

type Attachment struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Type    string `json:"type"`
	Time    `json:",inline"`
}
