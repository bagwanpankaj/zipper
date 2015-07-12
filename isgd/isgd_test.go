package isgd

import "testing"
import "fmt"

func TestShorten(t *testing.T) {
  v, _ := Shorten("http://bagwanpankaj.com")
  fmt.Println(v)
  if v != "http://is.gd/3UEtSd" {
    t.Error("Expected http://is.gd/3UEtSd, got ", v)
  }
}