package entity

type Article struct {
	Id      int64
	Title   string
	Status  int32
	Content string
}

type User struct {
	Id       int64
	Nickname string
}

type SearchResult struct {
	Users    []User
	Articles []Article
}
