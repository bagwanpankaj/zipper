// Package isgd provides wrapper for url shortener services provided by `is.gd`
package isgd

import (
  "io/ioutil"
  "net/http"
  "net/url"
  "encoding/json"
  "errors"
)

// Shorten calls to shortener services with data provided and returns string
// containing shortened url and error (if any)
func Shorten(longUrl string) (string, error) {
  client := &http.Client{}

  parameters := url.Values{}
  parameters.Add("format", "json")
  parameters.Add("url", longUrl)

  req, err := http.NewRequest("GET", "http://is.gd/create.php?"+parameters.Encode(), nil)
  if err != nil {
    return "", err
  }
  resp, err := client.Do(req)
  if resp.StatusCode != 200 {
    return "", errors.New(resp.Status)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  var f interface{}
  err = json.Unmarshal(body, &f)
  if err != nil {
    return "", err
  }
  urlHash := f.(map[string]interface{})
  // shortUrl := urlHash["shorturl"]
  return urlHash["shorturl"].(string), nil
}