package main

import (
	"snippetbox.abdou-salama-001.net/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2025, 1, 9, 12, 10, 0, 0, time.UTC),
			want: "09 Jan 2025 at 12:10",
		},
		{
			name: "empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2025, 1, 9, 12, 10, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "09 Jan 2025 at 11:10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			assert.Equal(t, hd, tt.want)
		})
	}
}
