package convertor

import (
	"newsCenter/idl/articleGrpc"
	"newsCenter/ranking/domain/entity"
	"time"
)

func InteractiveProtoToEntity(proto *articleGrpc.Interactive) entity.Interactive {
	return entity.Interactive{
		BizId:      proto.BizId,
		ReadCnt:    proto.ReadCnt,
		LikeCnt:    proto.LikeCnt,
		CollectCnt: proto.CollectCnt,
		Liked:      proto.Liked,
		Collected:  proto.Collected,
	}
}

func ArticleProtoToEntity(proto *articleGrpc.Article) entity.Article {
	createdAt, _ := time.Parse("2006-01-02 15:04:05", proto.CreateAt)
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", proto.UpdateAt)
	return entity.Article{
		Id:        uint64(proto.Id),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Author: entity.Author{
			Id:   uint64(proto.Author.Id),
			Name: proto.Author.Name,
		},
		Content:  proto.Content,
		Category: proto.Category,
		Title:    proto.Title,
	}
}
