package test

import "testing"

func TestNewTest(t *testing.T) {
	tmp := NewTest()
	println(tmp.num)
}
