package lec001

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello world"
	if ret := Lec001(); ret != expected {
		t.Errorf("Lec001() = %q, want %q", ret, expected)
	}
}