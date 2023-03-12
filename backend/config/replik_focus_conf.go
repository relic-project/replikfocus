package config

import "time"

type ReplikFocusConf struct {
	WorkDuration      time.Duration
	BreakDuration     time.Duration
	LongBreakDuration time.Duration
	BreaksBeforeLong  int
}
