package ranking

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"newsCenter/common/unierr"
	"newsCenter/idl/rankingGrpc"
	"time"
)

func GetTopN(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 180*time.Second)
	defer cancel()
	req := &rankingGrpc.GetTopNRequest{}
	resp, err := RankingServiceClient.GetTopN(ctx, req)
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
		"data":        resp.ArticleList,
	})
}

func TopN(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 180*time.Second)
	defer cancel()
	req := &rankingGrpc.TopNRequest{}
	resp, err := RankingServiceClient.TopN(ctx, req)
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
	})
}
