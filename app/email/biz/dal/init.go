package dal

import (
	"github.com/Whitea029/whmall/app/email/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
