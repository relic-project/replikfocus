package main

import (
	"log"
	"net/http"
	"os"
	"replikfocus/config"
	"replikfocus/msg"
	"replikfocus/users"
	"replikfocus/utils"
	"strconv"
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

var connectedUsers = users.NewConnectedUsers()

func hello(c echo.Context) error {
	defer log.Println("QUITTING ")
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	profile := new(msg.RegisterMsg)

	if err := ws.ReadJSON(profile); err != nil {
		log.Println(err)
		return err
	}

	if profile.Username == "" {
		profile.Username = "Anonymous"
	}

	if connectedUsers.Contains(profile.Username) {
		for i := 0; true; i++ {
			profile.Username = profile.Username + strconv.Itoa(i)
			if !connectedUsers.Contains(profile.Username) {
				profile.Username = profile.Username + strconv.Itoa(i)
				break
			}
		}
	}

	connectedUsersListener := make(chan []string)
	tChan := make(chan time.Time)
	defer func() {
		connectedUsers.RemoveListener(connectedUsersListener)
		connectedUsers.Remove(profile.Username)
		defer timeWatcher.RemoveListener(tChan)

	}()

	connectedUsers.Add(profile.Username)
	connectedUsers.AddListener(connectedUsersListener)
	log.Println("New user connected: " + profile.Username)

	if err := ws.WriteJSON(msg.NewRegisterMsg(profile.Username)); err != nil {
		log.Println(err)
		return err
	}

	if err := ws.WriteJSON(msg.NewConnectedUsersMsg(connectedUsers.Get())); err != nil {
		log.Println(err)
		return err
	}

	if err := ws.WriteJSON(msg.NewNewTimerMsg(CurrentMode, timeWatcher.GetTime())); err != nil {
		log.Println(err)
		return err
	}

	timeWatcher.AddListener(tChan)

	for {
		if err := ws.WriteJSON(msg.Message{Type: "ping"}); err != nil {
			log.Println(err)
			return err
		}
		select {
		case t := <-tChan:
			ws.WriteJSON(msg.NewPlaySoundMsg("start"))
			err := ws.WriteJSON(msg.NewNewTimerMsg(CurrentMode, t))
			if err != nil {
				log.Println(err)
				return err
			}
		case users := <-connectedUsersListener:
			err := ws.WriteJSON(msg.NewConnectedUsersMsg(users))
			if err != nil {
				log.Println(err)
				return err
			}
		default:
			continue
		}
		time.Sleep(1 * time.Second)
	}
	return nil
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
