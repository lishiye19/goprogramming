package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "C语言中文网\n")
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("C语言中文网\n"))
}
func main() {
	http.HandleFunc("/", handler)
	log.Printf("监听12345端口成功,可以通过https://127.0.0.1:12345/访问")
	err := http.ListenAndServeTLS(":12345", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Panicln(err)
	}
	//http.HandleFunc("/hello", HelloServer)
	//err := http.ListenAndServe(":12345", nil)
	//if err != nil {
	//	log.Panicln("ListenAndServe:", err)
	//}
	//get, err := http.Get("http://c.biancheng.net")
	//if err != nil {
	//	log.Panicln("http.Get", err)
	//}
	//defer get.Body.Close()
	//all, err := ioutil.ReadAll(get.Body)
	//log.Println(string(all))

	//resp, err := http.Head("http://c.biancheng.net")
	//if err != nil {
	//	log.Panicln("http.Head error", err.Error())
	//}
	//defer resp.Body.Close()
	//for key, value := range resp.Header {
	//	log.Println(key, ":", value)
	//}
}
