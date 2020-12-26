# redis-monitor

1) Start the redis server

2) Clone the repository
   ```
    git clone https://github.com/shivshankarsingh98/redis-monitor.git
   ```

3) Set GOPATH
   ```
   dir=$(PWD)
   export GOPATH=$dir/redis-monitor/
   ```
   
4) Move to src where main.go is located
   ```
   cd $dir/redis-monitor/src
   ```

5) Set this variable in main.go
   already default values are set to:
   ```
   var (
   
    metricFrequency time.Duration = 1  // how frequently the metrics should be published in seconds 
    redisHost = "127.0.0.1"            // ip address of the redis
    redisPort = "6379"                 // port of the redis
    password = ""                      // password to connect
   
   )   
   ```

6) Download and install third-party Go packages
   ```
   go get -v
   ```


7) Run the app 
   ```
   go run main.go
   ```

8) Open the below address in browser and click button "Stream Redis Metrics"
   ```
   http://127.0.0.1:8080/
   ```

9) We can also test in console, In the dev tools, type the following:
   ```
   var ws = new WebSocket("ws://localhost:8080/ws")
   ws.addEventListener("message", function(e) {console.log(e.data);})
   ```



