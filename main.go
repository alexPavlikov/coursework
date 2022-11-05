package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var m manager

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
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/registration", regHandler)
	http.HandleFunc("/regHandlerPost", regHandlerPost)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	admin := false

	if m.Role != "" {
		if m.Role == "admin" {
			admin = true
		} else {
			fmt.Println("NO admin")
		}
	}

	tmpl, err := template.ParseFiles(cfg.Html+"index.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	var ct categoty
	err = ct.Select()
	if err != nil {
		fmt.Println("Error - main.go ct.Select()", err.Error())
	}

	rows := map[string]interface{}{"Rows": ct.Rows}
	role := map[string]interface{}{"Role": admin}

	tmpl.ExecuteTemplate(w, "index", rows) //struct{ Admin string }{Admin: admin}
	tmpl.ExecuteTemplate(w, "index", role)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, er := template.ParseFiles(cfg.Html+"index.html", cfg.Html+"footer.html")
	if er != nil {
		http.NotFound(w, r)
	}

	data := json.NewDecoder(r.Body)
	data.DisallowUnknownFields()
	err := data.Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = m.Select()
	if err != nil {
		fmt.Println("Error - incorrect login/password", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(m.Login, m.Password, m.Name, m.Role)
	// if m.Name != "" {
	// men := map[string]interface{}{"Men": m.Name}
	tmpl.ExecuteTemplate(w, "index", nil)
	// http.Redirect(w, nil, "/", 200)
	// 	//struct{ Men string }{Men: m.Name}
	// }

	// tmpl.ExecuteTemplate(w, "index", nil)
}

func regHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"registration.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	tmpl.ExecuteTemplate(w, "registr", nil)
}

func regHandlerPost(w http.ResponseWriter, r *http.Request) {

	var us user
	us.Email = r.FormValue("email")
	us.Password = r.FormValue("pass1")
	checkpass := r.FormValue("pass2")
	us.Name = r.FormValue("name")

	fmt.Println(us.Email, us.Password, checkpass, us.Name)

	if us.Email == "" || len(us.Email) < 5 {
		fmt.Fprintf(w, "Неверный формат логина - %s", us.Email)
		if us.Password == "" || len(us.Password) < 8 {
			fmt.Fprintf(w, "Неверный формат пароля - %s", us.Password)
			if checkpass == "" {
				fmt.Fprintf(w, "Неверный формат пароля - %s", checkpass)
				if us.Name == "" {
					fmt.Fprintf(w, "Неверный формат имени - %s", us.Name)
				}
			}
		}
	}
	if us.Password != checkpass {
		fmt.Fprintf(w, "Пароли не соответствуют друг другу - %s и %s", us.Password, checkpass)
	}

	err = insert(db, us)
	if err != nil {
		log.Printf("Insert product failed with error %s", err)
		return
	}

	http.Redirect(w, r, "/", 200)
	// fmt.Fprintf(w, "Email: %s Password: %s  Pass: %s Name: %s", email, password, checkpass, name)
}
