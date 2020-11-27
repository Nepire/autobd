package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)
func Catflag()  (string){
    f, err := ioutil.ReadFile("./flag")
    if err != nil {
        fmt.Println("read fail", err)
    }
    return string(f)[0: len(f)-1]
}
func main() {
	var flag,token string
	token = "Teamtoken"
	flag = Catflag()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	data := url.Values{"flag":{flag}, "token":{token}}
	body := strings.NewReader(data.Encode())
	resp, err := client.Post("https://127.0.0.1:8888/submit_flag/", "application/x-www-form-urlencoded", body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(text))
}
