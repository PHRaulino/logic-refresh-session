// refresh_steps_test.go

package session

import (
	"reflect"
	"testing"
	"time"
)

func TestGetRefreshSteps(t *testing.T) {
	startTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	interval := 4 * time.Minute

	tests := []struct {
		name      string
		interval  time.Duration
		startTime time.Time
		endTime   time.Time
		want      []time.Time
		wantErr   bool
	}{
		{
			"Return array empty",
			interval,
			startTime,
			startTime.Add(3 * time.Minute),
			nil,
			false,
		},
		{
			"Return error with StartTime greater endTime",
			interval,
			startTime.Add(6 * time.Minute), startTime,
			nil,
			true,
		},
		{
			"Return empty array with start and endtime same date",
			interval,
			startTime, startTime,
			nil,
			false,
		},
		{
			"Return array with 1 refresh time",
			interval,
			startTime,
			startTime.Add(6 * time.Minute),
			[]time.Time{
				startTime.Add(4 * time.Minute),
			},
			false,
		},
		{
			"Return array with 3 refresh time",
			interval,
			startTime,
			startTime.Add(15 * time.Minute),
			[]time.Time{
				startTime.Add(4 * time.Minute),
				startTime.Add(8 * time.Minute),
				startTime.Add(12 * time.Minute),
			},
			false,
		},
		{
			"Return array with 4 refresh time",
			interval,
			startTime,
			startTime.Add(18 * time.Minute),
			[]time.Time{
				startTime.Add(4 * time.Minute),
				startTime.Add(8 * time.Minute),
				startTime.Add(12 * time.Minute),
				startTime.Add(16 * time.Minute),
			},
			false,
		},
		{
			"Return array with 3 refresh time with others periods",
			interval,
			startTime.Add(3 * time.Minute),
			startTime.Add(18 * time.Minute),
			[]time.Time{
				startTime.Add(7 * time.Minute),
				startTime.Add(11 * time.Minute),
				startTime.Add(15 * time.Minute),
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRefreshSteps(tt.interval, tt.startTime, tt.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRefreshSteps() expect error ? %v, error receive: '%v'", tt.wantErr, err)
			} else if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRefreshSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
