package article

type Article struct {
	Id        uint64 `json:"id" bson:"id,omitempty"`
	CreatedAt int64  `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt int64  `json:"deleted_at" bson:"deleted_at,omitempty"`
	UserId    uint64 `json:"user_id" bson:"user_id,omitempty"`
	Content   string `json:"content" bson:"content,omitempty"`
	Category  string `json:"category" bson:"category"`
	Title     string `json:"title" bson:"title"`
	Status    uint8  `bson:"status,omitempty"`
}

func (*Article) TableName() string {
	return "article"
}
