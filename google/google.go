// Package google provides wrapper for url shortener services provided by `goo.gl`
package google

import (
  "bytes"
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/http"
)

// Shorten calls to shortener services with data provided and returns string
// containing shortened url and error (if any)
func Shorten(url, key string) (string, error) {
  client := &http.Client{}
  req, err := http.NewRequest("POST", "https://www.googleapis.com/urlshortener/v1/url?key="+key,
    bytes.NewBufferString("{\"longUrl\": \""+url+"\"}"))
  if err != nil {
    panic(err)
  }
  req.Header.Add("Content-Type", "application/json")
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
  shortUrl := f.(map[string]interface{})["id"].(string)
  return shortUrl, nil
}