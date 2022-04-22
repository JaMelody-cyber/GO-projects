package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

// producerDemo.go

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("E:\\GoProjects\\webDemo\\hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w

	tmpl.Execute(w, " 黎声益")
}

func sayHello2(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("E:\\GoProjects\\webDemo\\hello.tmpl")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	// 自定义一个夸人的模板函数
	kua := func(arg string) (string, error) {
		return arg + "真帅a", nil
	}
	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
	tmpl := template.New("hello2")
	tmpl = tmpl.Funcs(template.FuncMap{"kuakuawo": kua})
	tmpl, err = tmpl.Parse(string(htmlByte))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	// 使用user渲染模板，并将结果写入w
	tmpl.Execute(w, user)
}
func main() {
	http.HandleFunc("/", sayHello2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
