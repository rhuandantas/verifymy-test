package errors

import (
	"github.com/joomcode/errorx"
)

var (
	NotFound     = errorx.CommonErrors.NewType("not_found", errorx.NotFound())
	BadRequest   = errorx.CommonErrors.NewType("bad_request")
	Unauthorized = errorx.CommonErrors.NewType("unauthorized")
)
