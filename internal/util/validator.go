package util

import "github.com/go-playground/validator/v10"

//go:generate mockgen -source=$GOFILE -package=mock_util -destination=../../test/mock/util/$GOFILE

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (cv *CustomValidator) ValidateStruct(i interface{}) error {
	return cv.validator.Struct(i)
}
