package business

import (
	"time"
	"zupu/pkg/model"
)

func GetPerson() (model.Person, error) {

	return model.Person{
		Id:      1,
		Name:    "于江江",
		Age:     20,
		Address: "杭州",
		Phone:   "18814866344",
		Time: model.Time{
			CreateTime: time.Now().UnixMilli(),
			UpdateTime: time.Now().UnixMilli(),
		},
	}, nil
}
