package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var a = 10
	var b = 20
	t.Logf("a = %d b = %d\n", a, b)
	c := Add(a, b)
	if c != 30 {
		t.Fatalf("invalid a + b, c=%d", c)
	}
	t.Logf("a = %d b = %d sum=%d\n", a, b, c)
}

func TestSub(t *testing.T) {
	var a = 10
	var b = 20
	t.Logf("a = %d b = %d\n", a, b)
	c := Sub(a, b)
	if c != -10 {
		t.Fatalf("invalid a - b, c=%d", c)
	}
	t.Logf("a = %d b = %d sub=%d\n", a, b, c)
}
