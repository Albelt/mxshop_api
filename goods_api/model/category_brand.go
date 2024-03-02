package model

import pb "albelt.top/mxshop_protos/albelt/good_srv/go"

type CategoryBrand struct {
	Id         int32
	BrandId    int32
	Brand      *Brand
	CategoryId int32
	Category   *Category
}

func NewCategoryBrand(cb *pb.CategoryBrand) *CategoryBrand {
	return &CategoryBrand{
		Id:         cb.Id,
		BrandId:    cb.BrandId,
		Brand:      NewBrandFromPb(cb.Brand),
		CategoryId: cb.CategoryId,
		Category:   NewCategoryFromPb(cb.Category),
	}
}
