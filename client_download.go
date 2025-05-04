package main

import (
  "io"
  "os"
  "fmt"
  "net/http"
)

func client_downloader(url *string, created_file *string) error {
  out, err := os.Create(*created_file)
  if err != nil {
    return err
  }
  resp, err := http.Get(*url)
  if err != nil {
    return err
  }
  _, err = io.Copy(out, resp.Body)
  if err != nil {
    return err
  }
  resp.Body.Close()
  out.Close()
  return nil
}

func main() {
  url :="http://0.0.0.0:8080/download/1.jpg_Lucas_40-212-169-92-204-122-186-29-63-46-160-58-168-24-195-250-" //do not forget https if TLS on server side
  created_file := "downloaded.jpg"
  err := client_downloader(&url, &created_file)
  if err != nil {
    fmt.Println(err)
    return
  }
  return
}


