package main

import (
  "spheregraph"
  "fmt"
)

func main() {
  icosa := spheregraph.Icosahedron()
  level_2 := spheregraph.SubdivideSet(icosa)




  fmt.Println(level_2)
}