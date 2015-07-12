package vgd

import "testing"
import "fmt"

func TestShorten(t *testing.T) {
  v, _ := Shorten("http://bagwanpankaj.com")
  fmt.Println(v)
  if v != "http://v.gd/16H3kS" {
    t.Error("Expected http://v.gd/16H3kS, got ", v)
  }
}