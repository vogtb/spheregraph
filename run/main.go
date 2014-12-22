package main

import (
  "spheregraph"
  "fmt"
)

func main() {
  level_1 := spheregraph.Icosahedron()
  level_2 := spheregraph.GenerateSet(level_1)
  fmt.Println(len(level_2))
  level_3 := spheregraph.GenerateSet(level_2)
  fmt.Println(len(level_3))
  fmt.Println(level_3)
}