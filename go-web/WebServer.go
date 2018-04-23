package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhellName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()     //解析参数
	fmt.Print(r.Form) //输出到服务端
	fmt.Print("path", r.URL.Path)
	fmt.Print("scheme", r.URL.Scheme)
	fmt.Print(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //写到客户端

}

func main() {
	http.HandleFunc("/", sayhellName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
