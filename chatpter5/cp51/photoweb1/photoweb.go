package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime/debug"
)

var UPLOAD_DIR = exeabspath() + "/upload"
var TEMPLATE_DIR = exeabspath() + "/views"

var templates = make(map[string]*template.Template)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//io.WriteString(w, "<html><body><form method=\"POST\" action=\"/upload\" "+
		//	" enctype=\"multipart/form-data\">"+
		//	"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
		//	"<input type=\"submit\" value=\"Upload\" />"+
		//	"</form></body></html>")
		//files, err := template.ParseFiles("D:/GIT/GoProgrammingtest/goprogramming/chatpter5/cp51/photoweb1/upload.html")
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//files.Execute(w, nil)
		//return
		err := renderHtml(w, "upload.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)

}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	dir, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//var listHtml string
	//for _, fileInfo := range dir {
	//	imaid := fileInfo.Name()
	//	listHtml += "<li><a href=\"/view?id=" + imaid + "\">" + imaid + "</a></li>"
	//}
	//io.WriteString(w, "<html><body><ol>"+listHtml+"</ol></body></html>")
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range dir {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	//t, err := template.ParseFiles("D:/GIT/GoProgrammingtest/goprogramming/chatpter5/cp51/photoweb1/list.html")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//t.Execute(w, locals)
	err = renderHtml(w, "list.html", locals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
	//t, err := template.ParseFiles(tmpl)
	//if err != nil {
	//	return err
	//}
	//err = t.Execute(w, locals)
	//return err
	err = templates[tmpl].Execute(w, locals)
	return err
}

func exeabspath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Panicln(err)
	}
	return dir
}

func init() {
	//for _, tmpl := range []string{"upload", "list"} {
	//	t := template.Must(template.ParseFiles(TEMPLATE_DIR + "/" + tmpl + ".html"))
	//
	//	templates[tmpl] = t
	//}
	dir, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		log.Panicln(err)
		return
	}
	var templateName, templatePath string
	for _, fileInfo := range dir {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
		log.Printf("templateName=[%v]", templateName)
		log.Printf("templatePath=[%v]", templatePath)
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN:panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func main() {
	log.Printf("UPLOAD_DIR=[%v]\nTEMPLATE_DIR=[%v]\n", UPLOAD_DIR, TEMPLATE_DIR)
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe("127.0.0.1:49961", nil)
	if err != nil {
		log.Panicf("listenandserve:", err)
	}

}
