package main

import (
	"io"
	"log"
	"net/http"
)

const form = `<html><body><form action ="#"method="post" name="bar">
<input type="text" name="in"/>
<input type="text" name="in"/>
<input type="submit" value="Submit"/>
</form></body></html>`

func SimleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello,world</h1>")
	panic("test,test") //错误测试
}
func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		request.ParseForm()
		io.WriteString(w, request.Form["in"][0])
		io.WriteString(w, "\n")
		io.WriteString(w, request.FormValue("in"))
	}
}
func main() {
	http.HandleFunc("/test1", logPanics(SimleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8800", nil); err != nil {

	}
}
func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic :%v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}
