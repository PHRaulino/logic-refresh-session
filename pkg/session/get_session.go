package session

import (
	"errors"
	"time"
)

func findClosestDateBefore(dates []time.Time, reference time.Time, firstToken time.Time) time.Time {
	closest := firstToken
	for _, date := range dates {
		if date.Before(reference) || date == reference {
			closest = date
		}
	}
	return closest
}

func GetTokenSession(
	refreshTime time.Duration,
	sessionTime time.Duration,
	currentTime time.Time,
	firstToken time.Time,
) (time.Time, error) {
	steps, err := GetRefreshSteps(
		refreshTime,
		firstToken,
		currentTime.Add(time.Duration(sessionTime.Minutes())*time.Minute),
	)
	if err != nil {
		return firstToken, errors.New("não foi possivel recuperar os steps")
	}

	return findClosestDateBefore(steps, currentTime, firstToken), nil
}
