package spheregraph_test

import (
    "spheregraph"
    "testing"
)

func TestFaceReferencesAtLevel2(t *testing.T) {
    spheregraph.Init()
    level_1 := spheregraph.Icosahedron()
    level_2 := spheregraph.GenerateSet(level_1)

    var indexList [80]int64

    for _, face := range level_2 {
        indexList[face.GetA()] += 1
        indexList[face.GetB()] += 1
        indexList[face.GetC()] += 1
    }

    for _, count := range indexList {
        if count != 3 {
            t.Errorf("Test expected each face to be referenced 3 times, got: %d", count)
        }
    }
}

func TestFaceReferencesAtLevel3(t *testing.T) {
    spheregraph.Init()
    level_3 := spheregraph.GenerateSet(spheregraph.GenerateSet(spheregraph.Icosahedron()))

    var indexList [320]int64

    for _, face := range level_3 {
        indexList[face.GetA()] += 1
        indexList[face.GetB()] += 1
        indexList[face.GetC()] += 1
    }

    for _, count := range indexList {
        if count != 3 {
            t.Errorf("Test expected each face to be referenced 3 times, got: %d", count)
        }
    }
}

func TestFaceReferencesAtLevel4(t *testing.T) {
    spheregraph.Init()
    level_4 := spheregraph.GenerateSet(spheregraph.GenerateSet(spheregraph.GenerateSet(spheregraph.Icosahedron())))

    var indexList [1280]int64

    for _, face := range level_4 {
        indexList[face.GetA()] += 1
        indexList[face.GetB()] += 1
        indexList[face.GetC()] += 1
    }

    for _, count := range indexList {
        if count != 3 {
            t.Errorf("Test expected each face to be referenced 3 times, got: %d", count)
        }
    }
}