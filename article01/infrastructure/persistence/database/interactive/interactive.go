package interactive

type Interactive struct {
	Id         int64  `json:"id" gorm:"primaryKey,autoIncrement"`
	BizId      int64  `json:"biz_id" gorm:"uniqueIndex:biz_type_id"`
	Biz        string `json:"biz" gorm:"type:varchar(128);uniqueIndex:biz_type_id"`
	ReadCnt    int64  `json:"read_cnt"`
	LikeCnt    int64  `json:"like_cnt"`
	CollectCnt int64  `json:"collect_cnt"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func (*Interactive) TableName() string {
	return "interactive"
}
