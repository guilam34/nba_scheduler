package main

import (
	"errors"
	"fmt"
	"time"
)

var CHRISTMAS = time.Date(2022, time.December, 25, 0, 0, 0, 0, time.UTC)
var ASG_START = time.Date(2023, time.February, 17, 0, 0, 0, 0, time.UTC)
var ASG_END = time.Date(2023, time.February, 19, 0, 0, 0, 0, time.UTC)
var NCAA_CHMP = time.Date(2023, time.April, 3, 0, 0, 0, 0, time.UTC)

const MIN_AVAILABLE_DAYS = 50

func validateAvailabilities(availabilities []time.Time) error {
	dedupedAvail := deduplicateAvailabilities(availabilities)

	// Min 4 Mondays and Thursdays
	var numMondays = 0
	var numThursdays = 0
	for _, day := range dedupedAvail {
		if day.Weekday() == time.Monday {
			numMondays++
		}
		if day.Weekday() == time.Thursday {
			numThursdays++
		}
		// No Christmas days, ASG weekend, or NCAA championship game
		if day.Equal(CHRISTMAS) || (!day.Before(ASG_START) && !day.After(ASG_END)) || day.Equal(NCAA_CHMP) {
			var format = "2022/12/01"
			return fmt.Errorf("expected no availabilities on Christmas (%q), All-Star Game Weekend (%q-%q), or during the NCAA Champaionship Game %q", CHRISTMAS.Format(format), ASG_START.Format(format), ASG_END.Format(format), NCAA_CHMP.Format(format))
		}

	}
	if numMondays < 4 || numThursdays < 4 {
		return errors.New("expected at least 4 Mondays and 4 Thursdays")
	}

	// Min 50 valid availabilities
	if len(dedupedAvail) < MIN_AVAILABLE_DAYS {
		return fmt.Errorf(`expected at least %d availabilities`, MIN_AVAILABLE_DAYS)
	}

	return nil
}

func deduplicateAvailabilities(availabilities []time.Time) []time.Time {
	keys := make(map[time.Time]bool)
	list := []time.Time{}

	for _, entry := range availabilities {
		if _, value := keys[entry]; !value {
			list = append(list, entry)
			keys[entry] = true
		}
	}
	return list
}
