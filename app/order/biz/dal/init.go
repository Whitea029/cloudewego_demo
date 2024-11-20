package dal

import (
	"github.com/Whitea029/whmall/app/order/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
