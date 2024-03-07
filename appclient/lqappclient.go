package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"time"
)

func main(){
	fmt.Println("type in a message which will be transfered in appserver")
	var message string
	fmt.Scanln(&message)
	data_aclient := url.Values{}
	data_aclient.Set("msg",message)
	new_ctclient_url := "http://127.0.0.1:9022/ctclient"
	uri_new, _ := url.ParseRequestURI(new_ctclient_url)
	uri_new.RawQuery=data_aclient.Encode()
	
	start := time.Now()
	
	resp,err := http.Get(uri_new.String())
	
	cost := time.Since(start)
	if err != nil {
		fmt.Println("err is",err)
		return
	}
	defer resp.Body.Close()
	cnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err is",err)
		return
	}
	fmt.Println(string(cnt))
	fmt.Println("cost is:",cost)
}