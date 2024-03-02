package api

import (
	pb "albelt.top/mxshop_protos/albelt/good_srv/go"
	"github.com/gin-gonic/gin"
	"mxshop_api/goods_api/global"
	"mxshop_api/goods_api/model"
	"net/http"
)

func ListGoods(ctx *gin.Context) {
	var (
		err  error
		req  model.ListGoodsRequest
		resp model.ListGoodsResponse
	)

	if err = ctx.ShouldBindQuery(&req); err != nil {
		handleParamErr(ctx, err)
		return
	}

	pbResp, err := global.GoodsSrvCli.GoodsList(ctx, &pb.GoodsListRequest{
		PriceMin:   req.PriceMin,
		PriceMax:   req.PriceMax,
		IsHot:      req.IsHot,
		IsNew:      req.IsNew,
		IsTab:      false,
		CategoryId: req.CategoryId,
		Keywords:   req.Keywords,
		Brand:      req.Brand,
		Page:       req.Page,
		Count:      req.Count,
	})
	if err != nil {
		handleGrpcErr(ctx, err)
		return
	}

	resp.Total = pbResp.Total
	for _, g := range pbResp.Goods {
		resp.Goods = append(resp.Goods, model.NewGoodsFromPb(g))
	}

	ctx.JSON(http.StatusOK, resp)
}

func GetGoods(ctx *gin.Context) {

}

func CreateGoods(ctx *gin.Context) {

}

func DeleteGoods(ctx *gin.Context) {

}

func UpdateGoods(ctx *gin.Context) {

}
