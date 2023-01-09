package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestValidateAvailabilitiesNotEnoughMondaysThursdays(t *testing.T) {
	err := validateAvailabilities([]time.Time{})
	if err.Error() != "expected at least 4 Mondays and 4 Thursdays" {
		t.Fatal("Expected failure due to not enough Mondays and Thursdays")
	}
}

func TestValidateAvailabilitiesInputIncludesInvalidDate(t *testing.T) {
	err := validateAvailabilities([]time.Time{CHRISTMAS})
	if !strings.Contains(err.Error(), "expected no availabilities on Christmas") {
		t.Fatal("Expected failure due to invalid date")
	}
}

func TestValidateAvailabilitiesInputTooSmall(t *testing.T) {
	dates := []time.Time{}
	curDate := NCAA_CHMP.AddDate(0, 0, 1)
	for i := 0; i < 49; i++ {
		dates = append(dates, curDate)
		curDate = curDate.AddDate(0, 0, 1)
	}

	err := validateAvailabilities(dates)
	if err.Error() != fmt.Sprintf(`expected at least %d availabilities`, MIN_AVAILABLE_DAYS) {
		t.Fatal("Expected failure due to input size")
	}
}
