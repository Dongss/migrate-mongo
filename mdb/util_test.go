package mdb

import (
	"testing"
)

func TestContains(t *testing.T) {
	str := []string{"a", "b", "c"}
	if r := contains(str, "a"); r != true {
		t.Fatal("should contaian")
	}
	if r := contains(str, "d"); r != false {
		t.Fatal("should not contaian")
	}
}

func TestFind(t *testing.T) {
	str := []string{"a", "b", "c"}
	if r := find(str, "a"); r != 0 {
		t.Fatal("should find index 0")
	}
	if r := find(str, "d"); r != -1 {
		t.Fatal("should not find")
	}
}
