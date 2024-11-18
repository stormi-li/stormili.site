package main

import (
	"github.com/go-redis/redis/v8"
	omiweb "github.com/stormi-li/omiweb-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	omiweb := omiweb.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	omiweb.GenerateTemplate()
	ws := omiweb.NewWebServer("stormili.site", "118.25.196.166:7070")
	ws.Start(1)
}
