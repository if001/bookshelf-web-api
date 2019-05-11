package form

import (
	"bookshelf-web-api/domain/model"
	"database/sql"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
)

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

var nullLiteral = []byte("null")
type NullInt64 struct {
	sql.NullInt64
}
func (i NullInt64) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int64)
	} else {
		return json.Marshal(nil)
	}
}
func (i *NullInt64) UnmarshalJSON(data []byte) error {
	var num int64
	if err := json.Unmarshal(data, &num); err != nil {
		return err
	}
	i.Int64 = num
	i.Valid = num != 0
	return nil
}

type NullTime struct {
	mysql.NullTime
}
func (t *NullTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	} else {
		return nullLiteral, nil
	}
}
