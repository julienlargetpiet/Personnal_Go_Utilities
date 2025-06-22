package main

import (
  "fmt"
  "os"
  "os/exec"
)

var first_file_audio bool = false
var audio_codec string = "aac"
var video_codec string = "mp4"
var command_val string
var width string = "854"
var height string = "480"

func IntToString(x int) string {
  const base int = 10
  var remainder int
  rtn_str := ""
  if x == 0 {
    return "0"
  }
  for x > 0 {
    remainder = x % base
    rtn_str = string(remainder + 48) + rtn_str
    x -= remainder
    x /= 10
  }
  return rtn_str
}

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
  base_file := args_v[1]
  ovrl_file := args_v[2]
  n = len(ovrl_file)
  var base_racine string
  var ovrl_racine string
  var ovrl_file_v []string
  var base_file_v []string
  var err error
  var str_i string
  var cmd *exec.Cmd
  all_files := ""
  cur_val := ""

  i := n - 1
  for ovrl_file[i] != '/' {
    i--
    if i == 0 {
      break
    }
  }
  i++
  ovrl_racine = ovrl_file[0:i]
  if ovrl_file[i] == '[' {
    i++
  }
  for i < n && ovrl_file[i] != ']' {
    if ovrl_file[i] != ',' {
      cur_val += string(ovrl_file[i])
    } else {
      if cur_val == "" {
        fmt.Println("Bad file name in overlay videos")
        return
      }
      ovrl_file_v = append(ovrl_file_v, "'" + cur_val + "'")
      cur_val = ""
    }
    i++
  }
  if cur_val == "" {
    fmt.Println("Bad file name in overlay videos")
    return
  }
  ovrl_file_v = append(ovrl_file_v, "'" + cur_val + "'")
  cur_val = ""
  n = len(base_file)
  i = n - 1
  for base_file[i] != '/' {
    i--
    if i == 0 {
      break
    }
  }
  i++
  base_racine = base_file[0:i]
  if base_file[i] == '[' {
    i++
  }
  for i < n && base_file[i] != ']' {
    if base_file[i] != ',' {
      cur_val += string(base_file[i])
    } else {
      if cur_val == "" {
        fmt.Println("Bad file name in background videos")
        return
      }
      base_file_v = append(base_file_v, "'" + cur_val + "'")
      cur_val = ""
    }
    i++
  }
  if cur_val == "" {
    fmt.Println("Bad file name in background videos")
    return
  }
  base_file_v = append(base_file_v, "'" + cur_val + "'")

  if len(base_file_v) != len(ovrl_file_v) {
    fmt.Println("The number of files are not equal")
    return
  }
  n = len(base_file_v)

  _, err = os.Stat("out." + video_codec)
  if err == nil {
    fmt.Println("Removing out." + video_codec + "...")
    err = os.Remove("out." + video_codec)
    if err != nil {
      fmt.Println(err)
      return
    }
  } else if !os.IsNotExist(err) {
    fmt.Println(err)
    return
  }

  for i = 0; i < n; i++ {
    _, err = os.Stat("audio." + audio_codec)
    if err == nil {
      fmt.Println("Removing audio." + audio_codec + "...")
      err = os.Remove("audio." + audio_codec)
      if err != nil {
        fmt.Println(err)
        return
      }
    } else if !os.IsNotExist(err) {
      fmt.Println(err)
      return
    }
    _, err = os.Stat("raw_overlayed." + video_codec)
    if err == nil {
      fmt.Println("Removing raw_overlayed." + video_codec + "...")
      err = os.Remove("raw_overlayed." + video_codec)
      if err != nil {
        fmt.Println(err)
        return
      }
    } else if !os.IsNotExist(err) {
      fmt.Println(err)
      return
    }
    _, err = os.Stat("overlay_file." + video_codec)
    if err == nil {
      fmt.Println("Removing overlay_file." + video_codec + "...")
      err = os.Remove("overlay_file." + video_codec)
      if err != nil {
        fmt.Println(err)
        return
      }
    } else if !os.IsNotExist(err) {
      fmt.Println(err)
      return
    }
    
    var duration_vid string
    command_val = `ffprobe -i ` + base_racine + base_file_v[i] + ` -show_entries format=duration -v quiet -of csv="p=0"`
    out, err := exec.Command("sh", "-c", command_val).Output()
    if err != nil {
      fmt.Println(err)
      return
    }
    duration_vid = string(out)
    duration_vid = duration_vid[0:len(duration_vid) - 1]
    fmt.Println("Duration of background video in seconds: ", duration_vid)

    fmt.Println("Cutting overlay file...")
    command_val = "ffmpeg -i " + ovrl_racine + ovrl_file_v[i] + " -c copy -ss 0 -t " + duration_vid + " overlay_file." + video_codec
    cmd := exec.Command("sh", "-c", command_val)
    cmd.Start()
    cmd.Wait()

    command_val = `ffprobe -i ` + ovrl_racine + ovrl_file_v[i] + ` -show_entries format=duration -v quiet -of csv="p=0"`
    out, err = exec.Command("sh", "-c", command_val).Output()
    if err != nil {
      fmt.Println(err)
      return
    }
    duration_vid = string(out)
    fmt.Println("Duration of overlay video in seconds: ", duration_vid)
 
    if first_file_audio {
      command_val = "ffmpeg -i " + base_racine + base_file_v[i] + " -vn -acodec copy audio." + audio_codec
    } else {
      command_val = "ffmpeg -i overlay_file." + video_codec + " -vn -acodec copy audio." + audio_codec
    }
    fmt.Println("Extracting Audio from video...")
    cmd = exec.Command("sh", "-c", command_val)
    cmd.Start()
    cmd.Wait()
    command_val = `ffmpeg -i ` + base_racine + base_file_v[i] + ` -i overlay_file.` + video_codec + ` -filter_complex "[1:v]scale=` + width + `:` + height + `[ovrl];[0:v][ovrl]overlay=x=(main_w-overlay_w):y=570[outv]" -map "[outv]" -shortest raw_overlayed.` + video_codec
    fmt.Println("Overlaying...")
    cmd = exec.Command("sh", "-c", command_val)
    cmd.Start()
    cmd.Wait()
    str_i = IntToString(i)
    fmt.Println("Merging Video and Audio...")
    command_val = `ffmpeg -i raw_overlayed.` + video_codec + ` -i audio.` + audio_codec + ` -shortest -c copy out` + str_i + `.` + video_codec
    cmd = exec.Command("sh", "-c", command_val)
    cmd.Start()
    cmd.Wait()
    all_files += "file 'out" + str_i + "." + video_codec + "'\n"
  }

  var ref_i int = i
  if n > 1 {
    all_files = all_files[0:len(all_files) - 1]
    err = os.WriteFile("files.txt", []byte(all_files), 0644)
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println("Concatenating files...")
    command_val = "ffmpeg -f concat -i files.txt out." + video_codec
    cmd = exec.Command("sh", "-c", command_val)
    cmd.Start()
    cmd.Wait()
    for i = 0; i < ref_i; i++ {
      str_i = IntToString(i)
      command_val = "rm out" + str_i + "." + video_codec
      cmd = exec.Command("sh", "-c", command_val)
      cmd.Start()
      cmd.Wait()
    }
  } else {
    command_val = "mv out0." + video_codec +  " out." + video_codec
    cmd = exec.Command("sh", "-c", command_val)
    cmd.Start()
    cmd.Wait()
  }

  err = os.Remove("overlay_file." + video_codec)
  fmt.Println("Removing overlay_file." + video_codec + "...")
  if err != nil {
    fmt.Println(err)
    return
  } 
  err = os.Remove("raw_overlayed." + video_codec)
  fmt.Println("Removing raw_overlayed." + video_codec + "...")
  if err != nil {
    fmt.Println(err)
    return
  }
  err = os.Remove("audio." + audio_codec)
  fmt.Println("Removing audio." + audio_codec + "...")
  if err != nil {
    fmt.Println(err)
    return
  }
  return
}


