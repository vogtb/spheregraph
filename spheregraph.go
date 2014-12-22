// A golang package to create an icosahedron and recusively subdivide the faces to generate a graph of a sphere.
package spheregraph

type Face struct {
    parentIndex     int64 // index of parent face
    index           int64 // index of this face
    a               int64 // neighbor a
    b               int64 // neighbor b
    c               int64 // neighbor c
    done            bool
    children        []Face
}

var icosahedron []Face
var count int64

func Icosahedron() []Face {
    icosahedron := []Face {
        Face {index: 0, a: 4, b: 1, c: 5, done: false,},
        Face {index: 1, a: 7, b: 0, c: 2, done: false,},
        Face {index: 2, a: 9, b: 3, c: 1, done: false,},
        Face {index: 3, a: 11, b: 2, c: 4, done: false,},
        Face {index: 4, a: 0, b: 13, c: 3, done: false,},
        Face {index: 5, a: 14, b: 6, c: 0, done: false,},
        Face {index: 6, a: 15, b: 5, c: 7, done: false,},
        Face {index: 7, a: 1, b: 8, c: 6, done: false,},
        Face {index: 8, a: 16, b: 7, c: 9, done: false,},
        Face {index: 9, a: 2, b: 10, c: 8, done: false,},
        Face {index: 10, a: 17, b: 9, c: 11, done: false,},
        Face {index: 11, a: 3, b: 12, c: 10, done: false,},
        Face {index: 12, a: 13, b: 11, c: 18, done: false,},
        Face {index: 13, a: 12, b: 4, c: 14, done: false,},
        Face {index: 14, a: 5, b: 19, c: 13, done: false,},
        Face {index: 15, a: 6, b: 16, c: 19, done: false,},
        Face {index: 16, a: 8, b: 15, c: 17, done: false,},
        Face {index: 17, a: 10, b: 18, c: 16, done: false,},
        Face {index: 18, a: 19, b: 17, c: 12, done: false,},
        Face {index: 19, a: 18, b: 14, c: 15, done: false,},
    }
    return icosahedron
}

func GenerateSet(parentFaceSet []Face) []Face {
    var fullListOfFaces []Face
    count = 0

    // Subdivide face, and assign as children
    for i := 0; i < len(parentFaceSet); i++ {
        parentFaceSet[i].children = subdivideFace(parentFaceSet[i])
        count += 4

        //If neighbor A is done, connect it, etc.
        if parentFaceSet[parentFaceSet[i].a].done {
            connectByRule(&parentFaceSet[i], &parentFaceSet[parentFaceSet[i].a], "A")
        }
        if parentFaceSet[parentFaceSet[i].b].done {
            connectByRule(&parentFaceSet[i], &parentFaceSet[parentFaceSet[i].b], "B")
        }
        if parentFaceSet[parentFaceSet[i].c].done {
            connectByRule(&parentFaceSet[i], &parentFaceSet[parentFaceSet[i].c], "C")
        }
        parentFaceSet[i].done = true
    }

    //For each parent faces
    for i := 0; i < len(parentFaceSet); i++ {
        for x := 0; x < len(parentFaceSet[i].children); x++ {
            fullListOfFaces = append(fullListOfFaces, parentFaceSet[i].children[x])
        }
    }
    return fullListOfFaces
}

//When given two faces, will connect by the given rule (A, B, or C)
func connectByRule(faceOne *Face, faceTwo *Face, rule string) {
    switch rule {
        case "A":
            // 1a1
            faceOne.children[1].a = faceTwo.children[1].index
            faceTwo.children[1].a = faceOne.children[1].index
            // 3a3
            faceOne.children[3].a = faceTwo.children[3].index
            faceTwo.children[3].a = faceOne.children[3].index
        case "B":
            // 2b2
            faceOne.children[2].b = faceTwo.children[2].index
            faceTwo.children[2].b = faceOne.children[2].index
            // 3b3
            faceOne.children[3].b = faceTwo.children[3].index
            faceTwo.children[3].b = faceOne.children[3].index
        case "C":
            // 2c2
            faceOne.children[2].c = faceTwo.children[2].index
            faceTwo.children[2].c = faceOne.children[2].index
            // 1c1
            faceOne.children[1].c = faceTwo.children[1].index
            faceTwo.children[1].c = faceOne.children[1].index
    }
}

// Will return 4 faces that makeup a given face
func subdivideFace(face Face) []Face {
    parent := face.index
    set := []Face {
        Face {parentIndex: parent, index:count, a: count + 2, b: count + 1, c: count + 3, done: false,}, // sub 0, always neighbors by a:2, b:1, c:3 relative to siblings
        Face {parentIndex: parent, index:count + 1, a: 0, b: count, c: 0, done: false,}, // sub 1, neighbor by b
        Face {parentIndex: parent, index:count + 2, a: count, b: 0, c: 0, done: false,}, // sub 2, neighbor by a
        Face {parentIndex: parent, index:count + 3, a: 0, b: 0, c: count, done: false,}, // sub 3, neighbor by c
    }
    return set
}