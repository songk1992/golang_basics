package lec002

import "testing"

func TestLearnNumber(t *testing.T) {
	var expected int32 = -2147475648 
	if ret := LearnNumber(); ret != expected {
		t.Errorf("LearnNumber() = %d, want %d", ret, expected)
	}
}

func TestLearnString(t *testing.T) {
	var expected string = "code master1 code master2  code master4"
	if ret := LearnString(); ret != expected {
		t.Errorf("LearnString() = %q, want %q", ret, expected)
	}
}

