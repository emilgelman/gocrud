package domain

type Article struct {
	Id string `json:"Id" validate:"required"`
	Title string `json:"Title" validate:"required"`
	Content string `json:"content"`
}
