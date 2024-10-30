package convertor

import (
	"pulseCommunity/article/domain/entity"
	"pulseCommunity/article/infrastructure/persistence/database/interactive"
)

func ToInteractiveEntity(inter interactive.Interactive) entity.Interactive {
	return entity.Interactive{
		BizId:      inter.BizId,
		ReadCnt:    inter.ReadCnt,
		LikeCnt:    inter.LikeCnt,
		CollectCnt: inter.CollectCnt,
	}
}
