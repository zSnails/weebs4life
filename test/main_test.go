package main

import (
	"testing"

	w4l ".."
)

var (
	weebClient = w4l.NewWeeb()
)

//TestHug test a hug
func TestHug(t *testing.T) {
	_, err := weebClient.Hug()
	if err != nil {
		t.Error(err)
	}
}

func TestNeko(t *testing.T) {
	_, err := weebClient.Neko()
	if err != nil {
		t.Error(err)
	}
}

func TestKiss(t *testing.T) {
	_, err := weebClient.Slap()
	if err != nil {
		t.Error(err)
	}
}

func TestSlap(t *testing.T) {
	_, err := weebClient.Slap()
	if err != nil {
		t.Error(err)
	}
}
