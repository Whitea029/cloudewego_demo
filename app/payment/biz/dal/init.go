package dal

import (
	"github.com/Whitea029/whmall/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
