package dal

import (
	"github.com/Whitea029/whmall/app/checkout/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
