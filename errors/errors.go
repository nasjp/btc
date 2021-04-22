package errors

import "fmt"

type ValueError string

func ValueErrorf(format string, a ...interface{}) ValueError {
	return ValueError(fmt.Sprintf(format, a...))
}

func (e ValueError) Error() string { return string(e) }

type TypeError string

func TypeErrorf(format string, a ...interface{}) TypeError {
	return TypeError(fmt.Sprintf(format, a...))
}

func (e TypeError) Error() string { return string(e) }
