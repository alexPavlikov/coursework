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
var posts post

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
	http.HandleFunc("/brandphone", brandHandler)
	http.HandleFunc("/discount", discountHandler)
	http.HandleFunc("/smartphone", smartphoneHandler)
	http.HandleFunc("/periphery", peripheryHandler)
	http.HandleFunc("/blog", blogHandler)
	http.HandleFunc("/blog/createPost", createPostHandler)
	http.HandleFunc("/admin", dbTableHandler)

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

	fmt.Println(us.Rows)

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
	err = insertUsers(db, us)
	if err != nil {
		log.Printf("Insert product failed with error %s", err)
		return
	} else {
		err = send(us.Email, us.Password, us.Name)
		if err != nil {
			fmt.Println("Error - send email message", err)
		}
	}

	// http.Redirect(w, r, "/", 200)
}

func brandHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"brandphone.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	tmpl.ExecuteTemplate(w, "brandphone", nil)
}

func smartphoneHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"smartphone.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	tmpl.ExecuteTemplate(w, "smartphone", nil)
}

func discountHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"discount.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	tmpl.ExecuteTemplate(w, "discount", nil)
}

func peripheryHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"periphery.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	tmpl.ExecuteTemplate(w, "periphery", nil)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	var p post
	p.Image = r.FormValue("Image")
	p.Title = r.FormValue("Title")
	p.Text = r.FormValue("Text")
	data := "Nov 7 2022"
	p.Data = data

	fmt.Println(p.Image, p.Title, p.Text, p.Data)

	if p.Image == "" {
		fmt.Fprint(w, "Не указали изображение")
		if p.Title == "" {
			fmt.Fprint(w, "Не указали заголовок поста")
			if p.Text == "" {
				fmt.Fprint(w, "Не указали текст поста")
			}
		}
	}

	err = insertPosts(db, posts)
	if err != nil {
		log.Printf("Insert product failed with error %s", err)
		return
	}
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"blog.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	// var posts post
	err = posts.Select()

	if err != nil {
		fmt.Println("Error - main.go posts.Select()", err.Error())
	}

	// drop := map[string]interface{}{"PostDrop": posts.Text}
	// ttl := map[string]interface{}{"Rows": posts.Rows}

	tmpl.ExecuteTemplate(w, "blog", nil)
}

/* доделать отправку данных поста на форму*/

func dbTableHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"userList.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}
	table := dbSelect()
	tmpl.ExecuteTemplate(w, "user", table)
}

func dbTable(w http.ResponseWriter, r *http.Request) {
	table := dbSelect()
	for i := range table {
		emp := table[i]
		fmt.Fprintf(w, "YESS|%12s|%12s|%12s|\n", emp.Email, emp.Name, emp.Password)
	}
}
