package main

//import (
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//)
//
//const form = `<html><body>
//        <form action="#" method="post" />
//            <input type="text" name="name" />
//            <input type="text" name="pass" />
//            <input type="submit" value="submit" />
//        </form>
//        </body></html>`
//
//func deal2(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html")
//
//	//Method 代表指定的http方法
//	switch r.Method {
//	case "GET":
//		//io.WriteString(w, s string) 将字符串的s写入w中
//		io.WriteString(w, form)
//	case "POST":
//		//func (*Request) PostFormValue
//		//PostFormValue返回key为键查询r.PostForm字段得到结果[]string切片的第一个值
//		io.WriteString(w, r.PostFormValue("name"))
//		io.WriteString(w, "\n")
//		io.WriteString(w, r.PostFormValue("pass"))
//	}
//}
//
////panic 处理
//func logsPanic(handle http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		defer func() {
//			if x := recover(); x != nil {
//				log.Printf("[%v] catch panic:%v", r.RemoteAddr, x)
//			}
//		}()
//		handle(w, r)
//	}
//}
//
//func main() {
//	http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "<h1>hello world</h1>")
//	})
//
//	http.HandleFunc("/test2", logsPanic(deal2))
//
//	if err := http.ListenAndServe("localhost:10010", nil); err != nil {
//		fmt.Println("listen server failed, error:", err)
//		return
//	}
//
//}
