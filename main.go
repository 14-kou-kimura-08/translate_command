package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "strings"
)

func main() {
  flag.Parse()
  args := strings.Join(flag.Args(), " ")
  source := "ja"
  target := "en"
  url := "https://script.google.com/macros/s/AKfycbzH5kPvN6nG0XnzfPJhbrxPlBhnb0vOXTdob0WKygLyYX7N2Fg/exec?text=" + url.QueryEscape(args) + "&source=" + source + "&target=" + target

  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
  }

  defer resp.Body.Close()

  byteArray, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(byteArray))
}
