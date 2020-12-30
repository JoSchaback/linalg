# linalg: Linear Algebra in Go for 3D graphics
This is a small and self-contained linear algebra library
for 3D graphics written in Go.

The purpose of this lib is to avoid arrays for data storage, avoid loops in favor of
explicitly unfolded liner algebra operations and avoid unecessary abstractions (e.g. interfaces) that
are often used in linear algebra libs. The intention is to support look-ahead cache utilization, 
memory handling and compiler optimization. The internal structure of the types follows the WebGL / OpenGL format to allow for fast and easy uploading (e.g. no additional conversions or transformations) to GPU shaders.

The library currently features the following types:
- **Matrix4**: A `float32` 4x4 matrix as commonly used for projection matrices or model-view matrics with WebGL or OpenGL.
- **Vector3**: A `float32`, 3-dimensional vector (x, y, z), the working horse of every 3d application.

## Usage
To include `linalg` in your project, you may import this
project into your module or source files.

```Go
import "github.com/schabby/linalg"
```

Depending of your dependency management, you may need to run 
```bash
go get github.com/schabby/linalg
```
to download the library to your local Go environment.

The package name is `linalg`. Everyhting is in that package.

At the initilization time of your program
you create your matrics. In this hypotentical example, we
create to matrices, one called `viewMatrix` that may be used for the transformation of world space coordinates into camera space coordinates and `projectionMatrix` to project the camera space onto the 2D screen canvas.
```Go
var viewMatrix = linalg.NewMatrix4()
var projectionMatrix = linalg.NewMatrix4()
```
To avoid unnecessary garbage collection runs, it is recommended to initialize the matrices during the initialization phase of your program and reuse them throughout the lifetime of your program.

To set up a perspective projection matrix, you can do
```Go
// couple of helper vars, just for clarification
viewAngle := 45 // field of view angle
width := 1024 // view port width, used for aspect ratio
height := 768 // view port height, used for aspect ratio
nearClippingPlaceDistance := 0.1
farClippingPlaceDistance := 100

// sets projectionMatrix to map from camera space to 2D canvas space
projectionMatrix.Projection(
    viewAngle, 
    width, 
    height, 
    nearClippingPlaceDistance, 
    farClippingPlaceDistance)
```
and upload it to your shader instance.

As for the model-view matrix, there are usually two or more steps involved to combine the "model" matrix and the "view"
matrix. A common pattern in simple shaders is to use separate model matrixes
to move objects / vertices around in world space and a single view matrix to map from world space into camera space.

```Go
eye := linalg.Vector3{2, 2, 3} // camera position
center := linalg.Vector3{0, 0, 0} // center where camera "looks at"
up := linalg.Vector3{0, 0, 1} // up vector of camera

// compute mapping for world- to camera space in viewMatrix.
viewMatrix.LookAt(eye, center, up) 

// set up a rotation matrix (around the z axis) in world space
rot := linalg.NewMatrix4()

// set up rotation as a quarter of Pi around z axis
rot.Rotation(math.Pi/4, linalg.Vector3{0, 0, 1})

// multiply rotation matrix with view matrix into a
// modelViewMatrix
viewMatrix.MultAssign(&rot)
```
Then, upload the modelViewMatrix to your shader, e.g. via `UniformMatrix4fv` or similarly. Note that the scope of the helper types `eye`, `center` and `up` etc. live only in the function scope such that the compiler can allocate the
space on the stack with minimal performance costs (no GC).