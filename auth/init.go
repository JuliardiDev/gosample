package auth

import (
	session_helper "github.com/tokopedia/helper/session"
)

var SessionHelper *session_helper.RedisCache

func init() {
	if SessionHelper == nil {
		//initilize the SessionHelper
		//the parameters are the redis that being use for the session helper
		//for real life, it uses the redis in biznet or aws
		//for this time, just use same redis on devel.
		SessionHelper = session_helper.NewRedisSession("devel-redis.tkpd:6379", "devel-redis.tkpd:6379")
	}
}
