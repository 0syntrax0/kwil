package api

import (
	"kwil/api/memcache"
	"kwil/api/model"
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	fileSaveLoc = "../files/"
)

func (a *API) fileCreate(c *gin.Context) {
	// limit file uploads to 5 MiB
	const maxBytes int64 = 5 << 20
	var w http.ResponseWriter = c.Writer
	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, maxBytes)

	// process uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		slog.Error("error accepting uploaded file", "error", err)
		errMsg := "unable to accept file"

		//
		if isErrFileTooLarge(err) {
			errMsg += ". file too large, 5mb max file size"
		}

		a.FailedResponse(c, http.StatusInternalServerError, errMsg)
		return
	}

	// set filename
	filename := filepath.Base(file.Filename)

	// save file to disk
	{
		err := c.SaveUploadedFile(file, fileSaveLoc+filename)
		if err != nil {
			slog.Error("error saving uploaded", "file", filename, "error", err)
			a.FailedResponse(c, http.StatusInternalServerError, "unable to save file, please try again later")
			return
		}
	}

	// generate unique file name
	fName := uuid.New()

	// save file to memcache
	a.Memcache.Insert(fName, memcache.File{
		Name:     filename,
		Location: fileSaveLoc + filename,
		Size:     file.Size,
	})

	// return file name
	a.SuccessResponse(c, model.FileCreateResp{
		Name: fName.String(),
	})
}
