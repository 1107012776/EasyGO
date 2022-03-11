package server

import (
	"fmt"
	"github.com/1107012776/EasyGO/tests/controller"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var realPath *string

type MyHttpServer struct {
	content string
}

func resource(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	var requestType string
	requestType = "/"
	if path != "" && path != "/" {
		requestType = path[strings.LastIndex(path, "."):]
	}
	switch requestType {
	case ".css":
		w.Header().Set("content-type", "text/css")
	case ".js":
		w.Header().Set("content-type", "text/javascript")
	default:
		str := controller.Index()
		io.WriteString(w, str)
		return
	}
	if isExist(*realPath + path) {
		fin, err := os.Open(*realPath + path)
		defer fin.Close()
		if err != nil {
			log.Fatal("static resource:", err)
		}
		fd, _ := ioutil.ReadAll(fin)
		w.Write(fd)
	} else {
		io.WriteString(w, "404")
	}
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func Listen(port string, path string) {
	realPath = &path
	http.HandleFunc("/", resource)
	fmt.Println("Listen 0.0.0.0:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
