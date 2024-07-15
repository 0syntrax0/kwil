package api

import "errors"

// check if the error given is related to a file being too large
func isErrFileTooLarge(err error) bool {
	ftl := errors.New("http: request body too large")
	return errors.Is(err, ftl)
}
