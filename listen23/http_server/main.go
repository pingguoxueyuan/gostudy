package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "%v\n", r.Form)
	fmt.Fprintf(w, "path:%s\n", r.URL.Path)
	fmt.Fprintf(w, "schema:%s\n", r.URL.Scheme)
	fmt.Fprintf(w, "hello world\n")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

}
