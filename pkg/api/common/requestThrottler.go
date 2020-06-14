package common

import (
	"sync"
	"time"
)

//Throttler abstracts limiting of API requests
type Throttler interface {
	Throttle()
}

type Sleeper func(sleepTime time.Duration)

var sleepThrottler *SleepThrottler

//SleepThrottler implements sleeping logic for requests throttling
type SleepThrottler struct {
	LimitPerSecond uint
	LastTimestamp  int64
	Count          uint
	sl             Sleeper
	lock           sync.Mutex
}

//NewSleepThrottler creates SleepThrottler
func NewSleepThrottler(limitPerSecond uint, sl Sleeper) *SleepThrottler {
	if sleepThrottler == nil {
		sleepThrottler = &SleepThrottler{
			LimitPerSecond: limitPerSecond,
			LastTimestamp:  time.Now().Unix(),
			Count:          0,
			sl:             sl,
			lock:           sync.Mutex{},
		}
	}

	return sleepThrottler
}

//Throttle implements throttling method
func (rt *SleepThrottler) Throttle() {
	rt.lock.Lock()
	defer rt.lock.Unlock()

	if rt.LimitPerSecond <= 0 {
		return
	}

	rt.Count++
	now := time.Now().Unix()
	if now != rt.LastTimestamp {
		rt.LastTimestamp = now
		rt.Count = 1
		return
	}

	if rt.Count >= rt.LimitPerSecond {
		rt.sl(time.Second)
	}
}

type ThrottlerMock struct {
	WasTriggered bool
}

func (tm *ThrottlerMock) Throttle(){
	tm.WasTriggered = true
}
