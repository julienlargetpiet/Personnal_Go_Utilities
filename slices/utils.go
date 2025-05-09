
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
