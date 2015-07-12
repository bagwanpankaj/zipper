package spito

import(
  //"net/url"
  "net/http"
  "io/ioutil"
  "errors"
  "encoding/json"
  "bytes"
  //"fmt"
  "mime/multipart"
)

// exp should be the expiry time in seconds - string format
func Shorten(longUrl string, exp string) (string, error) {

  client := &http.Client{}
 
  var b bytes.Buffer
  w := multipart.NewWriter(&b)

  fw, err := w.CreateFormField("spit_type")
  if err != nil {
    return "", err
  }
  if _, err = fw.Write([]byte("url")); err != nil {
    return "", err
  }
  if fw, err = w.CreateFormField("content"); err != nil {
    return "", err
  }
  if _, err = fw.Write([]byte(longUrl)); err != nil {
    return "", err
  }
  if fw, err = w.CreateFormField("exp"); err != nil {
    return "", err
  }
  if _, err = fw.Write([]byte(exp)); err != nil {
    return "", err
  }

  w.Close()

  req, err := http.NewRequest("POST", "http://spi.to/api/v1/spits", &b)
  if err != nil{
    return "", err
  }
  // Don't forget to set the content type, this will contain the boundary.
  req.Header.Set("Content-Type", w.FormDataContentType())
  resp, err := client.Do(req)
  //fmt.Println(resp)
  if(resp.StatusCode != 200){
    return "", errors.New(resp.Status)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  //fmt.Println(string(body))
  if(err != nil){
    return "", err
  }

  var f interface{}
  err = json.Unmarshal(body, &f)
  if(err != nil){
    return "", err
  }
  // fmt.Println(f)
  urlHash := f.(map[string]interface{})
  return urlHash["absolute_url"].(string), nil
}
