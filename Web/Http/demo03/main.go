package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"time"
)

func main() {
	http.HandleFunc("/", login)                       //设置访问的路由
	err := http.ListenAndServe("127.0.0.1:9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err) // 端口监听失败返回错误
	}
}

func login(w http.ResponseWriter, r *http.Request) { // 创建函数，用于访问对应路劲时处理对应事务
	if r.Method == "GET" { // 用户第一次访问网页基本都是GET请求
		t, _ := template.ParseFiles("C:/D/Go_code/Web/Http/demo02/login.html") // 得到GET请求后将直接跳转到登录页面
		log.Println(t.Execute(w, nil))
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		if user, _ := regexp.MatchString("^[a-zA-Z0-9]+$", username); !user {
			fmt.Println("err……")
		} else {
			fmt.Println(username)
		}
		if pass, _ := regexp.MatchString("^[a-zA-Z0-9]+$", password); !pass {
			fmt.Println("err……")
		} else {
			fmt.Println(password)
		}
		expiration := time.Now()
		expiration = expiration.AddDate(1, 0, 0)
		cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
		http.SetCookie(w, &cookie)
	}
}
