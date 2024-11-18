package main

import (
	"comment/database"
	"comment/routes"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiserd-v1"
)
var redisAddr = "118.25.196.166:3934"
var redisPassword = "12982397StrongPassw0rd"
func main() {

	// 初始化数据库
	database.InitDatabase()

	r := routes.SetupRouter()
	omiserd.NewClient(&redis.Options{Addr: redisAddr, Password: redisPassword}, omiserd.Server).NewRegister("comment_service", "118.25.196.166:9091").RegisterAndListen(1, func(port string) {
		r.Run(port)
	})
	// 启动服务
}
