package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Listen on - " + cfg.ServerHost + ":" + cfg.ServerPort)

	err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	handlerRequest()
	err = http.ListenAndServe(":"+cfg.ServerPort, nil)
	if err != nil {
		log.Fatal("Error ListenAndServe no worked ", err.Error())
	}
}

// -------------------------HandleFunc-------------------------
func handlerRequest() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html + "index.html")
	if err != nil {
		http.NotFound(w, r)
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}
