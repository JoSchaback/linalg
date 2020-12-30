package linalg

import "math"

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func NewVector3() Vector3 {
	return Vector3{0, 0, 0}
}

func (self *Vector3) Add(v Vector3) {
	self.X += v.X
	self.Y += v.Y
	self.Z += v.Z
}

func (self *Vector3) Set(v Vector3) {
	self.X = v.X
	self.Y = v.Y
	self.Z = v.Z
}

func (self *Vector3) Sub(v Vector3) {
	self.X -= v.X
	self.Y -= v.Y
	self.Z -= v.Z
}

func (self *Vector3) Mult(v Vector3) {
	self.X *= v.X
	self.Y *= v.Y
	self.Z *= v.Z
}

func (self *Vector3) MultScalar(s float32) {
	self.X *= s
	self.Y *= s
	self.Z *= s
}

func (self *Vector3) Div(v Vector3) {
	self.X /= v.X
	self.Y /= v.Y
	self.Z /= v.Z
}

func (self *Vector3) DivScalar(s float32) {
	self.X /= s
	self.Y /= s
	self.Z /= s
}

func (self *Vector3) L2Norm() float32 {
	return float32(math.Sqrt(float64(self.X*self.X + self.Y*self.Y + self.Z*self.Z)))
}

func abs(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

func (self *Vector3) L1Norm() float32 {
	return abs(self.X) + abs(self.Y) + abs(self.Z)
}

// Dot computes the dot product between self and a.
func (self *Vector3) Dot(a Vector3) float32 {
	return self.X*a.X + self.Y*a.Y + self.Z*a.Z
}

func (self *Vector3) Normalize() {
	self.DivScalar(self.L2Norm())
}

func (self *Vector3) Cross(a Vector3) {
	tempX := self.Y*a.Z - self.Z*a.Y
	tempY := self.Z*a.X - self.X*a.Z
	tempZ := self.X*a.Y - self.Y*a.X

	self.X = tempX
	self.Y = tempY
	self.Z = tempZ
}
