package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"go.form.cn/db"
)

var database *sql.DB

func sayHelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("Key:",k)
		fmt.Println("value:", strings.Join(v,","))
	}
	fmt.Fprint(w,"hello, welcome you!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ",r.Method)
	if r.Method == "GET"{
		t, _ := template.ParseFiles("C:\\Users\\xm\\Desktop\\GO_Materials\\go.form.cn\\ui\\login.html")
		t.Execute(w,nil)
	}else{
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		un := r.FormValue("username")
		pwd := r.FormValue("password")
		db.Insert(database,un,pwd)
	}

}
func main() {
	database = db.Db_connection()
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)         //设置访问的路由
	err:=http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	db.Query(database)
}

