// A golang package to create an icosahedron and recusively subdivide the faces to generate a graph of a sphere.
package spheregraph

import "fmt"

type Face struct {
  parentIndex   int64
  a             int64 // neighbor a
  b             int64 // neighbor b
  c             int64 // neighbor c
  done          bool
}

var icosahedron []Face
var worldOne []Face

func Icosahedron() []Face {
  icosahedron := []Face {
    Face {parentIndex: 0, a: 4, b: 1, c: 5, done: false,},
    Face {parentIndex: 1, a: 7, b: 0, c: 2, done: false,},
    Face {parentIndex: 2, a: 9, b: 3, c: 1, done: false,},
    Face {parentIndex: 3, a: 11, b: 2, c: 4, done: false,},
    Face {parentIndex: 4, a: 0, b: 13, c: 3, done: false,},
    Face {parentIndex: 5, a: 14, b: 6, c: 0, done: false,},
    Face {parentIndex: 6, a: 15, b: 5, c: 7, done: false,},
    Face {parentIndex: 7, a: 1, b: 8, c: 6, done: false,},
    Face {parentIndex: 8, a: 16, b: 7, c: 9, done: false,},
    Face {parentIndex: 90, a: 2, b: 10, c: 8, done: false,},
    Face {parentIndex: 10, a: 17, b: 9, c: 11, done: false,},
    Face {parentIndex: 11, a: 3, b: 12, c: 10, done: false,},
    Face {parentIndex: 12, a: 13, b: 11, c: 18, done: false,},
    Face {parentIndex: 13, a: 12, b: 4, c: 14, done: false,},
    Face {parentIndex: 14, a: 5, b: 19, c: 13, done: false,},
    Face {parentIndex: 15, a: 6, b: 16, c: 19, done: false,},
    Face {parentIndex: 16, a: 8, b: 15, c: 17, done: false,},
    Face {parentIndex: 17, a: 10, b: 18, c: 16, done: false,},
    Face {parentIndex: 18, a: 19, b: 17, c: 12, done: false,},
    Face {parentIndex: 19, a: 18, b: 14, c: 15, done: false,},
  }
  return icosahedron
}


func GenerateNextSet(icosa []Face) []Face {
  var newFaceSet []Face

  for i := 0; i < len(icosa); i++ {
    faceSetFromSubdivide := subdivideFace(icosa[i])
    newFaceSet = append(newFaceSet, faceSetFromSubdivide...)
  }

  fmt.Println("END")
  return newFaceSet
}

func ConnectSetFromParent(parentSet []Face, childSet []Face) []Face {
  var finalSet []Face

  // For each in the parent set, 

  return finalSet
}




// Will return 4 faces that makeup a given face
func subdivideFace(face Face) []Face {
  parent := face.parentIndex
  set := []Face {
    Face {parentIndex: parent, a: 0, b: 0, c: 0, done: false,}, // sub 1
    Face {parentIndex: parent, a: 0, b: 0, c: 0, done: false,}, // sub 2
    Face {parentIndex: parent, a: 0, b: 0, c: 0, done: false,}, // sub 3
    Face {parentIndex: parent, a: 0, b: 0, c: 0, done: false,}, // sub 4
  }
  return set
}








