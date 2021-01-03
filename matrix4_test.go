package linalg

import (
	"testing"
)

func TestMultVector3WriteBack(t *testing.T) {
	v := Vector3{1, 3, 7}
	m := NewMatrix4()
	m.MultVector3WriteBack(v, 1)
	if v.X != 1 { // float32 got numeric error?!
		t.Errorf("m.MultVector3WriteBack(v): expected 1, bot got %f", v.X)
	}
}

func TestMultVector3(t *testing.T) {
	v := Vector3{1, 3, 7}
	m := NewMatrix4()
	k := m.MultVector3(v, 1)
	if k.X != 1 { // float32 got numeric error?!
		t.Errorf("m.MultVector3WriteBack(v): expected 1, bot got %f", k.X)
	}
}
