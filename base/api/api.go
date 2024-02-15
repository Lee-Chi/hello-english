package api

import (
	"fmt"
)

type ResponseBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Code int

func (c Code) Dump(format string, args ...interface{}) string {
	if format == "" {
		return fmt.Sprintf("code: %d, message: %s", c, messages[c])
	}

	return fmt.Sprintf("code: %d, message: %s, error: %s", c, messages[c], fmt.Sprintf(format, args...))
}
func (c Code) Response() ResponseBase {
	return ResponseBase{
		Code:    int(c),
		Message: messages[c],
	}
}

const (
	Ok            Code = 0
	UnknownError  Code = 1000
	ArgumentError Code = 1001
	DatabaseError Code = 1002
)

var messages = map[Code]string{
	Ok:            "ok",
	ArgumentError: "argument error",
	DatabaseError: "database error",
}
