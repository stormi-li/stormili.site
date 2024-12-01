package main

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiserd-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	omiserdC := omiserd.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	}, omiserd.Config)
	r := omiserdC.NewRegister("mysql", "118.25.196.166:3933")
	r.Data["username"] = "root"
	r.Data["database"] = "USER"
	r.Data["password"] = "12982397StrongPassw0rd"
	r.Data["processid"] = strconv.Itoa(os.Getpid())
	r.Register(1)
	select {}
}
