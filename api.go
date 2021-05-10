package woocommerce

import (
	"bytes"
  "net/http"
	"encoding/base64"
)

var username string
var password string
var host string


// Sets the Hostname for the Octoprint Connection (needs to be in the format http(s)://<server>/)
func SetHost(hostname string) {
	host = hostname
}

func isAPIKeySet() bool{
	return len(apiKey) > 0
}

func validateHostName()  {

}
func SetUsername(usernameP string)  {
	username = usernameP
}

func SetPassword(passwordP string)  {
	password = passwordP
}

// Returns the Hostname
func GetHost() string{
	return host
}

func basicAuth(username, password string) string {
  auth := username + ":" + password
  return base64.StdEncoding.EncodeToString([]byte(auth))
}

//executes a GET request against the provided Path (using the apiKey as authorization)
func Get(path string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", host+path, nil)
	req.Header.Add("Authorization","Basic " + basicAuth(username,password))
	res, _ := client.Do(req)

	return res

}
//executes a POST request against the provided Path with the provided body (using the apiKey as authorization)
func Post(path string, body []byte) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", host+path, bytes.NewBuffer(body))
	req.Header.Add("Authorization","Basic " + basicAuth(username,password))
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)
	return res
}
