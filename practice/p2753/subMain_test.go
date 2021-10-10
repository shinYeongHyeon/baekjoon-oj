package p2753

import (
	"testing"
)

func Test2012IsLeap(t *testing.T) {
	isLeap := CalcLeap(2012)

	if isLeap != 1 {
		t.Error("2012 should be leap")
	}
}

func Test1900IsNotLeap(t *testing.T) {
	isLeap := CalcLeap(1900)

	if isLeap != 0 {
		t.Error("1900 should not be leap")
	}
}

func Test2000IsLeap(t *testing.T) {
	isLeap := CalcLeap(2000)

	if isLeap != 1 {
		t.Error("2000 should be leap")
	}
}