package search

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pulseCommunity/cmd/model/searchModel"
	"pulseCommunity/common/unierr"
	"pulseCommunity/idl/searchGrpc"
	"time"
)

func Search(c *gin.Context) {
	var searchReq searchModel.SearchRequest
	//1.绑定参数
	if err := c.ShouldBind(&searchReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorParams.ErrCode,
			"status_msg":  unierr.ErrorParams.ErrMsg,
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, 180*time.Second)
	defer cancel()
	req := &searchGrpc.SearchRequest{
		Expression: searchReq.Expression,
	}
	resp, err := SearchServiceClient.Search(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorInternal.ErrCode,
			"status_msg":  unierr.ErrorInternal.ErrMsg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"data1":       resp.Article,
		"data2":       resp.User,
	})
}
