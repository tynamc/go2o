/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2014-02-05 21:53
 * description :
 * history :
 */
package server

import (
	"github.com/jrsix/gof/net/jsv"
	"github.com/garyburd/redigo/redis"
	"go2o/src/core"
)

var (
	_redis *redis.Pool
)

func Redis() *redis.Pool {
	if _redis == nil {
		gc := jsv.Context.(*core.MainApp)
		_redis = gc.Redis()
	}
	return _redis
}
