import (
  "time"
)

var ref_ltr = [52]uint8{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var ref_nb = [10]uint8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var ref_spechr = [24]uint8{'!', '.', ':', ';', '\\', '-', '%', '*', ',', '_', '/', '<', '>', '=', '[', ']', '\'', '{', '}', '[', ']', '(', ')', '"'}

func GenerateAES256() string {
  rtn_str := ""
  var tm int64 = time.Now().Unix()
  for i := 0; i < 32; i++ {
    if tm % 3 == 0 {
      rtn_str += string(ref_ltr[tm % 52])
      tm *= 2
    } else if tm % 2 == 0 {
      rtn_str += string(ref_nb[tm % 10])
      tm /= 2
    } else {
      rtn_str += string(ref_spechr[tm % 24])
      tm *= 3
    }
  }
  return rtn_str
}
