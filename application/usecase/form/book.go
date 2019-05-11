package form

import "bookshelf-web-api/domain/model"

type BookRequest struct {
	Title string
	AuthorId int64
	CategoryIds []int64
	PrevBookId int64 // default 0
	NextBookId int64 // default 0
}

type DescriptionRequest struct {
	Description string
}

type BookStatusResponse struct {
	BookId int64
	ReadStatus model.ReadState
}