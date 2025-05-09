package main

import (
  "fmt"
)

func CompByteSlice(x *[]byte, x2 *[]byte) bool {
  n := len(*x) 
  if n != len(*x2) {
    return false
  }
  for i := 0; i < n; i++ {
    if (*x)[i] != (*x2)[i] {
      return false
    }
  }
  return true
}

func CompInt32Slice(x *[]int32, x2 *[]int32) bool {
  n := len(*x) 
  if n != len(*x2) {
    return false
  }
  for i := 0; i < n; i++ {
    if (*x)[i] != (*x2)[i] {
      return false
    }
  }
  return true
}

func CompSlices(x1 any, x2 any) bool {
  switch x1 := x1.(type) {
    case *[]byte:
      x2, is_same := x2.(*[]byte)
      if !is_same {
        return false
      }
      return CompByteSlice(x1, x2)
    case *[]int32:
      x2, is_same := x2.(*[]int32)
      if !is_same {
        return false
      }
      return CompInt32Slice(x1, x2)
    default:
      return false
  }
}

func main() {
  x1 := []int32{2, 1, 3}
  x2 := []int32{2, 1, 3}
  vl := CompSlices(&x1, &x2)
  fmt.Println(vl)
}


