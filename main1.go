package main
//模板渲染浏览器输出
//go run main.go

//然后在浏览器中输入 :
//localhost:8880/user/info
import (
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
	myTemplate.Execute(w,p)

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
