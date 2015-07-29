// Package zipper provides high level function to for url shortener services
// It depends on `bitdo`, `google`, `isgd` and `vgd`
package zipper

import (
  "github.com/bagwanpankaj/zipper/bitdo"
  "github.com/bagwanpankaj/zipper/google"
  "github.com/bagwanpankaj/zipper/isgd"
  "github.com/bagwanpankaj/zipper/vgd"
)

// Zipper is a construct that is to hold several variables required
// for sorting a url by service
type Zipper struct {
  Service string
  Url     string
  ApiKey  string
}

// New initializes Zipper with required data to call shortener services
// and returns a Zipper pointer
func New(s, u, k interface{}) *Zipper {
  return &Zipper{Service: s.(string), Url: u.(string), ApiKey: k.(string)}
}

// Shorten calls to shortener services with data provided and returns string
// containing shortened url and error (if any)
func (z *Zipper) Shorten() (string, error) {
  switch z.Service {
  case "google":
    return google.Shorten(z.Url, z.ApiKey)
  case "bitdo":
    return bitdo.Shorten(z.Url)
  case "isgd":
    return isgd.Shorten(z.Url)
  case "vgd":
    return vgd.Shorten(z.Url)
  default:
    return google.Shorten(z.Url, z.ApiKey)
  }
}
