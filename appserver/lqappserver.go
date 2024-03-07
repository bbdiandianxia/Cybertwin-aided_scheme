package main

import (
	"fmt"
	"net/http"
)

func appseverhandler(w http.ResponseWriter,r *http.Request){
	
	fmt.Fprintf(w,"This is from appserver")
	
}

func main(){
	fmt.Println("==app server works on 6060 ==")
	http.HandleFunc("/appserver",appseverhandler)
	http.ListenAndServe("0.0.0.0:6060",nil)
}