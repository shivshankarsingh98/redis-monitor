package monitoring

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"strings"
	"time"
)

var (
	RedisDb *redis.Client
	Ctx = context.Background()
	metricDescription = GetMetricDescription()
	redisMetricChannel = make(chan RedisKeyInfo)
	metricFreq time.Duration
)

func InitializeRedisClent(host, port, password string, metricFrequency time.Duration)  {
	RedisDb = redis.NewClient(&redis.Options{
		Addr: host+":"+port,
		Password: password,
		DB:       0,  // use default DB
	})
	if err := RedisDb.Ping(Ctx).Err(); err != nil {
		log.Print("Unable to connect to redis: " + err.Error())
		return
	}
	metricFreq  = metricFrequency
}

func GetRedisKeyInfo(redisInfoKey string){
	val, err := RedisDb.Do(Ctx,"info", redisInfoKey).Result()
	if err != nil {
		log.Print("Unable get redis info: " + err.Error())
		return
	}

	redisKeyInfoArr := []MetricDetails{}
	metricsNameWithValue := strings.Split(val.(string), "\n")
	for _, metricNameAndValue := range metricsNameWithValue{
		splitValues := strings.Split(metricNameAndValue, ":")
		if len(splitValues) == 2 {
			metricName := strings.TrimSpace(splitValues[0])
			metricValue := strings.TrimSpace(splitValues[1])
			metricInfo := metricDescription[metricName]

			metricDetails := MetricDetails{MetricName: metricName, MetricDescription: metricInfo, MetricValue: metricValue}
			redisKeyInfoArr = append(redisKeyInfoArr, metricDetails)
		}
	}

	redisMetricChannel <- RedisKeyInfo{redisInfoKey,redisKeyInfoArr}
}

func processKeys(redisInfoKeys []string) map[string]interface{} {

	redisStats := map[string]interface{}{}
	for _, redisInfoKey := range redisInfoKeys {
		go GetRedisKeyInfo(redisInfoKey)
	}
	for i:=0 ; i<len(redisInfoKeys) ; i ++ {
		metrics := <- redisMetricChannel
		redisStats[metrics.KeyName] = metrics.KeyMetrics
	}
	return redisStats
}

func ExampleClient() map[string]interface{}{
	redisInfoKeys := []string{"CPU", "Clients", "Cluster", "Server", "Memory"}
	redisStats := processKeys(redisInfoKeys)
	return redisStats
}