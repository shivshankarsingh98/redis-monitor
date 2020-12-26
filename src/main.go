package main

import (
	redisMonitor "monitoring"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

var(
	metricFrequency time.Duration = 1  // how frequently the metrics should be published in seconds
	redisHost = "127.0.0.1"            // ip address of the redis
	redisPort = "6379"                 // port of the redis
	password = ""                      // password to connect
)


func main() {
	redisMonitor.InitializeRedisClent(redisHost,redisPort,password,metricFrequency)

	http.Handle("/ws", websocket.Handler(redisMonitor.StreamRedisMetrics))
	http.HandleFunc("/", redisMonitor.RenderHomePage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}



