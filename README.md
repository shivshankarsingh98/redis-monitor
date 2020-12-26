# redis-monitor

cd workspace/

git clone https://github.com/shivshankarsingh98/redis-monitor.git

cd redis-monitor/src


export GOPATH=/Users/manasasingh/Downloads/workspace/redis-monitor/

go get -v

set this variable in main.go

var (
metricFrequency time.Duration = 1  // how frequently the metrics should be published in seconds
redisHost = "127.0.0.1"            // ip address of the redis
redisPort = "6379"                 // port of the redis
password = ""                      // password to connect
)

go run main.go

in browser open http://127.0.0.1:8080/

we can also test in console
In the dev tools, type the following:

var ws = new WebSocket("ws://localhost:8080/ws")
ws.addEventListener("message", function(e) {console.log(e.data);})


