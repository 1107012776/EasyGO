package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var realPath *string

func staticResource(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	request_type := path[strings.LastIndex(path, "."):]
	switch request_type {
	case ".css":
		w.Header().Set("content-type", "text/css")
	case ".js":
		w.Header().Set("content-type", "text/javascript")
	default:
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

//go run HttpServer.go --path=/tmp/static
func Listen(port string, path string) {
	//realPath = flag.String("path", "", "static resource path")
	//flag.Parse()
	realPath = &path
	http.HandleFunc("/", staticResource)
	fmt.Println("Listen 0.0.0.0:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
