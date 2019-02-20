package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/daved/sandbox/goplgreg/e0410/github"
)

type ghi = github.Issue

func msAgo(n int) time.Time {
	return time.Now().Add(time.Millisecond * -1 * time.Duration(n))
}

func TestFilterIssues(t *testing.T) {
	now := msAgo(0)
	d := []*ghi{
		{CreatedAt: msAgo(20)},
		{CreatedAt: msAgo(4000)},
		{CreatedAt: msAgo(10)},
		{CreatedAt: msAgo(12000)},
		{CreatedAt: msAgo(1200)},
		{CreatedAt: msAgo(15000)},
	}

	tests := []struct {
		name string
		fn   func(time.Time) bool
		want []*github.Issue
	}{
		{"lt 1 sec", now.Add(-time.Second).Before, []*ghi{d[0], d[2]}},
		{"mt 2 sec", now.Add(-2 * time.Second).After, []*ghi{d[1], d[3], d[5]}},
	}

	for _, tt := range tests {
		got := filterIssues(tt.fn, d)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: got %v, want %v", tt.name, got, tt.want)
		}
	}
}
