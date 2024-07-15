package api

import (
	"kwil/api/memcache"

	"github.com/gin-gonic/gin"
)

type API struct {
	*gin.Engine
	Memcache *memcache.Memcache
}

func Run() {
	var api API

	// ini what would've been the DB
	api.Memcache = memcache.New()

	// creates a router without any middleware by default
	r := gin.New()
	r.SetTrustedProxies([]string{localSchema + localhost + localPort})

	// global middleware
	r.Use(gin.Logger())

	// recovery middleware recovers from any panics and logs a 500 if there was one.
	r.Use(gin.Recovery())

	// init url routes
	api.initRoutes(r)

	// run server
	r.Run(localPort)
}
