package main

import (
	"auth/database"
	"auth/routes"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiserd-v1"
)

var redisAddr = "118.25.196.166:3934"
var redisPassword = "12982397StrongPassw0rd"
var serverName = "auth_service"

func main() {
	// Initialize database
	database.InitDatabase()

	// Setup routes
	r := routes.SetupRouter()

	// Start server
	omiserd.NewClient(&redis.Options{Addr: redisAddr, Password: redisPassword}, omiserd.Server).NewRegister(serverName, "118.25.196.166:9090").RegisterAndListen(1, func(port string) {
		r.Run(port) // 默认监听 8080 端口
	})
}
