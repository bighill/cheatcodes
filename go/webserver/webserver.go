package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type TplData struct {
	Clock string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("index.html"))
	data := TplData{
		Clock: time.Now().Format("15:04:05"),
	}

	tpl.Execute(w, data)
}

func main() {
	fmt.Printf("webserver started. open in browser...\nhttp://localhost/8080\n\n")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
