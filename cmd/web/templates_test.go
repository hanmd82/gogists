package main

import (
	"testing"
	"time"
)

func TestFormatDateTime(t *testing.T) {
	tm := time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC)
	fdt := formatDateTime(tm)

	if fdt != "17 Dec 20 10:00 UTC" {
		t.Errorf("want %q; got %q", "17 Dec 20 10:00 UTC", fdt)
	}
}
