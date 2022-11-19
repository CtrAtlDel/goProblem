package main

import (
	"testing"
)

func TestAddition(t *testting.T) {
	if Addition(2, 2) != 4 {
		t.Error("Some problems...")
	}
}
