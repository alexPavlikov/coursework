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

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir(cfg.Assets))))
	http.Handle("/data/", http.StripPrefix("/data", http.FileServer(http.Dir(cfg.Data))))

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
	tmpl, err := template.ParseFiles(cfg.Html+"index.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	var ct categoty
	err = ct.Select()
	if err != nil {
		fmt.Println("Error - main.go ct.Select()", err.Error())
	}

	tmpl.ExecuteTemplate(w, "index", ct.Rows)
}
