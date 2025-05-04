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
  url :="http://0.0.0.0:8080/download/file_you_want" //do not forget https if TLS on server side
  created_file := "downloaded.jpg"
  err := client_downloader(&url, &created_file)
  if err != nil {
    fmt.Println(err)
    return
  }
  return
}


