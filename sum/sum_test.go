package sum_test

import (
	"test/sum"
	"testing"
)

func TestMain(t *testing.T) {
	num := sum.Sum(5, 10)
	if num != 15 {
		t.Error("TEST FAIL")
	}
}

func TestMulti(t *testing.T) {
	num := sum.Multi(5, 10)
	if num != 50 {
		t.Error("TEST FAIL")
	}
}
