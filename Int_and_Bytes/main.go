package main

import (
  "fmt"
)

func IntToByteSlice(x int) []byte {
  if x == 256 {
    return []byte{0, 0}
  } else if x < 256 {
    return []byte{byte(x)}
  }
  var rtn_byte []byte
  var rest int = x % 256
  rtn_byte = append(rtn_byte, byte(rest))
  x -= rest
  x /= 256
  for x > 256 {
    rtn_byte = append(rtn_byte, 255)
    rest = x % 256
    rtn_byte = append(rtn_byte, byte(rest))
    x -= rest
    x /= 256
  }
  rtn_byte = append(rtn_byte, byte(x - 1))
  return rtn_byte
}

func ByteSliceToInt(x []byte) int {
  var rtn_int int = 256
  var ref_mult int = 256
  var i int = len(x) - 1
  if i == 0 {
    return int(x[0])
  }
  for i > -1 {
    rtn_int = ((int(x[i]) + 1) * ref_mult + int(x[i - 1]))
    ref_mult = rtn_int
    i -= 2
  }
  return rtn_int
}

func main() {
  var x int = 172564558
  x2 := IntToByteSlice(x)
  fmt.Println(x2)
  x = 65536
  x2 = IntToByteSlice(x)
  fmt.Println(x2)
  x = ByteSliceToInt(x2)
  fmt.Println(x)
  x2 = []byte{0, 255}
  x = ByteSliceToInt(x2)
  fmt.Println(x)
}



