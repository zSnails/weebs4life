package main

import (
	"testing"

	w4l ".."
)

var (
	wc = w4l.NewWeeb()
)

//TestHug test a hug
func TestHug(t *testing.T) {
	_, err := wc.Hug()
	if err != nil {
		t.Error(err)
	}
}

func TestNeko(t *testing.T) {
	_, err := wc.Neko()
	if err != nil {
		t.Error(err)
	}
}

func TestKiss(t *testing.T) {
	_, err := wc.Slap()
	if err == nil {
		t.Error(err)
	}
}

func TestSlap(t *testing.T) {
	_, err := wc.Slap()
	if err != nil {
		t.Error(err)
	}
}
