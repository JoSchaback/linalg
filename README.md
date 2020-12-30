# linalg: Linear Algebra in Go for 3D graphics
This is a small and self-contained linear algebra library
for 3D graphics written in Go.

The purpose of this lib is to avoid arrays for data storage, avoid loops in favor of
explicitly unfolded liner algebra operations and avoid unecessary abstractions (e.g. interfaces) that
are often used in linear algebra libs. The intention is to support look-ahead cache utilization, 
memory handling and compiler optimization. 

The library features the following types:
- *Matrix4*: A `float32` 4x4 matrix as commonly used for projection matrices or model-view matrics.
- *Vector3*: A `float32`, 3-dimensional vector (x, y, z), the working horse of every 3d application.


