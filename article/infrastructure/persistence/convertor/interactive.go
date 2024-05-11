package convertor

import (
	"newsCenter/article/domain/entity"
	"newsCenter/article/infrastructure/persistence/database/interactive"
)

func ToInteractiveEntity(inter *interactive.Interactive) entity.Interactive {
	return entity.Interactive{
		ReadCnt:    inter.ReadCnt,
		LikeCnt:    inter.LikeCnt,
		CollectCnt: inter.CollectCnt,
	}
}
