package main

import (
	"github.com/tonytheleg/snippetbox/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	/*
		Basic test
		// init new time.Time object and pass to humanDate func
		tm := time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC)
		hd := humanDate(tm)

		// check that output of human date is in format we expect
		// if not return t.Errorf() to indicate test failed
		if hd != "17 Mar 2022 at 10:15" {
			t.Errorf("got %q; want %q", hd, "17 Mar 2022 at 10:15")
	*/

	// Table test
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2022 at 10:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2022 at 09:15",
		},
	}

	for _, tt := range tests {
		// t.Run() function runs a sub-test for each test case
		// the first param is name of test to run
		// second param is function container actual test
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
