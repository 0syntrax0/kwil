package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errResponse struct {
	// error message response
	ErrorMessage string
}

// Note: these are just some basic response functions
// in production, this would be a more advance library handling many different types of return types and options
func (a *API) SuccessResponse(c *gin.Context, resp any) {
	c.JSON(http.StatusOK, resp)
}

func (a *API) FailedResponse(c *gin.Context, httpCode int, msg string) {
	c.JSON(httpCode, errResponse{
		ErrorMessage: msg,
	})
}
