package form

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type BookRequest struct {
	Title       string
	AuthorId    int64
	CategoryIds []int64
	PrevBookId  int64 // default 0
	NextBookId  int64 // default 0
}

type DescriptionRequest struct {
	Description string
}

type BookStatusResponse struct {
	BookId     int64
	ReadStatus model.ReadState
}

type BookResponse struct {
	Base
	Name         string                `json:"name"`
	Author       *AuthorResponse       `json:"author"`
	PublishedAt  service.NullTime      `json:"publish_at"`
	Publisher    *PublisherResponse    `json:"publisher"`
	AccountId    int64                 `json:"account_idn"`
	StartAt      service.NullTime      `json:"start_at"`
	EndAt        service.NullTime      `json:"end_at"`
	NextBookID   service.NullInt64     `json:"next_book_id"`
	PrevBookID   service.NullInt64     `json:"prev_book_id"`
	Descriptions []DescriptionResponse `json:"descriptions"`
	Categories   []CategoryResponse    `json:"categories"`
}

type AuthorResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type DescriptionResponse struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}
type CategoryResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type PublisherResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
