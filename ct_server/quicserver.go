package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

func app1 (w http.ResponseWriter,r *http.Request){

	r.ParseForm()
	msg, bfound := r.Form["msg"]
	if !bfound {
		fmt.Println("found false")
	}
 	message1 := msg[0]
	if message1 == "timetest"{

		client := &http.Client{}
		addr_appserver := "http://47.242.55.31:6060/appserver"
		start := time.Now()
		resp,err := client.Get(addr_appserver)
		cost :=time.Since(start)
		if err != nil {
			fmt.Println("get resp error is ",err)
		}
		content,err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			fmt.Println("read content error is",err1)
		}
		fmt.Println(cost)
		fmt.Fprintf(w,string(content))
	}else{
		fmt.Fprintf(w,"now you have a long QUIC connection")
	}
	
}

func main(){
	
	fmt.Println("server on 5050 connection lasting 30mins")
	quicConf := &quic.Config{
		MaxIdleTimeout: 30 *time.Minute,
	}
	server := http3.Server{
		Addr:	 "0.0.0.0:5050",
		QuicConfig: quicConf,
	}
	http.HandleFunc("/try",app1)
	err := server.ListenAndServeTLS("cert.pem","priv.key")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("success!")
	}
}