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
	var err error

	for {

		redisStats := ExampleClient()
		b, err1 := json.Marshal(redisStats)
		if err1 != nil {
			panic(err)
		}

		if err = websocket.Message.Send(ws, string(b)); err != nil {
			log.Print("Can't send")
			break
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

	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
