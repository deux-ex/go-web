package main
//模块渲染存储文件
import (
	"os"
	"fmt"
	"html/template"
	"net/http"
)

var myTemplate *template.Template

type Person struct {
	Name string
	age  string
}


func userInfo(w http.ResponseWriter,r *http.Request) {

	p := Person{Name:"Murphy",age:"28"}

	myTemplate.Execute(os.Stdout,p)
	file,err := os.OpenFile("demo.dat", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("open failed err:", err)
		return
	}
	myTemplate.Execute(file,p)
}

func initTemplate(fileName string) (err error){
	myTemplate,err  = template.ParseFiles(fileName)
	if err != nil{
		fmt.Println("parse file err:",err)
		return
	}
	return
}


func main() {
	initTemplate("./index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
