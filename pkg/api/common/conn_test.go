package common

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type WaiterMock struct {
	WaitingDurations []time.Duration
}

func (wm *WaiterMock) Wait(dur time.Duration) {
	wm.WaitingDurations = append(wm.WaitingDurations, dur)
}

func TestSingleConnectSuccess(t *testing.T) {
	sessionCleanerWasCalled := false
	connectAmount := 0
	w := &WaiterMock{}
	const connectsCount = 5
	c := Connector{
		SessionCleaner: func() error {
			sessionCleanerWasCalled = true
			return nil
		},
		Connect: func() (err error) {
			connectAmount++
			return
		},
		AttemptsCount:         connectsCount,
		Waiter:                w,
	}

	err := c.Run()
	assert.NoError(t, err)

	assert.False(t, sessionCleanerWasCalled)

	assert.Equal(t, 1, connectAmount)

	assert.Len(t, w.WaitingDurations, 0)
}

func TestSessionExpiration(t *testing.T) {
	sessionCleanerWasCalled := false
	connectAmount := 0
	w := &WaiterMock{}
	const connectsCount = 5
	c := Connector{
		SessionCleaner: func() error {
			sessionCleanerWasCalled = true
			return nil
		},
		Connect: func() (err error) {
			if connectAmount == 0 {
				err = &ErplyError{
					error:   errors.New("conn failure"),
					Status:  "Some status",
					Message: "Some message",
					Code:    APISessionExpired,
				}
			}
			connectAmount++
			return
		},
		AttemptsCount:         connectsCount,
		Waiter:                w,
		WaitingInterval:       time.Second,
		WaitingIncrementCoeff: 10,
	}

	err := c.Run()
	if err != nil {
		t.Error(err)
	}

	assert.True(t, sessionCleanerWasCalled)
	assert.Equal(t, 2, connectAmount)
	assert.Len(t, w.WaitingDurations, 1)
	assert.Equal(t, time.Second, w.WaitingDurations[0])
}

func TestMultipleConnectFailure(t *testing.T) {
	sessionCleanerWasCalled := false
	connectAmount := 0
	w := &WaiterMock{}
	const connectsCount = 5
	c := Connector{
		SessionCleaner: func() error {
			sessionCleanerWasCalled = true
			return nil
		},
		Connect: func() (err error) {
			connectAmount++
			err = errors.New("conn failure")
			return
		},
		AttemptsCount:         connectsCount,
		Waiter:                w,
		WaitingInterval:       time.Second,
		WaitingIncrementCoeff: 2,
	}

	err := c.Run()
	assert.Error(t, err, "conn failure")
	assert.False(t, sessionCleanerWasCalled)
	assert.Equal(t, 5, connectAmount)
	assert.Len(t, w.WaitingDurations, 5)
	assert.Equal(t, time.Second, w.WaitingDurations[0])
	assert.Equal(t, time.Second * 3, w.WaitingDurations[1])
	assert.Equal(t, time.Second * 5, w.WaitingDurations[2])
	assert.Equal(t, time.Second * 7, w.WaitingDurations[3])
	assert.Equal(t, time.Second * 9, w.WaitingDurations[4])
}

func TestSessionCleanFailure(t *testing.T) {
	connectAmount := 0
	w := &WaiterMock{}
	const connectsCount = 5
	c := Connector{
		SessionCleaner: func() error {
			return errors.New("some session failure")
		},
		Connect: func() (err error) {
			if connectAmount == 0 {
				err = &ErplyError{
					error:   errors.New("conn failure"),
					Status:  "Some status",
					Message: "Some message",
					Code:    APISessionExpired,
				}
			}
			connectAmount++
			return
		},
		AttemptsCount:         connectsCount,
		Waiter:                w,
	}

	err := c.Run()
	assert.Error(t, err, "some session failure")
	assert.Equal(t, 1, connectAmount)
	assert.Len(t, w.WaitingDurations, 0)
}
