package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type UserInfo struct {
	Name string
	Sex  string
	Age  int
}

func login(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./login.html")
	if err != nil {
		fmt.Fprintf(w, "load login.html failed")
		return
	}
	/*
		user := UserInfo{
			Name: "Mary",
			Sex:  "男",
			Age:  18,
		}
	*/
	m := make(map[string]interface{})
	m["username"] = "Mary"
	m["sex"] = "男"
	m["age"] = 18
	//t.Execute(w, "Mary")
	//t.Execute(w, 1000)
	//t.Execute(w,  )
	t.Execute(w, m)
	t.Execute(os.Stdout, m)
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("listen server failed, err:%v\n", err)
		return
	}
}
