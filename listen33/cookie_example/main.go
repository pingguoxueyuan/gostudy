package main

import (
	"fmt"
	"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {

	/*cookies := r.Cookies()
	for index, cookie := range cookies {
		fmt.Printf("index:%d cookie:%#v\n", index, cookie)
	}*/
	c, err := r.Cookie("sessionid")
	fmt.Printf("cookie:%#v, err:%v\n", c, err)

	cookie := &http.Cookie{
		Name:   "sessionid",
		Value:  "lkjsdfklsjfklsfdsfdjslf",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}

	http.SetCookie(w, cookie)

	//在具体数据返回之前设置cookie，否则cookie种不上
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.ListenAndServe(":9090", nil)
}
