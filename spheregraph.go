// A golang package to create an icosahedron and recusively subdivide the faces to generate a graph of a sphere.
package spheregraph

import "fmt"

type Face struct {
  index   int64
  a       int64 // neighbor a
  b       int64 // neighbor b
  c       int64 // neighbor c
  done    bool
}

var icosahedron []Face
var worldOne []Face

func Icosahedron() []Face {
  icosahedron := []Face {
    Face {index: 0, a: 1, b: 4, c: 5, done: false,},
    Face {index: 1, a: 2, b: 0, c: 7, done: false,},
    Face {index: 2, a: 1, b: 9, c: 3, done: false,},
    Face {index: 3, a: 2, b: 11, c: 4, done: false,},
    Face {index: 4, a: 2, b: 13, c: 0, done: false,},
    Face {index: 5, a: 0, b: 6, c: 14, done: false,},
    Face {index: 6, a: 5, b: 7, c: 15, done: false,},
    Face {index: 7, a: 6, b: 1, c: 8, done: false,},
    Face {index: 8, a: 7, b: 9, c: 16, done: false,},
    Face {index: 9, a: 2, b: 8, c: 10, done: false,},
    Face {index: 10, a: 9, b: 11, c: 17, done: false,},
    Face {index: 11, a: 3, b: 10, c: 12, done: false,},
    Face {index: 12, a: 11, b: 13, c: 18, done: false,},
    Face {index: 13, a: 4, b: 12, c: 14, done: false,},
    Face {index: 14, a: 13, b: 5, c: 19, done: false,},
    Face {index: 15, a: 19, b: 6, c: 16, done: false,},
    Face {index: 16, a: 15, b: 8, c: 17, done: false,},
    Face {index: 17, a: 10, b: 16, c: 18, done: false,},
    Face {index: 18, a: 12, b: 17, c: 19, done: false,},
    Face {index: 19, a: 18, b: 14, c: 15, done: false,},
  }
  return icosahedron
}

func SubdivideSet(icosa []Face) []Face {
  var newFaceSet []Face
  for _, e := range icosa {
    faceSet := subdivideFace(e)
    newFaceSet = append(newFaceSet, faceSet...)
  }
  fmt.Println(len(newFaceSet))
  return newFaceSet
}

// Will return 4 faces that makeup a given face
func subdivideFace(face Face) []Face {
  set := []Face {
    Face {index: 0, a: 0, b: 0, c: 0, done: false,},
    Face {index: 0, a: 0, b: 0, c: 0, done: false,},
    Face {index: 0, a: 0, b: 0, c: 0, done: false,},
    Face {index: 0, a: 0, b: 0, c: 0, done: false,},
  }
  return set
}








