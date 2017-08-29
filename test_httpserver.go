package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

func multiserver() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "server1")
	})
	go http.ListenAndServe(":8080", nil)

	mux := http.ServeMux{}
	mux.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "server2")
	})
	go http.ListenAndServe(":80", &mux)

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func postserver() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello/login:", request.Method)
		fmt.Println("url", request.URL)
		fmt.Println("url.path", request.URL.Path)
		fmt.Println("url.rawpath", request.URL.RawPath)
		fmt.Println("url.query", request.URL.Query())
		fmt.Println("url.rawquery", request.URL.RawQuery)
		fmt.Println("url.requstURI", request.URL.RequestURI())
		fmt.Println("requestURI", request.RequestURI)
		fmt.Println("form", request.Form == nil)
		request.ParseForm()
		fmt.Println("form", request.Form)
		name := request.FormValue("name")
		fmt.Println("name", name)
		notexist := request.FormValue("notexist")
		fmt.Println("notexist", notexist)
		fmt.Println("multipartForm", request.MultipartForm)
		fmt.Println("form", request.Form)
	})
	http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		if http.MethodGet == request.Method {
			content, _ := ioutil.ReadFile("uploadicon.html")
			writer.Write(content)
		} else {
			f, h, e := request.FormFile("icon")
			if e == nil {
				fmt.Println(h.Size)
				fmt.Println(f)
			} else {
				fmt.Println(e)
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}

func httpsserver() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello https")
	})
	e := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	fmt.Println(e)
}

func main() {
	httpsserver()
}
