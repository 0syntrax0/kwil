package api

import "github.com/gin-gonic/gin"

func (a *API) initRoutes(r *gin.Engine) {
	r.POST("/file", a.fileCreate)    // create file
	r.GET("/file/:uid", a.fileFetch) // get file
	r.GET("/files", a.fileFetchAll)  // Bonus: get all saved files
}
