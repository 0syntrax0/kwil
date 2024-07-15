package model

import "kwil/api/memcache"

// api json response fields
type FileFetchAllResp struct {
	UID  string
	Data memcache.File
}
