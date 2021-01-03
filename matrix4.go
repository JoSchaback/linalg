package linalg

import (
	"math"
)

// Matrix4 is a float32 4x4 matrix.
type Matrix4 struct {
	M_0_0, M_1_0, M_2_0, M_3_0 float32
	M_0_1, M_1_1, M_2_1, M_3_1 float32
	M_0_2, M_1_2, M_2_2, M_3_2 float32
	M_0_3, M_1_3, M_2_3, M_3_3 float32
}

// NewMatrix4 creates a brand new Matrix4 struct, set to identity.
func NewMatrix4() Matrix4 {
	m := Matrix4{}
	m.Eye()
	return m
}

func (self *Matrix4) Row(row int, x float32, y float32, z float32, w float32) {
	switch row {
	case 0:
		self.M_0_0 = x
		self.M_1_0 = y
		self.M_2_0 = z
		self.M_3_0 = w
		break
	case 1:
		self.M_0_1 = x
		self.M_1_1 = y
		self.M_2_1 = z
		self.M_3_1 = w
		break
	case 2:
		self.M_0_2 = x
		self.M_1_2 = y
		self.M_2_2 = z
		self.M_3_2 = w
		break
	case 3:
		self.M_0_3 = x
		self.M_1_3 = y
		self.M_2_3 = z
		self.M_3_3 = w
		break
	}
}

func (m *Matrix4) Eye() {
	m.M_0_0 = 1
	m.M_0_1 = 0
	m.M_0_2 = 0
	m.M_0_3 = 0
	m.M_1_0 = 0
	m.M_1_1 = 1
	m.M_1_2 = 0
	m.M_1_3 = 0
	m.M_2_0 = 0
	m.M_2_1 = 0
	m.M_2_2 = 1
	m.M_2_3 = 0
	m.M_3_0 = 0
	m.M_3_1 = 0
	m.M_3_2 = 0
	m.M_3_3 = 1
}

func (self *Matrix4) Frustum(left float32, right float32, bottom float32, top float32, near float32, far float32) {
	// http://www.songho.ca/opengl/gl_projectionmatrix.html
	self.Eye()

	self.M_0_0 = (2 * near) / (right - left)
	self.M_2_0 = (right + left) / (right - left)

	self.M_1_1 = (2 * near) / (top - bottom)
	self.M_2_1 = (top + bottom) / (top - bottom)

	self.M_2_2 = -(far + near) / (far - near)
	self.M_3_2 = -2 * (far * near) / (far - near)

	self.M_2_3 = -1
	self.M_3_3 = 0
}

const Pi32 = float32(math.Pi)

func (self *Matrix4) Projection(viewAngle float32, width float32, height float32, nearClippingPlaneDistance float32, farClippingPlaneDistance float32) {
	// http://www.geeks3d.com/20090729/howto-perspective-projection-matrix-in-opengl/
	radians := (viewAngle * Pi32 / 180)
	halfHeight := float32(math.Tan(float64(radians/2))) * nearClippingPlaneDistance
	halfScaledAspectRatio := halfHeight * (width / height)
	self.Frustum(-halfScaledAspectRatio, halfScaledAspectRatio, -halfHeight, halfHeight, nearClippingPlaneDistance, farClippingPlaneDistance)
}

func (self *Matrix4) MultAssign(that *Matrix4) {
	m00 := self.M_0_0*that.M_0_0 + self.M_1_0*that.M_0_1 + self.M_2_0*that.M_0_2 + self.M_3_0*that.M_0_3
	m01 := self.M_0_1*that.M_0_0 + self.M_1_1*that.M_0_1 + self.M_2_1*that.M_0_2 + self.M_3_1*that.M_0_3
	m02 := self.M_0_2*that.M_0_0 + self.M_1_2*that.M_0_1 + self.M_2_2*that.M_0_2 + self.M_3_2*that.M_0_3
	m03 := self.M_0_3*that.M_0_0 + self.M_1_3*that.M_0_1 + self.M_2_3*that.M_0_2 + self.M_3_3*that.M_0_3

	m10 := self.M_0_0*that.M_1_0 + self.M_1_0*that.M_1_1 + self.M_2_0*that.M_1_2 + self.M_3_0*that.M_1_3
	m11 := self.M_0_1*that.M_1_0 + self.M_1_1*that.M_1_1 + self.M_2_1*that.M_1_2 + self.M_3_1*that.M_1_3
	m12 := self.M_0_2*that.M_1_0 + self.M_1_2*that.M_1_1 + self.M_2_2*that.M_1_2 + self.M_3_2*that.M_1_3
	m13 := self.M_0_3*that.M_1_0 + self.M_1_3*that.M_1_1 + self.M_2_3*that.M_1_2 + self.M_3_3*that.M_1_3

	m20 := self.M_0_0*that.M_2_0 + self.M_1_0*that.M_2_1 + self.M_2_0*that.M_2_2 + self.M_3_0*that.M_2_3
	m21 := self.M_0_1*that.M_2_0 + self.M_1_1*that.M_2_1 + self.M_2_1*that.M_2_2 + self.M_3_1*that.M_2_3
	m22 := self.M_0_2*that.M_2_0 + self.M_1_2*that.M_2_1 + self.M_2_2*that.M_2_2 + self.M_3_2*that.M_2_3
	m23 := self.M_0_3*that.M_2_0 + self.M_1_3*that.M_2_1 + self.M_2_3*that.M_2_2 + self.M_3_3*that.M_2_3

	m30 := self.M_0_0*that.M_3_0 + self.M_1_0*that.M_3_1 + self.M_2_0*that.M_3_2 + self.M_3_0*that.M_3_3
	m31 := self.M_0_1*that.M_3_0 + self.M_1_1*that.M_3_1 + self.M_2_1*that.M_3_2 + self.M_3_1*that.M_3_3
	m32 := self.M_0_2*that.M_3_0 + self.M_1_2*that.M_3_1 + self.M_2_2*that.M_3_2 + self.M_3_2*that.M_3_3
	m33 := self.M_0_3*that.M_3_0 + self.M_1_3*that.M_3_1 + self.M_2_3*that.M_3_2 + self.M_3_3*that.M_3_3

	self.M_0_0 = m00
	self.M_0_1 = m01
	self.M_0_2 = m02
	self.M_0_3 = m03

	self.M_1_0 = m10
	self.M_1_1 = m11
	self.M_1_2 = m12
	self.M_1_3 = m13

	self.M_2_0 = m20
	self.M_2_1 = m21
	self.M_2_2 = m22
	self.M_2_3 = m23

	self.M_3_0 = m30
	self.M_3_1 = m31
	self.M_3_2 = m32
	self.M_3_3 = m33
}

func (self *Matrix4) Copy(vec []float32) {
	vec[0] = self.M_0_0
	vec[1] = self.M_0_1
	vec[2] = self.M_0_2
	vec[3] = self.M_0_3
	vec[4] = self.M_1_0
	vec[5] = self.M_1_1
	vec[6] = self.M_1_2
	vec[7] = self.M_1_3
	vec[8] = self.M_2_0
	vec[9] = self.M_2_1
	vec[10] = self.M_2_2
	vec[11] = self.M_2_3
	vec[12] = self.M_3_0
	vec[13] = self.M_3_1
	vec[14] = self.M_3_2
	vec[15] = self.M_3_3
}

func (self *Matrix4) Rotation(alpha float32, u Vector3) {
	self.Eye()
	c := float32(math.Cos(float64(alpha)))
	s := float32(math.Sin(float64(alpha)))
	t := 1 - c

	self.M_0_0 = t*u.X*u.X + c
	self.M_1_0 = t*u.X*u.Y - u.Z*s
	self.M_2_0 = u.X*u.Z*t + u.Y*s
	self.M_3_0 = 0
	self.M_0_1 = t*u.Y*u.X + u.Z*s
	self.M_1_1 = t*u.Y*u.Y + c
	self.M_2_1 = u.Y*u.Z*t - u.X*s
	self.M_3_1 = 0
	self.M_0_2 = t*u.Z*u.X - u.Y*s
	self.M_1_2 = t*u.Z*u.Y + u.X*s
	self.M_2_2 = u.Z*u.Z*t + c
	self.M_3_2 = 0
	self.M_0_3 = 0
}

func (self *Matrix4) LookAt(eye Vector3, center Vector3, up Vector3) {

	u := Vector3{}
	v := Vector3{}

	// the w vector is computed by w = eye - center which means
	// it is the inverse of the viewing direction.
	w := Vector3{eye.X, eye.Y, eye.Z}
	w.Sub(center) // -= center
	w.Normalize()

	//dom.Log(fmt.Sprintf("w: %v, length %f", w, w.L2Norm()))

	// compute cross product
	u.Set(up)
	u.Cross(w)
	u.Normalize() // side = (0,0,1) x w

	//dom.Log(fmt.Sprintf("u: %v", u))

	// up = side x look
	v.Set(w)
	v.Cross(u)
	v.Normalize()

	//dom.Log(fmt.Sprintf("v: %v", v))

	// note the format: set(COLUMN, ROW, value)
	// it may be different for your matrix implementation
	self.Eye()

	self.Row(0, u.X, u.Y, u.Z, 0)
	self.Row(1, v.X, v.Y, v.Z, 0)
	self.Row(2, w.X, w.Y, w.Z, 0)

	translation := Matrix4{}
	translation.Eye()
	translation.M_3_0 = -eye.X
	translation.M_3_1 = -eye.Y
	translation.M_3_2 = -eye.Z

	self.MultAssign(&translation)
}

func (m *Matrix4) MultVector3(v Vector3, w float32) Vector3 {
	nx := v.X*m.M_0_0 + v.Y*m.M_1_0 + v.Z*m.M_2_0 + w*m.M_3_0
	ny := v.X*m.M_0_1 + v.Y*m.M_1_1 + v.Z*m.M_2_1 + w*m.M_3_1
	nz := v.X*m.M_0_2 + v.Y*m.M_1_2 + v.Z*m.M_2_2 + w*m.M_3_2
	//nw := x * m.M_0_3 + y * m.M_1_3 + z * m.M_2_3 + w * m.M_3_3

	return Vector3{nx, ny, nz}
}

func (m *Matrix4) MultVector3WriteBack(v Vector3, w float32) {
	nx := v.X*m.M_0_0 + v.Y*m.M_1_0 + v.Z*m.M_2_0 + w*m.M_3_0
	ny := v.X*m.M_0_1 + v.Y*m.M_1_1 + v.Z*m.M_2_1 + w*m.M_3_1
	nz := v.X*m.M_0_2 + v.Y*m.M_1_2 + v.Z*m.M_2_2 + w*m.M_3_2
	//nw := x * m.M_0_3 + y * m.M_1_3 + z * m.M_2_3 + w * m.M_3_3

	v.X = nx
	v.Y = ny
	v.Z = nz
	// dropping w
	// v.W= nw
}
