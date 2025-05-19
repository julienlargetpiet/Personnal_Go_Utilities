package main

import (
  "fmt"
)

func IntToByteSlice(x int) []byte {
  var rtn_byte []byte
  var rest int = x % 256
  rtn_byte = append(rtn_byte, byte(rest))
  fmt.Println("rest:", rest)
  x -= rest
  fmt.Println(x)
  x /= 256
  fmt.Println(x)
  for x > 256 {
    rtn_byte = append(rtn_byte, 255)
    rest = x % 256
    rtn_byte = append(rtn_byte, byte(rest))
    x -= rest
    x /= 256
    fmt.Println(x)
  }
  rtn_byte = append(rtn_byte, byte(x - 1))
  return rtn_byte
}



func main() {
  var x int = 102564558
  x2 := IntToByteSlice(x)
  fmt.Println(x2)
}
