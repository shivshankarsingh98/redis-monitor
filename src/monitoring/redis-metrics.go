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

func GetRedisKeyInfo(redisInfoKeyName string){
	infoVal, err := RedisDb.Do(Ctx,"info", redisInfoKeyName).Result()
	if err != nil {
		log.Print("Unable get redis info: " + err.Error())
		return
	}

	redisKeyMetrics := []MetricDetails{}
	metricsList := strings.Split(infoVal.(string), "\n")
	for _, metricNameWithValue := range metricsList{
		splitValues := strings.Split(metricNameWithValue, ":")
		if len(splitValues) == 2 {
			metricName := strings.TrimSpace(splitValues[0])
			metricValue := strings.TrimSpace(splitValues[1])
			metricInfo := metricDescription[metricName]

			metricDetails := MetricDetails{MetricName: metricName, MetricDescription: metricInfo, MetricValue: metricValue}
			redisKeyMetrics = append(redisKeyMetrics, metricDetails)
		}
	}

	redisMetricChannel <- RedisKeyInfo{redisInfoKeyName,redisKeyMetrics}
}

func processKeys(redisInfoKeys []string) map[string]interface{} {
	for _, redisInfoKey := range redisInfoKeys {
		go GetRedisKeyInfo(redisInfoKey)
	}

	redisStats := map[string]interface{}{}
	for i:=0 ; i<len(redisInfoKeys) ; i ++ {
		infoKeyStats := <- redisMetricChannel
		redisStats[infoKeyStats.KeyName] = infoKeyStats.KeyMetrics
	}
	return redisStats
}

func GetRedisMetrics() map[string]interface{} {
	redisInfoKeys := []string{"CPU", "Clients", "Cluster", "Server", "Memory"}
	redisStats := processKeys(redisInfoKeys)
	return redisStats
}