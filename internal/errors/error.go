package errors

import (
	"fmt"
	"github.com/joomcode/errorx"
)

var (
	NotFound     = errorx.CommonErrors.NewType("not_found", errorx.NotFound())
	BadRequest   = errorx.CommonErrors.NewType("bad_request")
	Unauthorized = errorx.CommonErrors.NewType("unauthorized")
)

func InvalidRecipeAlternative(recipeType string) *errorx.Error {
	return errorx.InternalError.New(fmt.Sprintf("Unsupported recipe alternative: %s", recipeType))
}

func ServiceCardstackClientProxyError(cause error) *errorx.Error {
	return errorx.InternalError.Wrap(cause, "Couldn't proxy to Service Cardstack Client")
}

func PageResponseError(cause error, id string) *errorx.Error {
	message := fmt.Sprintf("Couldn't retrieve page with ID: %s", id)

	return errorx.ExternalError.Wrap(cause, message)
}

func SectionHeaderInvalidFilterOperationType() *errorx.Error {
	return errorx.IllegalArgument.New("Invalid filter operation type for section header")
}
