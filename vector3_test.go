package linalg

import "testing"

func TestVectorAdd(t *testing.T) {
	v := Vector3{4, 2, 1}
	a := Vector3{4, 3, 2}
	v.Add(a)
	if v.X != (4 + 4) {
		t.Errorf("v.Add(a): expected %d, bot got %f", 4+4, v.X)
	}
	if v.Y != (2 + 3) {
		t.Errorf("v.Add(a): expected %d, bot got %f", 2+3, v.Y)
	}
	if v.Z != (1 + 2) {
		t.Errorf("v.Add(a): expected %d, bot got %f", 1+2, v.Z)
	}
}
