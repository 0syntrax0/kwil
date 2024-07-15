package model

import (
	"errors"
	"mime/multipart"
)

var (
	ErrFileUploadMissingFile = errors.New("missing file")
)

// api json field needed to upload a file
type FileUploadReq struct {
	File *multipart.FileHeader
}

// check for requried fields
func (fur FileUploadReq) Validate() error {
	if fur.File == nil {
		return ErrFileUploadMissingFile
	}

	return nil
}
