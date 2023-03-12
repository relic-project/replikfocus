package main

import (
	"log"
	"net/http"
	"os"
	"replikfocus/config"
	"replikfocus/msg"
	"replikfocus/utils"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var timeWatcher = utils.NewTimeWatcher()

var pref = config.ReplikFocusConf{
	WorkDuration:      time.Duration(25) * time.Minute,
	BreakDuration:     time.Duration(5) * time.Minute,
	LongBreakDuration: time.Duration(15) * time.Minute,
	BreaksBeforeLong:  4,
}

/*var pref = config.ReplikFocusConf{
	WorkDuration:      time.Duration(30) * time.Second,
	BreakDuration:     time.Duration(10) * time.Second,
	LongBreakDuration: time.Duration(15) * time.Second,
	BreaksBeforeLong:  4,
}*/

var CurrentMode = ""
var CurrentBreak = 0

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	if err := ws.WriteJSON(msg.NewNewTimerMsg(CurrentMode, timeWatcher.GetTime())); err != nil {
		log.Println(err)
		return err
	}

	tChan := make(chan time.Time)
	timeWatcher.AddListener(tChan)
	defer timeWatcher.RemoveListener(tChan)
	for {
		select {
		case t := <-tChan:
			ws.WriteJSON(msg.NewPlaySoundMsg("start"))
			err := ws.WriteJSON(msg.NewNewTimerMsg(CurrentMode, t))
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}

}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../frontend/dist/")
	e.GET("/ws", hello)

	go func() {
		// timer to handle work time, break time, long break time

		for {
			if CurrentMode != "work" {
				CurrentMode = "work"
				go timeWatcher.SetTime(time.Now().Add(pref.WorkDuration))
				log.Println("Starting work timer")
				time.Sleep(pref.WorkDuration)
			} else if CurrentMode != "break" && CurrentMode != "longbreak" {
				toadd := pref.BreakDuration
				if CurrentBreak == pref.BreaksBeforeLong {
					CurrentMode = "longbreak"
					CurrentBreak = 0
					toadd = pref.LongBreakDuration
					log.Println("Starting long break timer")
				} else {
					CurrentMode = "break"
					CurrentBreak++
					log.Println("Starting break timer")
				}
				go timeWatcher.SetTime(time.Now().Add(toadd))
				time.Sleep(toadd)
			}
			// sleep until timewatcher.getTime

		}
	}()

	host := ""
	port := "1323"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	}

	e.Logger.Fatal(e.Start(host + ":" + port))

}
