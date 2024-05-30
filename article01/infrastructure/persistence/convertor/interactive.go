package convertor

import (
	"newsCenter/article01/domain/entity"
	"newsCenter/article01/infrastructure/persistence/database/interactive"
)

func ToInteractiveEntity(inter interactive.Interactive) entity.Interactive {
	return entity.Interactive{
		BizId:      inter.BizId,
		ReadCnt:    inter.ReadCnt,
		LikeCnt:    inter.LikeCnt,
		CollectCnt: inter.CollectCnt,
	}
}
