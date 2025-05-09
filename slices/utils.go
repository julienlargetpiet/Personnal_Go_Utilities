package main

import (
  "fmt"
)

func CompByteSlice(x *[]byte, x2 *[]byte) bool { //also uint8
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

func CompInt8Slice(x *[]int8, x2 *[]int8) bool {
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

func CompInt16Slice(x *[]int16, x2 *[]int16) bool {
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

func CompInt64Slice(x *[]int64, x2 *[]int64) bool {
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

func CompUInt16Slice(x *[]uint16, x2 *[]uint16) bool {
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

func CompUInt32Slice(x *[]uint32, x2 *[]uint32) bool {
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

func CompUInt64Slice(x *[]uint64, x2 *[]uint64) bool {
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

func CompFloat32Slice(x *[]float32, x2 *[]float32) bool {
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

func CompFloat64Slice(x *[]float64, x2 *[]float64) bool {
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

func CompStringSlice(x *[]string, x2 *[]string) bool {
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

func CompBoolSlice(x *[]bool, x2 *[]bool) bool {
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
    case *[]int8:
      x2, is_same := x2.(*[]int8)
      if !is_same {
        return false
      }
      return CompInt8Slice(x1, x2)
    case *[]int16:
      x2, is_same := x2.(*[]int16)
      if !is_same {
        return false
      }
      return CompInt16Slice(x1, x2)
    case *[]int32:
      x2, is_same := x2.(*[]int32)
      if !is_same {
        return false
      }
      return CompInt32Slice(x1, x2)
    case *[]int64:
      x2, is_same := x2.(*[]int64)
      if !is_same {
        return false
      }
      return CompInt64Slice(x1, x2)
    case *[]uint16:
      x2, is_same := x2.(*[]uint16)
      if !is_same {
        return false
      }
      return CompUInt16Slice(x1, x2)
    case *[]uint32:
      x2, is_same := x2.(*[]uint32)
      if !is_same {
        return false
      }
      return CompUInt32Slice(x1, x2)
    case *[]uint64:
      x2, is_same := x2.(*[]uint64)
      if !is_same {
        return false
      }
      return CompUInt64Slice(x1, x2)
    case *[]float32:
      x2, is_same := x2.(*[]float32)
      if !is_same {
        return false
      }
      return CompFloat32Slice(x1, x2)
    case *[]float64:
      x2, is_same := x2.(*[]float64)
      if !is_same {
        return false
      }
      return CompFloat64Slice(x1, x2)
    case *[]string:
      x2, is_same := x2.(*[]string)
      if !is_same {
        return false
      }
      return CompStringSlice(x1, x2)
    case *[]bool:
      x2, is_same := x2.(*[]bool)
      if !is_same {
        return false
      }
      return CompBoolSlice(x1, x2)
    default:
      return false
  }
}

func main() {
  x1 := []int8{2, 1, 3}
  x2 := []int8{2, 1, 3}
  vl := CompSlices(&x1, &x2)
  fmt.Println(vl)
}


