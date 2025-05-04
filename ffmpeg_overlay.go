package main

import (
  "fmt"
  "os"
  "os/exec"
)

func main() {
  args_v := os.Args
  var n int = len(os.Args)
  if n < 2 {
    fmt.Println("Missing files")
    return
  }
  if n < 3 {
    fmt.Println("Missing second file")
    return
  }
  if n < 4 {
    fmt.Println("Missing output file")
    return
  }
  base_file := args_v[1]
  ovrl_file := args_v[2]
  out_file := args_v[3]
  command_val := `ffmpeg -i ` + base_file + ` -i ` + ovrl_file + ` -filter_complex "[1:v]scale=690:430[ovrl];[0:v][ovrl]overlay=x=(main_w-overlay_w):y=570[outv]" -map "[outv]" -y ` + out_file
  fmt.Println("Running...")
  cmd := exec.Command("sh", "-c", command_val)
  _, err := cmd.Output()
  if err != nil {
    fmt.Println(err)
    return
  }
  return
}


