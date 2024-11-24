package dal

import (
	"github.com/Whitea029/whmall/app/product/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
