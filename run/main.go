package main

import (
  "spheregraph"
  "fmt"
)

func main() {
  icosa := spheregraph.Icosahedron()
  level_2 := spheregraph.GenerateNextSet(icosa)
  level_2 = spheregraph.ConnectSetFromParent(icosa, level_2)




  fmt.Println(level_2)
}