package main

import (
  "spheregraph"
  "fmt"
  "encoding/json"
)

func main() {
  icosa := spheregraph.Icosahedron()
  level_2 := spheregraph.GenerateSet(icosa)

  level_2_json, err := json.Marshal(level_2)
  if err != nil {
      fmt.Println(err)
      return
  }
  fmt.Println(string(level_2_json))
}