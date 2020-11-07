package helpers

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

// TODO: remove all these functions
func TimeToTs(time *time.Time) *timestamp.Timestamp {
	if time == nil {
		return nil
	}
	t := time.UTC()
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func TimeToTs2(time time.Time) *timestamp.Timestamp {
	return TimeToTs(&time)
}

func TsToTime(ts *timestamp.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	res := &time.Time{}
	*res = time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
	return res
}

func TsToTime2(ts *timestamp.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	v := TsToTime(ts)
	return *v
}

func Now() *time.Time {
	now := time.Now()
	return &now
}
