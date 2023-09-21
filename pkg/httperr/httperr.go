package httperr

import (
	"hangout/pkg/richerror"
	"net/http"
)

func Error(err error) (int, string) {
	switch err.(type) {
	case richerror.RichError:
		rErr := err.(richerror.RichError)
		code := mapKindToStatusCode(rErr.Kind())
		msg := rErr.Message()

		if code >= 500 {
			msg = "something went wrong"
		}

		return code, msg
	default:
		return http.StatusBadRequest, err.Error()
	}
}

func mapKindToStatusCode(k richerror.Kind) int {
	switch k {
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
		return http.StatusBadRequest
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	case richerror.KindNotFound:
		return http.StatusNotFound
	default:
		return http.StatusBadRequest
	}
}
