package bitdo

import (
  "bytes"
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/http"
  "net/url"
  // "fmt"
)

func Shorten(longUrl string) (string, error) {

  client := &http.Client{}

  parameters := url.Values{}
  parameters.Add("action", "shorten")
  parameters.Add("url_hash", "")
  parameters.Add("url", longUrl)

  req, err := http.NewRequest("POST", "http://bit.do/mod_perl/url-shortener.pl",
    bytes.NewBufferString(parameters.Encode()))
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
  // {"url_hash":"jmot"}
  urlHash := f.(map[string]interface{})["url_hash"].(string)
  shortUrl := "http://bit.do/" + urlHash
  return shortUrl, nil
}