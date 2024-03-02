package api

import (
	pb "albelt.top/mxshop_protos/albelt/order_srv/go"
	"github.com/gin-gonic/gin"
	"mxshop_api/order_api/global"
	"mxshop_api/order_api/model"
	"net/http"
)

func CreateOrder(ctx *gin.Context) {
	var (
		err  error
		req  model.CreateOrderRequest
		resp model.CreateOrderResponse
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		handleParamErr(ctx, err)
		return
	}

	pbResp, err := global.OrderSrvCli.CreateOrderFromCart(ctx, &pb.CreateOrderFromCartRequest{
		UserId: req.UserId,
		Order: &pb.Order{
			UserId:      req.UserId,
			UserName:    req.UserName,
			UserMobile:  req.UserMobile,
			UserAddress: req.UserAddress,
		},
	})
	if err != nil {
		handleGrpcErr(ctx, err)
		return
	}

	resp.OrderInfo = model.NewOrderFromPb(pbResp.Order)

	ctx.JSON(http.StatusOK, resp)
}

func GetOrder(ctx *gin.Context) {

}
