package monitoring

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
	"time"
)

func StreamRedisMetrics(ws *websocket.Conn) {
	clentAddress := ws.RemoteAddr().String()
	log.Print("Connected to client: ", clentAddress)

	for {
		redisStats := GetRedisMetrics()
		redisStatsJson, marshalErr := json.Marshal(redisStats)
		if marshalErr != nil {
			log.Print("Error json marshaling: ",marshalErr)
			return
		}

		if err := websocket.Message.Send(ws, string(redisStatsJson)); err != nil {
			log.Print("Can't send metrics to clent: ",clentAddress)
			return
		}
		time.Sleep(time.Second*metricFreq)
	}
}

type PageVariables struct {
	Date         string
	Time         string
}

func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	HomePageVars := PageVariables{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	indexTemplate, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Print("Template parsing error: ", err)
	}

	err = indexTemplate.Execute(w, HomePageVars)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}
