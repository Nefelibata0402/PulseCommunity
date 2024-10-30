package article

import (
	"github.com/gin-gonic/gin"
	"pulseCommunity/cmd/middlewares/tokenVerify"
)

func InitArticleRouter(r *gin.Engine) {
	router := r.Group("newsCenter/article")
	router.Use(tokenVerify.TokenVerify())
	router.POST("/edit", Edit)         // 编辑
	router.POST("/publish", Publish)   //发布
	router.POST("/withdraw", Withdraw) //撤回 改变状态仅自己可见
	//作者自己的接口
	router.GET("/detail/:id", Detail) //作者查看详情
	router.POST("/list", List)        //作者的列表接口
	router.GET("/read/:id", Read)     //读者阅读文章
	router.POST("/like", Like)        //点赞文章
	router.POST("/collect", Collect)  //收藏文章
}
