package dal

import (
	"judgeMore/biz/dal/cache"
	"judgeMore/biz/dal/mysql"
)

func Init() {
	mysql.Init()
	cache.Init()
}
