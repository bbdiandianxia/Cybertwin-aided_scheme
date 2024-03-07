package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	
	//"net"
)

var clienttry *http.Client


func handle(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	msg, bfound := r.Form["msg"]
	if !bfound {
		fmt.Println("found false")
	}
	
 	message1 := msg[0]
	if message1 == "timetest" {
	
		addr_handle := "https://47.242.55.31:5050/try?msg=timetest"
		start2 := time.Now()
		
		resp2,err :=clienttry.Get(addr_handle)
		cost2 := time.Since(start2)
		
		if err != nil{
			fmt.Println(err)
		}
		content2, err := ioutil.ReadAll(resp2.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(content2))
		fmt.Printf("cost2=[%s]",cost2)
		
		fmt.Fprintf(w,string(content2))

	}else {
		fmt.Fprint(w,"only sent to ct-ctclient")

	}
}

func main(){
	pool := x509.NewCertPool()
	caCertPath := "C:/Users/PC/Desktop/token/ca.pem"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil{
		fmt.Println("Read flie erros :",err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	var qconf quic.Config
	//qconf.Allow0RTT = func(net.Addr) bool {return true}
	qconf.MaxIdleTimeout = 30 * time.Minute
	roundTripper := http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
			InsecureSkipVerify: true,
		},
		QuicConfig: &qconf,
	}
	defer roundTripper.Close()
	clienttry = &http.Client{
		Transport: &roundTripper,
	}
	addr := "https://47.242.55.31:5050/try?msg=connection"
	start := time.Now()
	resp,err := clienttry.Get(addr)
	cost := time.Since(start)
	
	if err != nil{
		fmt.Println(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
	fmt.Println("cost is ",cost)
	fmt.Println("quic-connection already built lasting 30 mins")
	//
	fmt.Println("now start listening on 9022")
	http.HandleFunc("/ctclient",handle)
	http.ListenAndServe("127.0.0.1:9022",nil)
	
	

}

