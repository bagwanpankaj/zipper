package main;

import(
  "flag"
  "fmt"
  "./google"
  "./bitdo"
  "./isgd"
  "./vgd"
  "./spito"
)

func main(){

  var s = flag.String("s", "google", "specify service to use")
  var u = flag.String("u", "http://xyz.com", "specify url to shorten")
  var k = flag.String("k", "api key for google url shortner", "specify api key")
  var exp = flag.String("exp", "3600", "specify expiry time in seconds")
  flag.Parse()

  switch *s {
    case "spito":
      resp, err := spito.Shorten(*u, *exp)
      if err != nil{
        panic(err)
      }
      fmt.Println(resp)
    case "google":
      resp, err := google.Shorten(*u, *k)
      if err != nil{
        panic(err)
      }
      fmt.Println(resp)
    case "bitdo":
      respb, errr := bitdo.Shorten(*u)
      if errr != nil{
        panic(errr)
      }
      fmt.Println(respb)
    case "isgd":
      isgdr, isgderr := isgd.Shorten(*u)
      if isgderr != nil{
        panic(isgderr)
      }
      fmt.Println(isgdr)
    case "vgd":
      vgdr, vgderr := vgd.Shorten(*u)
      if vgderr != nil{
        panic(vgderr)
      }
      fmt.Println(vgdr)
    default:
      resp, err := google.Shorten(*u, *k)
      if err != nil{
        panic(err)
      }
      fmt.Println(resp)
  }
}
