package main

import (
	"fmt"
	"os"
	"text/template"
)

//模板渲染终端
//模板替换 {{.字段名}}

type Person struct {
	Name string
	age  string
}

func main() {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", age: "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}

