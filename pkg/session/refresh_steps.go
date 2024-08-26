package session

import (
	"errors"
	"math"
	"time"
)

func GetRefreshSteps(interval time.Duration, startTime, endTime time.Time) ([]time.Time, error) {
	var steps []time.Time

	if endTime.Before(startTime) {
		return nil, errors.New("startTime date greater endTime")
	}

	if endTime == startTime || endTime.Sub(startTime) <= interval {
		return steps, nil
	}

	lastStep := endTime.Add(-5 * time.Minute)
	intervalSteps := lastStep.Sub(startTime)
	minutes := intervalSteps.Minutes()
	roundedMinutes := math.Ceil(minutes)
	stepsRange := int(math.Ceil(roundedMinutes / interval.Minutes()))

	for i := 1; i <= stepsRange; i++ {
		refreshDate := startTime.Add(time.Duration(i) * interval)
		steps = append(steps, refreshDate)
	}

	return steps, nil
}
