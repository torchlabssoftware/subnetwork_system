package server

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type OTP struct {
	Key     string
	Created time.Time
}

type RetentionMap map[string]*OTP

func NewRetentionMap(context context.Context, retentionPeriod time.Duration) RetentionMap {
	rm := make(RetentionMap)
	go rm.Retention(context, retentionPeriod)
	return rm
}

func (rm RetentionMap) NewOTP() OTP {
	o := OTP{
		Key:     uuid.NewString(),
		Created: time.Now(),
	}
	rm[o.Key] = &o
	return o
}

func (rm RetentionMap) VerifyOTP(key string) bool {
	if _, ok := rm[key]; !ok {
		return false
	}
	delete(rm, key)
	return true
}

func (rm RetentionMap) Retention(ctx context.Context, timePeriod time.Duration) {
	ticker := time.NewTicker(400 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			for _, otp := range rm {
				if otp.Created.Add(timePeriod).Before(time.Now()) {
					delete(rm, otp.Key)
				}
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}
