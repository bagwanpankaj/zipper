package google

import(
  "net/http"
  "io/ioutil"
  "bytes"
  "errors"
  "encoding/json"
)
func Shorten(url string) (string, error) {
  client := &http.Client{}
  req, err := http.NewRequest("POST", "https://www.googleapis.com/urlshortener/v1/url", 
    bytes.NewBufferString("{\"longUrl\": \"" + url + "\"}"))
  if err != nil{
    panic(err)
  }
  req.Header.Add("Content-Type", "application/json")
  resp, err := client.Do(req)
  // return resp, err
  if(resp.StatusCode != 200){
    return "", errors.New(resp.Status)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if(err != nil){
    return "", err
    // return make(map[string]interface{}), err
  }
  // var stringify = string(body)
  // fmt.Println(stringify)

  var f interface{}
  err = json.Unmarshal(body, &f)
  if(err != nil){
    return "", err
  }
  shortUrl := f.(map[string]interface{})["id"].(string)
  return shortUrl, nil
  // fmt.Println(f.(map[string]interface{})["id"]);
}