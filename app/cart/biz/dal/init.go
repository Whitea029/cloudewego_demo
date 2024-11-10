package dal

import (
	"github.com/Whitea029/whmall/app/cart/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
