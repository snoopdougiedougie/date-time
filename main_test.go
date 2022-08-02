package main

import (
	"testing"
)

func TestDateTime(t *testing.T) {
	myDateTime := newDateTime(2022, 8, 2, 18, 7, 0, 0)
	myDateTimeString := getDateTimeAsString(*myDateTime)

	correctDtString := "2022-08-02T18:07:00:00Z"

	if myDateTimeString != correctDtString {
		t.Errorf("got %q, wanted %q", myDateTimeString, correctDtString)
	}
}
