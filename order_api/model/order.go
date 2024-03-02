package model

import (
	pb "albelt.top/mxshop_protos/albelt/order_srv/go"
)

type CreateOrderRequest struct {
	UserId      int32      `json:"user_id" binding:"required"`
	UserName    string     `json:"user_name" binding:"required"`
	UserMobile  string     `json:"user_mobile" binding:"required"`
	UserAddress string     `json:"user_address" binding:"required"`
	PayType     pb.PayType `json:"pay_type" binding:"required"`
}

type CreateOrderResponse struct {
	OrderInfo *Order
}

type Order struct {
	Id      int32          `json:"id,omitempty"`
	No      string         `json:"no,omitempty"`
	PayType pb.PayType     `json:"pay_type,omitempty"`
	Status  pb.OrderStatus `json:"status,omitempty"`
	Amount  float32        `json:"amount,omitempty"`
	Goodss  []*OrderGoods  `json:"goodss"`
}

type OrderGoods struct {
	Id         int32   `json:"id,omitempty"`
	OrderId    int32   `json:"order_id,omitempty"`
	GoodsId    int32   `json:"goods_id,omitempty"`
	GoodsNum   int32   `json:"goods_num,omitempty"`
	GoodsName  string  `json:"goods_name,omitempty"`
	GoodsPrice float32 `json:"goods_price,omitempty"`
}

func NewOrderFromPb(o *pb.Order) *Order {
	if o == nil {
		return nil
	}

	order := &Order{
		Id:      o.Id,
		No:      o.No,
		PayType: o.PayType,
		Status:  o.Status,
		Amount:  o.Amount,
	}

	for _, g := range o.Goodss {
		order.Goodss = append(order.Goodss, NewOrderGoodsFromPb(g))
	}

	return order
}

func NewOrderGoodsFromPb(g *pb.OrderGoods) *OrderGoods {
	if g == nil {
		return nil
	}

	return &OrderGoods{
		Id:         g.Id,
		OrderId:    g.OrderId,
		GoodsId:    g.GoodsId,
		GoodsNum:   g.GoodsNum,
		GoodsName:  g.GoodsName,
		GoodsPrice: g.GoodsPrice,
	}
}
