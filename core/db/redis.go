/*
@Author : Bruce Bee
@Date   : 2020/1/2 17:38
@Email  : mzpy_1119@126.com
*/

// package db is ...
package db

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/garyburd/redigo/redis"
)

// InitRedis ...
func InitRedis() (redisConn redis.Conn,  err error){
	cfg, err := goconfig.LoadConfigFile("conf/app.ini")

	if err != nil {
		panic("panic")
	}

	redis_host, err := cfg.GetValue("redis","redis_host")
	redis_port, err := cfg.Int("redis","redis_port")
	//redis_user, err := cfg.GetValue("redis","redis_user")
	//redis_pwd, err := cfg.GetValue("redis","redis_pwd")
	//redis_db, err := cfg.GetValue("redis","redis_db")

	url := fmt.Sprintf("%s:%d",
		redis_host,
		redis_port)

	redisConn, err = redis.Dial("tcp",url)

	if err != nil {
		fmt.Println("connect redis error :",err)
	}
	return
}
