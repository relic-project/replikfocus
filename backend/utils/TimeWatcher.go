package utils

import (
	"sync"
	"time"
)

type TimeWatcher struct {
	Time      time.Time
	Listeners []chan time.Time
	mx        sync.Mutex
}

func NewTimeWatcher() *TimeWatcher {
	return &TimeWatcher{
		Time:      time.Now(),
		Listeners: []chan time.Time{},
	}
}

func (tw *TimeWatcher) AddListener(listener chan time.Time) {
	tw.mx.TryLock()
	defer tw.mx.Unlock()
	tw.Listeners = append(tw.Listeners, listener)
}

func (tw *TimeWatcher) RemoveListener(listener chan time.Time) {
	tw.mx.TryLock()
	defer tw.mx.Unlock()
	for i, l := range tw.Listeners {
		if l == listener {
			tw.Listeners = append(tw.Listeners[:i], tw.Listeners[i+1:]...)
			return
		}
	}
}

func (tw *TimeWatcher) SetTime(t time.Time) {
	tw.Time = t
	tw.mx.TryLock()
	defer tw.mx.Unlock()
	for _, l := range tw.Listeners {
		l <- t
	}
}

func (tw *TimeWatcher) GetTime() time.Time {
	return tw.Time
}
