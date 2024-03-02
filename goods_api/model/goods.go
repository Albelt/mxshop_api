package model

import (
	pb "albelt.top/mxshop_protos/albelt/good_srv/go"
	"mxshop_api/goods_api/utils"
	"time"
)

type Goods struct {
	Id              int32      `json:"id"`
	Name            string     `json:"name"`
	GoodsSn         string     `json:"goods_sn"`
	ClickNum        int32      `json:"click_num"`
	SoldNum         int32      `json:"sold_num"`
	FavNum          int32      `json:"fav_num"`
	MarketPrice     float32    `json:"market_price"`
	ShopPrice       float32    `json:"shop_price"`
	GoodsBrief      string     `json:"goods_brief"`
	ShipFree        bool       `json:"ship_free"`
	GoodsFrontImage string     `json:"goods_front_image"`
	IsNew           bool       `json:"is_new"`
	IsHot           bool       `json:"is_hot"`
	OnSale          bool       `json:"on_sale"`
	CategoryId      int32      `json:"category_id"`
	Category        *Category  `json:"category"`
	BrandId         int32      `json:"brand_id"`
	Brand           *Brand     `json:"brand"`
	CreateTime      *time.Time `json:"create_time"`
}

func NewGoodsFromPb(g *pb.Good) *Goods {
	if g == nil {
		return nil
	}

	goods := &Goods{
		Id:              g.Id,
		Name:            g.Name,
		GoodsSn:         g.GoodsSn,
		ClickNum:        g.ClickNum,
		SoldNum:         g.SoldNum,
		FavNum:          g.FavNum,
		MarketPrice:     g.MarketPrice,
		ShopPrice:       g.ShopPrice,
		GoodsBrief:      g.GoodsBrief,
		ShipFree:        g.ShipFree,
		GoodsFrontImage: g.GoodsFrontImage,
		IsNew:           g.IsNew,
		IsHot:           g.IsHot,
		OnSale:          g.OnSale,
		CategoryId:      g.CategoryId,
		Category:        NewCategoryFromPb(g.Category),
		BrandId:         g.BrandId,
		Brand:           NewBrandFromPb(g.Brand),
		CreateTime:      utils.Val2Ptr(g.CreateTime.AsTime()),
	}

	return goods
}

type ListGoodsRequest struct {
	PriceMin   int32  `form:"price_min" binding:""`
	PriceMax   int32  `form:"price_max" binding:""`
	IsHot      bool   `form:"is_hot" binding:""`
	IsNew      bool   `form:"is_new" binding:""`
	IsTab      bool   `form:"is_tab" binding:""`
	CategoryId int32  `form:"category_id" binding:""`
	Keywords   string `form:"keywords" binding:""`
	Brand      int32  `form:"brand" binding:""`
	Page       int32  `form:"page" binding:"required"`
	Count      int32  `form:"count" binding:"required"`
}

type ListGoodsResponse struct {
	Total int32    `json:"total"`
	Goods []*Goods `json:"goods"`
}
