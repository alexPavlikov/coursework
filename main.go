package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
)

var m manager
var pe purchase
var ajax_post_data string
var new_post_data string
var incorrect string
var tr *string

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

// -------------------------Объявление HandleFunc-------------------------
func handlerRequest() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/login/access", loginAccessHandler)
	http.HandleFunc("/registration", regHandler)
	http.HandleFunc("/regHandlerPost", regHandlerPost)
	http.HandleFunc("/brandphone", brandHandler)
	http.HandleFunc("/discount", discountHandler)
	http.HandleFunc("/smartphone", smartphoneHandler)
	http.HandleFunc("/periphery", peripheryHandler)
	http.HandleFunc("/blog", blogHandler)
	http.HandleFunc("/blog/", createPostHandler)
	http.HandleFunc("/admin", dbTableHandler)
	http.HandleFunc("/admin/man", addManagerHandler)
	http.HandleFunc("/admin/user", addUserHandler)
	http.HandleFunc("/admin/del", delUserHandler)
	http.HandleFunc("/admin/series", addSerialHandler)
	http.HandleFunc("/admin/manager", adminTableHandler)
	http.HandleFunc("/productList", productTableHandler)
	http.HandleFunc("/productList/addProduct", productAddHandler)
	http.HandleFunc("/productList/addProduct/add", addPrHandler)
	http.HandleFunc("/postList", postTableHandler)
	http.HandleFunc("/postList/", postDelHandler)
	http.HandleFunc("/purchaseList", purchaseTableHandler)
	http.HandleFunc("/delUserList", delUserTableHandler)
	http.HandleFunc("/seriesList", seriesTableHandler)
	http.HandleFunc("/purchaseList/addPruchase", addPurchaseHandler)
	http.HandleFunc("/purchaseList/addPruchase/buy", buyPurchaseHandler) //-------------------
	http.HandleFunc("/admin/serdel", delSeriesHandler)
	http.HandleFunc("/product/del", productDelHandler)
	http.HandleFunc("/statistics", statHandler)
	http.HandleFunc("/brandphone/access", accessHandler)
}

// -------------------------Реализация HandleFunc-------------------------
func loginHandler(w http.ResponseWriter, r *http.Request) {

	var ct series

	err = ct.Select()
	if err != nil {
		fmt.Println("Error - main.go ct.Select()", err.Error())
	}

	data := json.NewDecoder(r.Body)
	data.DisallowUnknownFields()
	err = data.Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = m.Select()
	if err != nil {
		fmt.Println("Error - incorrect login/password", err.Error())
		incorrect = "Такого пользователя не существует"
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(m.Login, m.Password, m.Name, m.Role)
	tr = &m.Name

	err = loginWarning(m.Login, m.Name)
	if err != nil {
		fmt.Println("Возникла ошибка в отправе сообщения")
	}

	//fmt.Fprintf(w, `<h1 class="News-title">Вы вошли в аккаунт email:%s; password: %s; name:%s</h1>`, m.Login, m.Password, m.Name)
}

func loginAccessHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}
	brick := *tr
	if brick != "" && m.Role == "" {
		fmt.Fprintf(w, `<h1 class="News-title">Вы успешно авторизовались</h1>`)
		//brick = ""
	} else if m.Role != "" {
		fmt.Fprintf(w, `<h1 class="News-title">Вы успешно авторизовались под ролью администратора</h1>`)
	} else {
		fmt.Fprintf(w, `<h1 class="News-title">Ошибка! Такого пользователя не существует</h1>`)
		//brick = ""
	}
	tmpl.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(cfg.Html+"index.html", cfg.Html+"footer.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}
	m.Role = ""
	m.Name = ""
	var ct series
	err = ct.Select()
	if err != nil {
		fmt.Println("Error - main.go ct.Select()", err.Error())
	}

	log := r.FormValue("lgn")
	pass := r.FormValue("pass")
	LogUser(log, pass)

	rows := map[string]interface{}{"Rows": ct.Rows}
	tmpl.ExecuteTemplate(w, "index", nil)
	tmpl.ExecuteTemplate(w, "header", rows)
}

func regHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"registration.html", cfg.Html+"footer.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	var ct series
	err = ct.Select()
	if err != nil {
		fmt.Println("Error - main.go ct.Select()", err.Error())
	}

	tmpl.ExecuteTemplate(w, "registr", nil)
	rows := map[string]interface{}{"Rows": ct.Rows}
	tmpl.ExecuteTemplate(w, "header", rows)
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
		return
	}
	if us.Password != checkpass {
		fmt.Fprintf(w, "Пароли не соответствуют друг другу - %s и %s", us.Password, checkpass)
		return
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
	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)

	tmpl.Execute(w, nil)
}

var prodq product

func brandHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"brandphone.html", cfg.Html+"footer.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	var ct series
	err = ct.Select()
	if err != nil {
		fmt.Println("Error - main.go ct.Select()", err.Error())
	}
	table := productSelect()
	per := peripherySelect()
	r.ParseForm()
	if r.Method == "POST" {
		ajax_post_data = r.FormValue("ajax_post_data")
		fmt.Println(ajax_post_data)
		prodq = oneProductSelect(ajax_post_data)
	}
	if r.Method == "GET" {
		new_post_data = r.FormValue("new_post_data")
		prodq = oneProductSelect(new_post_data)
	}
	data := map[string]interface{}{"Product": table, "Periphery": per}
	tmpl.ExecuteTemplate(w, "brandphone", data)
	rows := map[string]interface{}{"Rows": ct.Rows}
	tmpl.ExecuteTemplate(w, "header", rows)
}

func accessHandler(w http.ResponseWriter, r *http.Request) {

	html := fmt.Sprintf(`
	<head>
	<meta http-equiv="refresh" content="1" />
    <link rel="shortcut icon" href="/data/site.ico">
    <title>%s</title>
	</head>
	<body>
		<section class="section">
			<div class="brand_list">
				<div class="brand-half">
					<div id="%d" class="item itemN">
						<div class="discount">Скидка 13%%</div>
						<div class="text">
							
							<div class="name"> %s <span id="spanName"> %s </span></div>
							<div class="description"> %s <span></span><br>В наличие (шт): %d</div>
							<div class="price">
								
								<div class="store-week_newPrice"> %.2f <span>рублей</span></div>
								<div class="store-week_oldPrice"></div>
							</div>
						</div>
						<div class="btns">
							<a style="color:white" href="/purchaseList/addPruchase" class="pre-order__button">Купить</a>							
						</div>
						<div class="image">
							<img src="%s" alt="">
						</div>
					</div>
				</div>
			</div>
		</section>
	</body>	
	`, prodq.Name, prodq.Article, prodq.Series, prodq.Name, prodq.Description, prodq.Count, prodq.Price, prodq.Image)
	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}
	fmt.Fprint(w, html)
	tmpl.Execute(w, nil)
}

func smartphoneHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"smartphone.html", cfg.Html+"footer.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	var ct series
	err = ct.Select()
	if err != nil {
		fmt.Println("Error - main.go ct.Select()", err.Error())
	}

	redmi := "Redmi"
	xiaomi := "Xiaomi"
	poco := "Poco"

	res1 := brandSelect(redmi, 8)
	res2 := brandSelect(xiaomi, 4)
	res3 := brandSelect(poco, 4)

	r.ParseForm()
	if r.Method == "POST" {
		ajax_post_data = r.FormValue("ajax_post_data")
		fmt.Println(ajax_post_data)
		prodq = oneProductSelect(ajax_post_data)
	}
	if r.Method == "GET" {
		new_post_data = r.FormValue("new_post_data")
		prodq = oneProductSelect(new_post_data)
	}

	br := map[string]interface{}{"Redmi": res1, "Xiaomi": res2, "Poco": res3}
	tmpl.ExecuteTemplate(w, "smartphone", br)
	rows := map[string]interface{}{"Rows": ct.Rows}
	tmpl.ExecuteTemplate(w, "header", rows)
}

func discountHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"discount.html", cfg.Html+"footer.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}
	var ct series
	err = ct.Select()
	if err != nil {
		fmt.Println("Error - main.go ct.Select()", err.Error())
	}

	tmpl.ExecuteTemplate(w, "discount", nil)
	rows := map[string]interface{}{"Rows": ct.Rows}
	tmpl.ExecuteTemplate(w, "header", rows)
}

func peripheryHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"periphery.html", cfg.Html+"footer.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}
	var ct series
	err = ct.Select()
	if err != nil {
		fmt.Println("Error - main.go ct.Select()", err.Error())
	}

	per := peripherySelect()
	row := map[string]interface{}{"Periphery": per}
	tmpl.ExecuteTemplate(w, "periphery", row)
	rows := map[string]interface{}{"Rows": ct.Rows}
	tmpl.ExecuteTemplate(w, "header", rows)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {

	var pst post
	pst.Image = r.FormValue("Image")
	pst.Title = r.FormValue("Title")
	pst.Text = r.FormValue("Text")
	data := time.Now().Format("2006-01-02 15:04")
	pst.Data = data

	fmt.Println(pst.Image, pst.Title, pst.Text, pst.Data)

	if pst.Image == "" {
		fmt.Fprint(w, "Не указали изображение")
		if pst.Title == "" {
			fmt.Fprint(w, "Не указали заголовок поста")
			if pst.Text == "" {
				fmt.Fprint(w, "Не указали текст поста")
			}
		}
	}

	err = insertPosts(db, pst)
	if err != nil {
		log.Printf("Insert product failed with error %s", err)
		return
	}

	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)

	tmpl.Execute(w, nil)
}

var views int64

func blogHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(cfg.Html+"blog.html", cfg.Html+"footer.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	} else {
		views++
	}
	var ct series
	err = ct.Select()
	if err != nil {
		fmt.Println("Error - main.go ct.Select()", err.Error())
	}

	var posts post
	err = posts.Select()
	if err != nil {
		fmt.Println("Error - main.go posts.Select()", err.Error())
	}

	rows := map[string]interface{}{"Rows": ct.Rows}

	this := map[string]interface{}{"Post": posts.Rows, "Views": views}
	tmpl.ExecuteTemplate(w, "blog", this)
	tmpl.ExecuteTemplate(w, "header", rows)

}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	var us user

	us.Email = r.FormValue("Email")
	us.Password = r.FormValue("Password")
	us.Name = r.FormValue("Name")

	fmt.Println(us.Email, us.Password, us.Name)

	if us.Email == "" || len(us.Email) < 5 {
		fmt.Fprintf(w, "Неверный формат логина - %s", us.Email)
		if us.Password == "" || len(us.Password) < 8 {
			fmt.Fprintf(w, "Неверный формат пароля - %s", us.Password)
			if us.Name == "" {
				fmt.Fprintf(w, "Неверный формат имени - %s", us.Name)
			}
		}
		return
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

	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)

	tmpl.Execute(w, nil)

}

func addManagerHandler(w http.ResponseWriter, r *http.Request) {

	m.Login = r.FormValue("Email")
	m.Password = r.FormValue("Password")
	m.Name = r.FormValue("Name")
	m.Role = "admin"

	fmt.Println(m.Login, m.Password, m.Name, m.Role)

	if m.Login == "" || len(m.Login) < 5 {
		fmt.Fprintf(w, "Неверный формат логина - %s", m.Login)
		if m.Password == "" || len(m.Password) < 8 {
			fmt.Fprintf(w, "Неверный формат пароля - %s", m.Password)
			if m.Name == "" {
				fmt.Fprintf(w, "Неверный формат имени - %s", m.Name)
			}
		}
		return
	}
	err = insertManager(db, m)
	if err != nil {
		log.Printf("Insert product failed with error %s", err)
		return
	} else {
		err = send(m.Login, m.Password, m.Name)
		if err != nil {
			fmt.Println("Error - send email message", err)
		}
	}

	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)

	tmpl.Execute(w, nil)

}

func addSerialHandler(w http.ResponseWriter, r *http.Request) {
	var s series

	s.Name = r.FormValue("Name")

	fmt.Println(s.Name)

	if s.Name == "" {
		fmt.Fprintf(w, "Неверный формат логина - %s", s.Name)
		return
	}
	err = insertSeries(db, s)
	if err != nil {
		log.Printf("Insert product failed with error %s", err)
		return
	}

	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)

	tmpl.Execute(w, nil)
}

/* Обработчики таблиц*/

func dbTableHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"userList.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}
	table := dbSelect()
	tmpl.ExecuteTemplate(w, "user", table)
}

func adminTableHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html + "adminList.html")
	if err != nil {
		http.NotFound(w, r)
	}
	table := adminSelect()
	tmpl.ExecuteTemplate(w, "admin", table)
}

func productTableHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html + "productList.html")
	if err != nil {
		http.NotFound(w, r)
	}
	table := productSelect()
	tmpl.ExecuteTemplate(w, "product", table)
}

func postTableHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html + "postList.html")
	if err != nil {
		http.NotFound(w, r)
	}
	table := postSelect()
	tmpl.ExecuteTemplate(w, "postList", table)
}

func purchaseTableHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html + "purchaseList.html")
	if err != nil {
		http.NotFound(w, r)
	}
	table := purchaseSelect()

	tmpl.ExecuteTemplate(w, "purchase", table)
}

func seriesTableHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html + "seriesList.html")
	if err != nil {
		http.NotFound(w, r)
	}
	table := seriesSelect()
	tmpl.ExecuteTemplate(w, "series", table)
}

func delUserTableHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html + "delUserList.html")
	if err != nil {
		http.NotFound(w, r)
	}
	table := delUserSelect()

	tmpl.ExecuteTemplate(w, "delUser", table)
}

func delUserHandler(w http.ResponseWriter, r *http.Request) {
	var usb userBanned

	usb.Email = r.FormValue("Email")
	usb.Reason = r.FormValue("Reason")

	fmt.Println(usb.Email, usb.Reason)

	if usb.Email == "" || len(usb.Email) < 5 {
		fmt.Fprintf(w, "Неверный формат логина - %s", usb.Email)
		return
	}
	err = deleteUsers(db, usb)
	if err != nil {
		log.Printf("Delete DelUsers failed with error %s", err)
		return
	}

	err = insertDelUser(db, usb)
	if err != nil {
		log.Printf("Insert DelUsers failed with error %s", err)
		return
	}

	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)

	tmpl.Execute(w, nil)
}

type foo struct {
	id       int
	valuePur int
	userPur  string
	products int
	price    float64
	tprice   float64
	data     string
}

var f foo

func addPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"addPurchase.html", cfg.Html+"header.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	item := purSelect()

	elem := purSelect()
	i := r.FormValue("count")
	f.valuePur, _ = strconv.Atoi(i)

	item.TotalPrice = item.Product.Price * float64(f.valuePur)
	fmt.Println(item.TotalPrice)
	f.id = elem.Id + 1
	f.userPur = r.FormValue("selectuser")
	b := r.FormValue("selectproduct")
	f.products, _ = strconv.Atoi(b)
	f.price = item.Product.Price
	f.tprice = item.TotalPrice
	f.data = time.Now().Format("2006-01-02 15:04")

	purch := map[string]interface{}{"User": item.User.Rows, "Product": item.Product.Rows, "Price": item.Product.Price, "TotalPrice": item.TotalPrice}
	tmpl.ExecuteTemplate(w, "addPurchase", purch)

}

func buyPurchaseHandler(w http.ResponseWriter, r *http.Request) {

	err = insertPurchase(db, f)
	if err != nil {
		log.Printf("Insert purchase failed with error %s", err)
		return
	}

	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)
	err = sendPurchase(f)
	if err != nil {
		fmt.Println("buyPurchaseHandler", err)
	}
	tmpl.Execute(w, nil)
}

func delSeriesHandler(w http.ResponseWriter, r *http.Request) {
	var s series

	s.Name = r.FormValue("Series")
	fmt.Println(s.Name)

	err = deleteSeries(db, s)
	if err != nil {
		log.Printf("Delete DelUsers failed with error %s", err)
		return
	}

	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)

	tmpl.Execute(w, nil)

}

var prod product

func productAddHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"addProduct.html", cfg.Html+"header.html", cfg.Html+"footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Println("Go")
	}
	err = r.ParseMultipartForm(1200)
	if err != nil {
		fmt.Println("Go")
	}

	table := seriesSelect()
	art := r.FormValue("article")
	prod.Article, _ = strconv.Atoi(art)
	prod.Series = r.FormValue("series")
	fmt.Println(prod.Series)
	prod.Name = r.FormValue("name")
	prod.Price, _ = strconv.ParseFloat(r.FormValue("price"), 64)
	count := r.FormValue("count")
	prod.Count, _ = strconv.Atoi(count)
	prod.Image = r.FormValue("image")
	prod.Description = r.FormValue("description")

	fmt.Println("Считан товар - ", prod)

	rows := map[string]interface{}{"Series": table}

	tmpl.ExecuteTemplate(w, "addProduct", rows)
}

func addPrHandler(w http.ResponseWriter, r *http.Request) {
	err := insertProduct(db, prod)
	if err != nil {
		fmt.Println("Error - insertProduct()", err)
	}

	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)

	tmpl.Execute(w, nil)
}

func postDelHandler(w http.ResponseWriter, r *http.Request) {
	var pt post

	pt.Id, _ = strconv.Atoi(r.FormValue("ID"))
	fmt.Println(pt.Id)

	err = deletePost(db, pt)
	if err != nil {
		log.Printf("Delete DelPost failed with error %s", err)
		return
	}

	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)

	tmpl.Execute(w, nil)
}

func productDelHandler(w http.ResponseWriter, r *http.Request) {
	var pd product

	pd.Article, _ = strconv.Atoi(r.FormValue("ID"))

	err = deleteProduct(db, pd)
	if err != nil {
		log.Printf("Delete DelPost failed with error %s", err)
		return
	}

	tmpl, err := template.ParseFiles(cfg.Html+"result.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, `<h1 class="News-title">Операция успешна выполнена</h1>`)

	tmpl.Execute(w, nil)
}

func statHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.Html+"statistics.html", cfg.Html+"footer.html", cfg.Html+"header.html")
	if err != nil {
		http.NotFound(w, r)
	}
	redmi := "Redmi"
	xiaomi := "Xiaomi"
	poco := "Poco"

	col1 := valueBrandSelect(redmi)
	col2 := valueBrandSelect(xiaomi)
	col3 := valueBrandSelect(poco)

	resRedmi := priceBrandSelect(redmi)
	var result1 float64
	for _, v := range resRedmi {
		result1 += v
	}

	resXiaomi := priceBrandSelect(xiaomi)
	var result2 float64
	for _, b := range resXiaomi {
		result2 += b
	}

	resPoco := priceBrandSelect(poco)
	var result3 float64
	for _, c := range resPoco {
		result3 += c
	}

	pr, dt := purchPriceSelect()
	type area struct {
		Price []float64
		Date  []string
	}
	var ar area
	ar.Price = pr
	ar.Date = dt
	rows := map[string]interface{}{"Col1": col1, "Col2": col2, "Col3": col3, "TPRedmi": result1, "TPXiaomi": result2, "TPPoco": result3, "Date": dt, "Price": pr}
	tmpl.ExecuteTemplate(w, "stat", rows)
}

func getElem(page string, elem string, id string) string {
	data, err := ioutil.ReadFile(page)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}
	str := fmt.Sprintf(elem + "#" + id)
	text := doc.Find(str)
	return text.Text()
}
