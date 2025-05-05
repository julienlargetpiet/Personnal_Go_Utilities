package main

import (
  "os"
  "io"
  "net/http"
  "mime/multipart"
  "bytes"
  "fmt"
)

func client_upload(file_to_upload *string, url *string) error {
  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  part, err := writer.CreateFormFile("myfile", "myfile")
  if err != nil {
    return err
  }
  file, err := os.Open(*file_to_upload)
  if err != nil {
    return err
  }
  _, err = io.Copy(part, file)
  if err != nil {
    return err
  }
  err = writer.Close()
  if err != nil {
    return err
  }
  file.Close()
  client := &http.Client{}
  req, err := http.NewRequest("POST", *url, body)
  if err != nil {
    return err
  }
  req.Header.Add("Content-Type", writer.FormDataContentType())
  resp, err := client.Do(req)
  if err != nil {
    return err
  }
  resp.Body.Close()
  return nil
}

func main() {
  url := "http://0.0.0.0:8080/upload" // do not forget https if TLS on server side
  file_to_upload := "1.jpg"
  err := client_upload(&file_to_upload, &url)
  if err != nil {
    fmt.Println(err)
  }
  return
}


