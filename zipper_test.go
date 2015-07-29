package zipper

import "testing"
import "fmt"

func TestShorten(t *testing.T) {
  z := &Zipper{Service: "isgd", Url: "http://bagwanpankaj.com"}
  v, _ := z.Shorten()
  if v != "http://is.gd/3UEtSd" {
    t.Error("Expected http://is.gd/3UEtSd, got ", v)
  }
  zz := &Zipper{Service: "vgd", Url: "http://bagwanpankaj.com"}
  vv, _ := zz.Shorten()
  fmt.Println(vv)
  if vv != "http://v.gd/16H3kS" {
    t.Error("Expected http://v.gd/16H3kS, got ", vv)
  }
}
