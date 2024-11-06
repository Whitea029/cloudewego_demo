package dal

import (
	"github.com/Whitea029/whmall/app/user/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
