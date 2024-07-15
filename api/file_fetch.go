package api

import (
	"io"
	"kwil/api/model"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *API) fileFetch(c *gin.Context) {
	// check for valid UUID
	uid, err := uuid.Parse(c.Param("uid"))
	if err != nil {
		slog.Warn("invalid", "uuid", c.Param("uid"))
		a.FailedResponse(c, http.StatusBadRequest, "invalid file key")
		return
	}

	// fetch file from storage
	file, err := a.Memcache.Select(uid)
	if err != nil {
		slog.Warn(err.Error())
		a.FailedResponse(c, http.StatusNotFound, "file not found")
		return
	}
	fileSavedLoc := file.Name

	// force browser to download file
	header := c.Writer.Header()
	header["Content-type"] = []string{"application/octet-stream"}
	header["Content-Disposition"] = []string{"attachment; filename= " + file.Name}

	// return file
	{
		file, err := os.Open(fileSaveLoc + fileSavedLoc)
		if err != nil {
			slog.Warn("file not found", "uuid", uid.String(), "fileLocation", fileSavedLoc)
			a.FailedResponse(c, http.StatusNotFound, "file not found")
			return
		}
		defer file.Close()
		io.Copy(c.Writer, file)
	}
}

func (a *API) fileFetchAll(c *gin.Context) {
	resp := make([]model.FileFetchAllResp, 0)

	for uid, file := range a.Memcache.SelectAll() {
		resp = append(resp, model.FileFetchAllResp{
			UID:  uid.String(),
			Data: file,
		})
	}

	// return all files currently in memcache
	a.SuccessResponse(c, resp)
}
