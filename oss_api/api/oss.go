package api

import (
	"fmt"
	"mxshop_api/oss_api/global"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"mxshop_api/oss_api/utils"
)

func Token(c *gin.Context) {
	response := utils.Get_policy_token(global.Config.Oss)
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Origin", "*")
	c.String(200, response)
}

func HandlerRequest(ctx *gin.Context) {
	// Get PublicKey bytes
	bytePublicKey, err := utils.GetPublicKey(ctx)
	if err != nil {
		utils.ResponseFailed(ctx)
		return
	}

	// Get Authorization bytes : decode from Base64String
	byteAuthorization, err := utils.GetAuthorization(ctx)
	if err != nil {
		utils.ResponseFailed(ctx)
		return
	}

	// Get MD5 bytes from Newly Constructed Authrization String.
	byteMD5, bodyStr, err := utils.GetMD5FromNewAuthString(ctx)
	if err != nil {
		utils.ResponseFailed(ctx)
		return
	}

	decodeurl, err := url.QueryUnescape(bodyStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodeurl)
	params := make(map[string]string)
	datas := strings.Split(decodeurl, "&")
	for _, v := range datas {
		sdatas := strings.Split(v, "=")
		params[sdatas[0]] = sdatas[1]
	}
	fileName := params["filename"]
	fileUrl := fmt.Sprintf("%s/%s", global.Config.Oss.Host, fileName)

	// verifySignature and response to client
	if utils.VerifySignature(bytePublicKey, byteMD5, byteAuthorization) {
		// do something you want accoding to callback_body ...
		ctx.JSON(http.StatusOK, gin.H{
			"url": fileUrl,
		})
		//utils.ResponseSuccess(ctx)  // response OK : 200
	} else {
		utils.ResponseFailed(ctx) // response FAILED : 400
	}
}
