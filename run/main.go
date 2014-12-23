package main

import (
    "spheregraph"
    "fmt"
)

func main() {
    spheregraph.Init()
    level_1 := spheregraph.Icosahedron()
    level_2 := spheregraph.GenerateSet(level_1)
    fmt.Println(len(level_2))
    level_3 := spheregraph.GenerateSet(level_2)
    fmt.Println(len(level_3))
    level_4 := spheregraph.GenerateSet(level_3)
    fmt.Println(len(level_4))
    level_5 := spheregraph.GenerateSet(level_4)
    fmt.Println(len(level_5))
    level_6 := spheregraph.GenerateSet(level_5)
    fmt.Println(len(level_6))
    level_7 := spheregraph.GenerateSet(level_6)
    fmt.Println(len(level_7))
    level_8 := spheregraph.GenerateSet(level_7)
    fmt.Println(len(level_8))
    level_9 := spheregraph.GenerateSet(level_8)
    fmt.Println(len(level_9))
    level_10 := spheregraph.GenerateSet(level_9)
    fmt.Println(len(level_10))
    // Hit a limit here where parallelization would be very useful.
}