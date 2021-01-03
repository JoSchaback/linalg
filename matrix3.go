package linalg

// Matrix4 is a float32 4x4 matrix.
type Matrix3 struct {
	M_0_0, M_1_0, M_2_0 float32
	M_0_1, M_1_1, M_2_1 float32
	M_0_2, M_1_2, M_2_2 float32
}

// NewMatrix4 creates a brand new Matrix4 struct, set to identity.
func NewMatrix3() Matrix3 {
	m := Matrix3{}
	m.Eye()
	return m
}

func (m *Matrix3) Eye() {
	m.M_0_0 = 1
	m.M_0_1 = 0
	m.M_0_2 = 0
	m.M_1_0 = 0
	m.M_1_1 = 1
	m.M_1_2 = 0
	m.M_2_0 = 0
	m.M_2_1 = 0
	m.M_2_2 = 1
}

func (m *Matrix3) Invert() {
	det := m.Determinant()
	m.Transpose()
	m.M_0_0 /= det
	m.M_0_1 /= det
	m.M_0_1 /= det
	m.M_1_0 /= det
	m.M_1_1 /= det
	m.M_1_2 /= det
	m.M_2_0 /= det
	m.M_2_1 /= det
	m.M_2_2 /= det
}

func (m *Matrix3) Determinant() float32 {
	// http://en.wikipedia.org/wiki/Invertible_matrix
	a := m.M_0_0
	b := m.M_1_0
	c := m.M_2_0
	d := m.M_0_1
	e := m.M_1_1
	f := m.M_2_1
	g := m.M_0_2
	h := m.M_1_2
	k := m.M_2_2

	return a*(e*k-f*h) - b*(k*d-f*g) + c*(d*h-e*g)
}

func (m *Matrix3) Transpose() {
	tmp := float32(0)

	tmp = m.M_0_1
	m.M_0_1 = m.M_1_0
	m.M_1_0 = tmp
	tmp = m.M_0_2
	m.M_0_2 = m.M_2_0
	m.M_2_0 = tmp

	tmp = m.M_1_2
	m.M_1_2 = m.M_2_1
	m.M_2_1 = tmp
}

func (m *Matrix3) MakeNormalMatrix(view Matrix4) {
	// set upper left
	m.M_0_0 = view.M_0_0
	m.M_1_0 = view.M_1_0
	m.M_2_0 = view.M_2_0

	m.M_0_1 = view.M_0_1
	m.M_1_1 = view.M_1_1
	m.M_2_1 = view.M_2_1

	m.M_0_2 = view.M_0_2
	m.M_1_2 = view.M_1_2
	m.M_2_2 = view.M_2_2

	m.Invert()
	m.Transpose()
}

func (m *Matrix3) multVector3(v Vector3) Vector3 {
	nx := v.X*m.M_0_0 + v.Y*m.M_1_0 + v.Z*m.M_2_0
	ny := v.X*m.M_0_1 + v.Y*m.M_1_1 + v.Z*m.M_2_1
	nz := v.X*m.M_0_2 + v.Y*m.M_1_2 + v.Z*m.M_2_2

	return Vector3{nx, ny, nz}
}

func (m *Matrix3) multVector3WriteBack(v Vector3) {
	nx := v.X*m.M_0_0 + v.Y*m.M_1_0 + v.Z*m.M_2_0
	ny := v.X*m.M_0_1 + v.Y*m.M_1_1 + v.Z*m.M_2_1
	nz := v.X*m.M_0_2 + v.Y*m.M_1_2 + v.Z*m.M_2_2

	v.X = nx
	v.Y = ny
	v.Z = nz
}
