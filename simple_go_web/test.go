// go web 简单登录例子 ， 输入 admin admin 才能登录
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", handler_index)      // each request calls handler
	http.HandleFunc("/login", handler_login) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 处理首页请求，返回登录页面
func handler_index(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", b)

}

// 处理登录请求，返回登录成功或者失败
func handler_login(w http.ResponseWriter, r *http.Request) {

	type login struct {
		Name   string //登录用户名
		Result string //登录结果
	}
	var login_info = new(login)

	//解析参数
	r.ParseForm()

	fmt.Println("username", r.Form["name"])
	fmt.Println("password", r.Form["password"])

	//遍历解析的参数，注意是一个数组
	for k, v := range r.Form {
		fmt.Println("key:", k, "value : ", v)
	}

	if len(r.Form["name"]) == 1 && len(r.Form["password"]) == 1 && r.Form["name"][0] == "admin" && r.Form["password"][0] == "admin" {
		login_info.Name = "admin"
		login_info.Result = "成功"
	} else {
		login_info.Name = r.Form["name"][0]
		login_info.Result = "失败"
	}

	tmpl, err := template.New("test").Parse("<html>执行结果 {{.Result}}, 用户名 {{.Name}}  </html>") //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, login_info) //将string与模板合成，变量name的内容会替换掉{{.}}
	//合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
}
