package model

import pb "albelt.top/mxshop_protos/albelt/good_srv/go"

type Category struct {
	Id               int32
	Name             string
	ParentCategoryId int32
	Level            int32
	IsTab            bool
}

func NewCategoryFromPb(c *pb.Category) *Category {
	return &Category{
		Id:               c.Id,
		Name:             c.Name,
		ParentCategoryId: c.ParentCategoryId,
		Level:            c.Level,
		IsTab:            c.IsTab,
	}
}
