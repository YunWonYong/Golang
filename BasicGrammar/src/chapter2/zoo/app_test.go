package main

import "testing"

func TestAppName(t *testing.T) {
	expect := "Zoo Application"
	actual := AppName()

	if expect != actual {
		t.Errorf("expect: %s, actual: %s", expect, actual)
	}
}
