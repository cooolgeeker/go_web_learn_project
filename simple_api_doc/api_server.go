// go web 简单登录例子 ， 输入 admin admin 才能登录
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

//小写的id 不能跨包调用 必须用大写
type api_struct struct {
	Id      string
	Url     string
	Method  string
	Args    string
	Content string
}

type api_all struct {
	Items []api_struct
}

func PrintApi(api api_struct) {

	data, err := json.MarshalIndent(api, "", " ")
	if err != nil {
		log.Fatalf("JSON ", err)
	}
	fmt.Printf("%s \n", data)

}

func ReadApi(file string) api_struct {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var data = string(b)
	t := strings.Split(data, "#####")
	if len(t) != 4 {
		fmt.Println("不能处理 必须是 #####  划分后为4组 ， 当前分组为" + strconv.Itoa(len(t)) + " 错误文件为 ： " + file)
		os.Exit(0)
	}
	var api = api_struct{}
	api.Id = filepath.Base(file)
	api.Url = t[0]
	api.Method = t[1]
	api.Args = t[2]
	api.Content = t[3]
	return api

}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

func getApiAll() api_all {

	files, err := ListDir("apis", ".txt")
	if err != nil {
		log.Fatal(" 读取文件失败 ")
	}
	fmt.Println(files)

	var apiAll = api_all{}

	for _, file := range files {
		//fmt.Println(path.Base(file))
		var api = ReadApi(file)
		PrintApi(api)
		apiAll.Items = append(apiAll.Items, api)
	}
	return apiAll

}

func test() {
	var a = "apis\api_1.txt"
	b := strings.Split(a, "\\")
	fmt.Println("长度 ： ")
	fmt.Println(len(b))
}
func main() {
	test()

	var apiAll = getApiAll()
	b, err := ioutil.ReadFile("api.html")
	if err != nil {
		log.Fatal(err)
	}

	var template_data = string(b)
	tmpl, err := template.New("test").Parse(template_data)
	if err != nil {
		panic(err)
	}

	var logFilename string = "api_out.html"
	logFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("open file error=%s\r\n", err.Error())
		os.Exit(-1)
	}

	defer logFile.Close()

	err = tmpl.Execute(logFile, apiAll) //将string与模板合成，变量name的内容会替换掉{{.}}
	//合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}

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
