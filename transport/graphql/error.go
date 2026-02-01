package graphql

import (
	"go/helperpkg/errors"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Error(err *errors.AppError) *gqlerror.Error {
	return &gqlerror.Error{
		Message: err.Message,
		Extensions: map[string]interface{}{
			"code": err.Code,
			"type": err.Type,
		},
	}
}
