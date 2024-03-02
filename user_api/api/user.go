package api

import (
	pb "albelt.top/mxshop_protos/albelt/user_srv/go"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop_api/user_api/global"
	"mxshop_api/user_api/middleware"
	"mxshop_api/user_api/model"
	"net/http"
	"strconv"
)

func ListUsers(ctx *gin.Context) {
	var (
		err  error
		req  model.ListUsersRequest
		resp model.ListUsersResponse
	)

	if err = ctx.ShouldBindQuery(&req); err != nil {
		handleParamErr(ctx, err)
		return
	}

	tmp, err := global.UserSrvCli.GetUserList(ctx, &pb.GetUserListRequest{
		Page: uint32(req.Page),
		Size: uint32(req.Count),
	})
	if err != nil {
		handleGrpcErr(ctx, err)
		return
	}

	resp.Total = tmp.Total
	for _, u := range tmp.Users {
		resp.Users = append(resp.Users, model.NewUserFromPb(u))
	}

	ctx.JSON(http.StatusOK, resp)
}

func GetUser(ctx *gin.Context) {
	var (
		err    error
		userId int64
		user   *model.User
		pbUser *pb.User
	)

	userId, err = strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		handleParamErr(ctx, err)
		return
	}

	pbUser, err = global.UserSrvCli.GetUserById(ctx, &pb.GetUserByIdRequest{Id: int32(userId)})
	if err != nil {
		handleGrpcErr(ctx, err)
		return
	}

	user = model.NewUserFromPb(pbUser)

	ctx.JSON(http.StatusOK, user)
}

func LoginUser(ctx *gin.Context) {
	var (
		err error
		req model.LoginUserRequest

		pbUser *pb.User
		user   *model.User
		token  string
	)

	// 校验参数
	if err = ctx.ShouldBindJSON(&req); err != nil {
		handleParamErr(ctx, err)
		return
	}

	// 校验验证码
	if req.CaptchaId == "000000" && req.CaptchaAnswer == "000000" {
		// 万能验证码
	} else {
		if !global.RedisStore.Verify(req.CaptchaId, req.CaptchaAnswer, true) {
			handleParamErr(ctx, errors.New("验证码错误"))
			return
		}
	}

	// 查询用户
	pbUser, err = global.UserSrvCli.GetUserByMobile(ctx, &pb.GetUserByMobileRequest{Mobile: req.Mobile})
	if err != nil {
		handleGrpcErr(ctx, err)
		return
	}

	// 校验用户密码
	check, err := global.UserSrvCli.CheckPassWord(ctx, &pb.CheckPassWordRequest{
		Password:          req.Password,
		EncryptedPassword: pbUser.Password,
	})
	if err != nil {
		handleGrpcErr(ctx, err)
		return
	}

	if !check.Success {
		handleParamErr(ctx, errors.New("密码错误"))
		return
	}

	user = model.NewUserFromPb(pbUser)

	// 生成jwt-token
	token, err = middleware.GlobalJwt.GenToken(user)
	if err != nil {
		handleInternalErr(ctx, err, "生成token失败")
		return
	}

	resp := model.LoginUserResponse{
		Id:       user.Id,
		NickName: user.NickName,
		Token:    token,
	}

	ctx.JSON(http.StatusOK, resp)
}

func RegisterUser(ctx *gin.Context) {
	var (
		err  error
		req  model.RegisterUserRequest
		user *model.User
	)

	// 校验参数
	if err = ctx.ShouldBindJSON(&req); err != nil {
		handleParamErr(ctx, err)
		return
	}

	// 校验短信验证码
	if ok, err := global.Sms.VerifyCode(req.Mobile, req.Code); err != nil {
		handleInternalErr(ctx, err, "校验短信验证码错误")
		return
	} else if !ok {
		handleParamErr(ctx, errors.New("短信验证码错误"))
		return
	}

	// 创建用户
	pbUser, err := global.UserSrvCli.CreateUser(ctx, &pb.CreateUserRequest{
		NickName: req.NickName,
		Password: req.Password,
		Mobile:   req.Mobile,
	})
	if err != nil {
		handleGrpcErr(ctx, err)
		return
	}

	user = model.NewUserFromPb(pbUser)
	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context) {

}

func handleParamErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}

func handleGrpcErr(ctx *gin.Context, err error) {
	// 将grpc code转换成http code
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.AlreadyExists:
			ctx.JSON(http.StatusConflict, gin.H{
				"message": fmt.Sprintf("数据已存在:%s", s.Message()),
			})
		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("数据不存在:%s", s.Message()),
			})
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("参数错误:%s", s.Message()),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    s.Code().String(),
				"message": s.Message(),
			})

		}
	}
}

func handleInternalErr(ctx *gin.Context, err error, note string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": fmt.Sprintf("%s:%s", note, err.Error()),
	})
}
