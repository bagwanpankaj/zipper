package main

import(
  "flag"
  "fmt"
  "github.com/bagwanpankaj/zipper"
)

func main() {

  var s = flag.String("s", "isgd", "specify service to use")
  var u = flag.String("u", "http://xyz.com", "specify url to shorten")
  var k = flag.String("k", "apikey", "specify api key")
  flag.Parse()

  z := zipper.New(*s, *u, *k)

  res, err := z.Shorten()
  if err != nil{
    panic(err)
  }
  fmt.Println(res)
}