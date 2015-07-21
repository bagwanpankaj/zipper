package isgd

import (
  "io/ioutil"
  "net/http"
  "net/url"
  // "bytes"
  "encoding/json"
  "errors"
  // "fmt"
)

func Shorten(longUrl string) (string, error) {
  // func main(){

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
  // fmt.Println(string(body))
  if err != nil {
    return "", err
  }

  var f interface{}
  err = json.Unmarshal(body, &f)
  if err != nil {
    return "", err
  }
  // fmt.Println(f)
  urlHash := f.(map[string]interface{})
  // shortUrl := urlHash["shorturl"]
  return urlHash["shorturl"].(string), nil
}