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

  _, err := os.Stat("audio." + audio_codec)
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

  var duration_vid string
  command_val = `ffprobe -i ` + base_file + ` -show_entries format=duration -v quiet -of csv="p=0"`
  out, err := exec.Command("sh", "-c", command_val).Output()
  if err != nil {
    fmt.Println(err)
    return
  }
  duration_vid = string(out)
  duration_vid = duration_vid[0:len(duration_vid) - 1]
  fmt.Println("Duration of background video in seconds: ", duration_vid)

  fmt.Println("Cutting overlay file...")
  command_val = "ffmpeg -i " + ovrl_file + " -c copy -ss 0 -t " + duration_vid + " overlay_file." + video_codec
  cmd := exec.Command("sh", "-c", command_val)
  cmd.Start()
  cmd.Wait()

  command_val = `ffprobe -i ` + ovrl_file + ` -show_entries format=duration -v quiet -of csv="p=0"`
  out, err = exec.Command("sh", "-c", command_val).Output()
  if err != nil {
    fmt.Println(err)
    return
  }
  duration_vid = string(out)
  fmt.Println("Duration of overlay video in seconds: ", duration_vid)
 
  if first_file_audio {
    command_val = "ffmpeg -i " + base_file + " -vn -acodec copy audio." + audio_codec
  } else {
    command_val = "ffmpeg -i overlay_file." + video_codec + " -vn -acodec copy audio." + audio_codec
  }
  fmt.Println("Extracting Audio from video...")
  cmd = exec.Command("sh", "-c", command_val)
  cmd.Start()
  cmd.Wait()
  command_val = `ffmpeg -i ` + base_file + ` -i overlay_file.` + video_codec + ` -filter_complex "[1:v]scale=` + width + `:` + height + `[ovrl];[0:v][ovrl]overlay=x=(main_w-overlay_w):y=570[outv]" -map "[outv]" -shortest raw_overlayed.` + video_codec
  fmt.Println("Overlaying...")
  cmd = exec.Command("sh", "-c", command_val)
  cmd.Start()
  cmd.Wait()
  fmt.Println("Merging Video and Audio...")
  command_val = `ffmpeg -i raw_overlayed.` + video_codec + ` -i audio.` + audio_codec + ` -shortest -c copy out.` + video_codec
  cmd = exec.Command("sh", "-c", command_val)
  cmd.Start()
  cmd.Wait()
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


