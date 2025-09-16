package main

import (
	"testing"
	"time"

	"github.com/tabriz-gulmammadov/snippets/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2025, 9, 16, 11, 45, 0, 0, time.UTC),
			want: "16 Sep 2025 at 11:45",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2025, 9, 16, 11, 45, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "16 Sep 2025 at 10:45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			if hd != tt.want {
				assert.Equal(t, hd, tt.want)
			}
		})
	}
}
