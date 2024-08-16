package entity

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	DateTime string `json:"date"`
}
