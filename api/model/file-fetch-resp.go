package model

import "kwil/api/memcache"

// api json response fields
type FileFetchAllResp struct {
	UID  string        `json:"uid"`
	Data memcache.File `json:"data"`
}
