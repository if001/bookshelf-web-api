package form

type BookRequest struct {
	Title string
	Author string
	Categories []string // default null
	PrevBookId int64 // default 0
	NextBookId int64 // default 0
}

type DescriptionRequest struct {
	Description string
}
