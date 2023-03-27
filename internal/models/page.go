package models

type Pagination struct {
	Page int `query:"page" validate:"min=0"`
	Size int `query:"size" validate:"min=10"`
}
