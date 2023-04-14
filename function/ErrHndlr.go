package function

import (
	"net/http"
	"strings"
)

const (
	Sep             = ":$:"
	newLine         = "\n"
	Unauthorized    = "unauthorized"
	NotFoundError   = "not_found"
	DuplicatedError = "duplicated"
	ValidationError = "validation_error"
	unhandledError  = "unhandled_error"
)

// return httpCode, customCode, messages
func ErrHndlr(e error) (int, string, []string) {
	code := ""
	messages := []string{}
	httpStatus := http.StatusBadRequest

	text := strings.Split(e.Error(), Sep)

	// error text is separated by :$: value where:
	// http code
	// custom code
	// text code
	// other text
	if len(text) != 4 {
		panic("message error must be in four parts")
	}

	switch text[0] { // http code
	case Unauthorized:
		httpStatus = http.StatusUnauthorized
	case NotFoundError:
		httpStatus = http.StatusNotFound
	case DuplicatedError:
		httpStatus = http.StatusConflict
	case ValidationError:
		httpStatus = http.StatusBadRequest
	default:
		httpStatus = http.StatusBadRequest
	}

	code = text[1] // code

	if text[2] != "" { // text code
		messages = append(messages, text[2])
	}

	if text[3] != "" { // other text
		if strings.Contains(text[3], newLine) {
			messages = append(messages, strings.Split(text[3], newLine)...)
		} else {
			messages = append(messages, text[3])
		}
	}

	return httpStatus, code, messages
}
