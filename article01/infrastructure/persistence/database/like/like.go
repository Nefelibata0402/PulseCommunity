package like

type Like struct {
	Id        int64  `gorm:"uniqueIndex:uid_biz_type_id"`
	BizId     int64  `gorm:"uniqueIndex:uid_biz_type_id"`
	Biz       string `gorm:"type:varchar(128);uniqueIndex:uid_biz_type_id"`
	Status    int
	CreatedAt int64
	UpdatedAt int64
}

func (*Like) TableName() string {
	return "like"
}
