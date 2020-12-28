package common

import (
	"github.com/erply/api-go-wrapper/pkg/api/log"
	"time"
)

type Waiter interface {
	Wait(dur time.Duration)
}

type SleepWaiter struct {}

func (sw SleepWaiter) Wait(dur time.Duration) {
	time.Sleep(dur)
}

type Connector struct {
	SessionCleaner        func() error
	Connect               func() (err error)
	AttemptsCount         uint
	Waiter                Waiter
	WaitingInterval       time.Duration
	WaitingIncrementCoeff uint
}

func (c Connector) Run() error {
	var i uint
	var err error
	if c.Waiter == nil {
		c.Waiter = SleepWaiter{}
	}
	for i = 0; i < c.AttemptsCount; i++ {
		err = c.Connect()
		if err == nil {
			return nil
		}

		switch e := err.(type) {
		case *ErplyError:
			if e.Code == APISessionExpired {
				log.Log.Log(log.Error, "failed to connect because auth session is expired: %v", err)
				log.Log.Log(log.Debug, "will invalidate session")
				err := c.SessionCleaner()
				if err != nil {
					return err
				}
			}
		default:
			log.Log.Log(log.Error, "failed to connect: %v", err)
		}

		log.Log.Log(log.Debug, "will retry the connection attempt after %v sleep interval, connections attempts left: %d", c.WaitingInterval, c.AttemptsCount-i)
		c.Waiter.Wait(c.WaitingInterval * time.Duration(c.WaitingIncrementCoeff*i+1))
	}

	return err
}
