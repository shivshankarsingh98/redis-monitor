package main

import (
	redisMonitor "monitoring"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

var(
	metricFrequency time.Duration = 1
	redisHost = "127.0.0.1"
	redisPort = "6379"
	password = ""
)


func main() {
	redisMonitor.InitializeRedisClent(redisHost,redisPort,password,metricFrequency)

	http.Handle("/ws", websocket.Handler(redisMonitor.StreamRedisMetrics))
	http.HandleFunc("/", redisMonitor.RenderHomePage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}



