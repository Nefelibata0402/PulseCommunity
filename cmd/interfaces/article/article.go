package article

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pulseCommunity/cmd/model/articleModel"
	"pulseCommunity/common/unierr"
	"pulseCommunity/idl/articleGrpc"
	"strconv"
	"time"
)

func Edit(c *gin.Context) {
	var articleReq articleModel.ArticleRequest
	//1.绑定参数
	if err := c.ShouldBind(&articleReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorParams.ErrCode,
			"status_msg":  unierr.ErrorParams.ErrMsg,
		})
		return
	}
	//2.校验参数
	if err := articleModel.ValidateArticleRequest(&articleReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ArticleTitleOrContentNotNil.ErrCode,
			"status_msg":  unierr.ArticleTitleOrContentNotNil.ErrMsg,
		})
		return
	}
	ctx, cancel := context.WithTimeout(c, 180*time.Second)
	defer cancel()
	userId, _ := c.Get("userId")
	req := &articleGrpc.EditRequest{
		ArticleId: articleReq.ArticleId,
		AuthorId:  userId.(uint64),
		Title:     articleReq.Title,
		Data:      articleReq.Content,
		Category:  articleReq.Category,
	}
	resp, err := ArticleServiceClient.Edit(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorInternal.ErrCode,
			"status_msg":  unierr.ErrorInternal.ErrMsg,
		})
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
	})
}

func Publish(c *gin.Context) {
	var articleReq articleModel.ArticleRequest
	//1.绑定参数
	if err := c.ShouldBind(&articleReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorParams.ErrCode,
			"status_msg":  unierr.ErrorParams.ErrMsg,
		})
		return
	}
	//2.校验参数
	if err := articleModel.ValidateArticleRequest(&articleReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ArticleTitleOrContentNotNil.ErrCode,
			"status_msg":  unierr.ArticleTitleOrContentNotNil.ErrMsg,
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	userId, _ := c.Get("userId")
	req := &articleGrpc.PublishRequest{
		ArticleId: articleReq.ArticleId,
		Title:     articleReq.Title,
		AuthorId:  userId.(uint64),
		Data:      articleReq.Content,
		Category:  articleReq.Category,
	}
	resp, err := ArticleServiceClient.Publish(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorInternal.ErrCode,
			"status_msg":  unierr.ErrorInternal.ErrMsg,
		})
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
	})
}

func Withdraw(c *gin.Context) {
	var withdrawReq articleModel.ArticleWithdrawRequest
	if err := c.ShouldBind(&withdrawReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorParams.ErrCode,
			"status_msg":  unierr.ErrorParams.ErrMsg,
		})
		return
	}
	if err := articleModel.ValidateArticleWithdrawRequest(&withdrawReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.WithdrawArticleIDNotNIl.ErrCode,
			"status_msg":  unierr.WithdrawArticleIDNotNIl.ErrMsg,
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	userId, _ := c.Get("userId")
	req := &articleGrpc.WithdrawRequest{
		ArticleId: withdrawReq.ArticleId,
		AuthorId:  userId.(uint64),
	}
	resp, err := ArticleServiceClient.WithdrawArticle(ctx, req)
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

func Detail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorParams.ErrCode,
			"status_msg":  unierr.ErrorParams.ErrMsg,
		})
		return
	}
	userId, _ := c.Get("userId")
	req := &articleGrpc.GetDetailRequest{
		ArticleId: uint64(id),
		AuthorId:  userId.(uint64),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	resp, err := ArticleServiceClient.GetDetail(ctx, req)
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
		"data":        resp.Article,
	})
}

func List(c *gin.Context) {
	var page articleModel.Page
	if err := c.ShouldBind(&page); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorParams.ErrCode,
			"status_msg":  unierr.ErrorParams.ErrMsg,
		})
		return
	}
	userId, _ := c.Get("userId")
	req := &articleGrpc.GetListByAuthorRequest{
		AuthorId: userId.(uint64),
		Offset:   page.Offset,
		Limit:    page.Limit,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	resp, err := ArticleServiceClient.GetList(ctx, req)
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

func Read(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorParams.ErrCode,
			"status_msg":  unierr.ErrorParams.ErrMsg,
		})
		return
	}
	userId, _ := c.Get("userId")
	req := &articleGrpc.ReadRequest{
		UserId:    userId.(uint64),
		ArticleId: uint64(id),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	resp, err := ArticleServiceClient.Read(ctx, req)
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
		"data":        resp.Article,
	})
}

func Like(c *gin.Context) {
	var Req articleModel.Like
	if err := c.ShouldBind(&Req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorParams.ErrCode,
			"status_msg":  unierr.ErrorParams.ErrMsg,
		})
		return
	}
	if err := articleModel.ValidateLikeRequest(&Req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.WithdrawArticleIDNotNIl.ErrCode,
			"status_msg":  unierr.WithdrawArticleIDNotNIl.ErrMsg,
		})
		return
	}
	userId, _ := c.Get("userId")
	req := &articleGrpc.LikeRequest{
		Id:     Req.Id,
		Like:   Req.Like,
		UserId: userId.(uint64),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	resp, err := ArticleServiceClient.Like(ctx, req)
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

// Collect 和点赞差不多先不写
func Collect(c *gin.Context) {

}
