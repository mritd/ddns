package main

import "testing"

func TestRun(t *testing.T) {
	initConf()
	err := run()
	if err != nil {
		t.Fatal(err)
	}
}
