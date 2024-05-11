package collection

type Collection struct {
	Id    int64  `gorm:"uniqueIndex:uid_biz_type_id"`
	BizId int64  `gorm:"uniqueIndex:uid_biz_type_id"`
	Biz   string `gorm:"type:varchar(128);uniqueIndex:uid_biz_type_id"`
	// 收藏夹的ID
	// 收藏夹ID本身有索引
	Cid       int64 `gorm:"index"`
	CreatedAt int64
	UpdatedAt int64
}

func (*Collection) TableName() string {
	return "collection"
}
