package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type categoty struct {
	Name string
	Rows []categoty
}

type manager struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
	Role     string `json:"Role"`

	Rows []manager
}

type user struct {
	Email    string
	Password string
	Name     string

	Rows []user
}

type userBanned struct {
	Email  string
	Reason string

	Rows []userBanned
}

type post struct {
	Id    int
	Image string
	Title string
	Text  string
	Data  string

	Rows []post
}

type product struct {
	Article     int
	Series      string
	Name        string
	Price       float64
	Count       int
	Image       string
	Description string

	Rows []product
}

type purchase struct {
	Id         int     `json:"Id"`
	User       string  `json:"User"`
	Product    int     `json:"Product"`
	Count      int     `json:"Count"`
	Price      float64 `json:"Price"`
	Date       string  `json:"Date"`
	TotalPrice float64 `json:"TotalPrice"`

	Rows []purchase
}

type pur struct {
	Id         int
	User       user
	Product    product
	Count      uint16
	Price      product
	Date       string
	TotalPrice float64

	Rows []pur
}

type series struct {
	Name string

	Rows []series
}

var (
	db         *sql.DB
	err        error
	Queries    map[string]*sql.Stmt
	countPosts int
)

func connect() error {
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgName))
	if err != nil {
		log.Fatal("Error - database connect ")
		return err
	}
	Queries = make(map[string]*sql.Stmt)

	prepareQueries()

	return nil
}

func prepareQueries() {
	Queries["Select#Series"], err = db.Prepare(`SELECT "Name" FROM "Series" ORDER BY "Name"`)
	if err != nil {
		fmt.Println("Ошибка запроса Select#Series ")
	}

	Queries["Select#User"], err = db.Prepare(`SELECT "Name" FROM "Users" WHERE "Email"=$1 AND "Password"=$2`)
	if err != nil {
		fmt.Println("Ошибка запроса Select#User ")
	}

	Queries["Select#Users"], err = db.Prepare(`SELECT * FROM "Users"`)
	if err != nil {
		fmt.Println("Ошибка запроса Select#Users ")
	}

	Queries["Select#Manager"], err = db.Prepare(`SELECT "Name", "Role" FROM "Manager" WHERE "Email"=$1 AND "Password"=$2`)
	if err != nil {
		fmt.Println("Ошибка запроса Select#Manager ")
	}

	Queries["Select#Posts"], err = db.Prepare(`SELECT * FROM "Posts"`)
	if err != nil {
		fmt.Println("Ошибка запроса Select#Posts ")
	}

}

func (ct *series) Select() error {
	stmt, ok := Queries["Select#Series"]
	if !ok {
		err = errors.New("db.go Select() - Select#Series")
		return err
	}

	rows, err := stmt.Query()
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&ct.Name)
		if err != nil {
			fmt.Println("Error - Scan", err.Error())
		}

		ct.Rows = append(ct.Rows, series{Name: ct.Name})
	}

	return nil
}

func (m *manager) Select() error {
	stmtUser, ok := Queries["Select#User"]
	if !ok {
		err = errors.New("db.go Select() - Select#User")
		return err
	}

	stmtManager, ok := Queries["Select#Manager"]
	if !ok {
		err = errors.New("db.go Select() - Select#Manager")
		return err
	}

	r := stmtUser.QueryRow(m.Login, m.Password)

	err := r.Scan(&m.Name)
	if err != nil {
		fmt.Println(err.Error())

		r = stmtManager.QueryRow(m.Login, m.Password)

		err = r.Scan(&m.Name, &m.Role)
		if err != nil {
			fmt.Println("Error - incorrect password or login", err.Error())
			return err
		}
	}
	return nil
}

func (posts *post) Select() error {
	stmtPost, ok := Queries["Select#Posts"]
	if !ok {
		err = errors.New("db.go Select() - Select#Posts")
		return err
	}
	// defer stmtPost.Close()

	rows, err := stmtPost.Query()
	if err != nil {
		return err
	}

	for rows.Next() {
		err := rows.Scan(&posts.Id, &posts.Image, &posts.Title, &posts.Text, &posts.Data)
		if err != nil {
			fmt.Println(err.Error())
		}

		posts.Rows = append(posts.Rows, post{
			Id:    posts.Id,
			Image: posts.Image,
			Title: posts.Title,
			Text:  posts.Text,
			Data:  posts.Data,
		})
	}

	return nil
}

func insertUsers(db *sql.DB, us user) error {
	query := `INSERT INTO "Users"("Email", "Password", "Name") VALUES ($1, $2, $3)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, us.Email, us.Password, us.Name)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

func deleteUsers(db *sql.DB, usb userBanned) error {
	res, err := db.Exec(`DELETE FROM "Users" WHERE "Email" = ($1)`, usb.Email)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			fmt.Println(count)
		}
		return nil
	}
	return err
}

func insertManager(db *sql.DB, m manager) error {
	query := `INSERT INTO "Manager"("Email", "Name", "Role", "Password") VALUES ($1, $2, $3, $4)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, m.Login, m.Name, m.Role, m.Password)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

func insertPosts(db *sql.DB, p post) error {
	query := `INSERT INTO "Posts"("Image", "Title", "Text", "Data") VALUES ($1, $2, $3, $4)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, p.Image, p.Title, p.Text, p.Data)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

func dbSelect() []user {
	rows, err := db.Query(`SELECT * FROM "Users"`)
	if err != nil {
		fmt.Println("Error = dbSelect() db.go")
		panic(err)
	}

	employee := user{}
	employees := []user{}

	for rows.Next() {
		var name, pass, email string
		err = rows.Scan(&email, &pass, &name)
		if err != nil {
			fmt.Println("Error = dbSelect() rows.Next()  db.go")
			panic(err)
		}
		employee.Email = email
		employee.Password = pass
		employee.Name = name
		employees = append(employees, employee)

	}
	//defer db.Close()
	return employees
}

func adminSelect() []manager {
	rows, err := db.Query(`SELECT * FROM "Manager"`)
	if err != nil {
		fmt.Println("Error = adminSelect() db.go")
		panic(err)
	}
	employee := manager{}
	employeesAdmin := []manager{}

	for rows.Next() {
		var email, name, role, pass string
		err = rows.Scan(&email, &name, &role, &pass)
		if err != nil {
			fmt.Println("Error = dbSelect() rows.Next()  db.go")
			panic(err)
		}
		employee.Login = email
		employee.Name = name
		employee.Role = role
		employee.Password = pass
		employeesAdmin = append(employeesAdmin, employee)

	}
	//defer db.Close()
	return employeesAdmin
}

func postSelect() []post {
	rows, err := db.Query(`SELECT * FROM "Posts"`)
	if err != nil {
		fmt.Println("Error = postSelect() db.go")
		panic(err)
	}
	employee := post{}
	employeesPost := []post{}

	for rows.Next() {
		var id int
		var image, title, text, data string
		err = rows.Scan(&id, &image, &title, &text, &data)
		if err != nil {
			fmt.Println("Error = postSelect() rows.Next()  db.go")
			panic(err)
		}
		employee.Id = id
		employee.Image = image
		employee.Title = title
		employee.Text = text
		employee.Data = data
		employeesPost = append(employeesPost, employee)

	}
	countPosts = len(employeesPost)
	//defer db.Close()
	return employeesPost
}

func productSelect() []product {
	rows, err := db.Query(`SELECT * FROM "Products"`)
	if err != nil {
		fmt.Println("Error = productSelect() db.go")
		panic(err)
	}
	employee := product{}
	employeesProduct := []product{}

	for rows.Next() {
		var article, count int
		var price float64
		var series, name, image, description string
		err = rows.Scan(&article, &series, &name, &price, &count, &image, &description)
		if err != nil {
			fmt.Println("Error = productSelect() rows.Next()  db.go")
			panic(err)
		}
		employee.Article = article
		employee.Series = series
		employee.Name = name
		employee.Price = price
		employee.Count = count
		employee.Image = image
		employee.Description = description
		employeesProduct = append(employeesProduct, employee)

	}
	//defer db.Close()
	return employeesProduct
}

var idPur int

func purchaseSelect() []purchase {
	rows, err := db.Query(`SELECT * FROM "Purchase"`)
	if err != nil {
		fmt.Println("Error = purchaseSelect() db.go")
		panic(err)
	}

	employee := purchase{}
	employeesPurchase := []purchase{}

	for rows.Next() {
		var count, product int
		var price, totalprice float64
		var user, date string
		err = rows.Scan(&idPur, &user, &product, &count, &price, &date, &totalprice)
		if err != nil {
			fmt.Println("Error = purchaseSelect() rows.Next()  db.go")
			panic(err)
		}
		employee.Id = idPur
		employee.User = user
		employee.Product = product
		employee.Count = count
		employee.Price = price
		employee.Date = date
		employee.TotalPrice = totalprice
		employeesPurchase = append(employeesPurchase, employee)

	}
	//defer db.Close()
	return employeesPurchase
}

func seriesSelect() []series {
	rows, err := db.Query(`SELECT * FROM "Series"`)
	if err != nil {
		fmt.Println("Error = seriesSelect() db.go")
		panic(err)
	}
	employee := series{}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Error = productSelect() rows.Next()  db.go")
			panic(err)
		}
		employee.Name = name
		employee.Rows = append(employee.Rows, employee)

	}
	//defer db.Close()
	return employee.Rows
}

func delUserSelect() []userBanned {
	rows, err := db.Query(`SELECT * FROM "DelUsers"`)
	if err != nil {
		fmt.Println("Error = seriesSelect() db.go")
		panic(err)
	}

	employeeUserBan := userBanned{}

	for rows.Next() {
		var email, reason string
		err = rows.Scan(&email, &reason)
		if err != nil {
			fmt.Println("Error = productSelect() rows.Next()  db.go")
			panic(err)
		}
		employeeUserBan.Email = email
		employeeUserBan.Reason = reason
		employeeUserBan.Rows = append(employeeUserBan.Rows, employeeUserBan)

	}
	//defer db.Close()

	return employeeUserBan.Rows
}

func insertSeries(db *sql.DB, s series) error {
	query := `INSERT INTO "Series"("Name") VALUES ($1)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.Name)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

func purSelect() pur {

	var stuff pur

	rowsProduct, err := db.Query(`SELECT * FROM "Products"`)
	if err != nil {
		fmt.Println("Error = purSelect() db.go")
		panic(err)
	}

	rowsUser, err := db.Query(`SELECT * FROM "Users"`)
	if err != nil {
		fmt.Println("Error = purSelect() db.go")
		panic(err)
	}

	for rowsProduct.Next() {
		var article, count int
		var price float64
		var series, name, image, description string
		err = rowsProduct.Scan(&article, &series, &name, &price, &count, &image, &description)
		if err != nil {
			fmt.Println("Error = productSelect() rows.Next()  db.go")
			panic(err)
		}
		stuff.Product.Article = article
		stuff.Product.Series = series
		stuff.Product.Name = name
		stuff.Product.Price = price
		stuff.Product.Count = count
		stuff.Product.Image = image
		stuff.Product.Description = description
		stuff.Product.Rows = append(stuff.Product.Rows, stuff.Product)
	}
	for rowsUser.Next() {
		var email, name, pass string
		err = rowsUser.Scan(&email, &pass, &name)
		if err != nil {
			fmt.Println("Error = dbSelect() rows.Next()  db.go")
			panic(err)
		}

		stuff.User.Email = email
		stuff.User.Password = pass
		stuff.User.Name = name

		stuff.User.Rows = append(stuff.User.Rows, stuff.User)
	}
	//defer db.Close()
	return stuff
}

func insertPurchase(db *sql.DB, f foo) error {
	query := `INSERT INTO "Purchase"("ID", "User", "Product", "Count", "Price", "Date", "TotalPrice") VALUES ($1, $2, $3, $4, $5, $6, $7)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}

	defer stmt.Close()
	f.id = 11
	f.userPur = "a.pavlikov22@gmail.com"
	f.products = 5435
	fmt.Println(f.id, f.userPur, f.products, f.valuePur, f.price, f.data, f.tprice)
	res, err := stmt.ExecContext(ctx, f.id, f.userPur, f.products, f.valuePur, f.price, f.data, f.tprice)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}

	queryRes := fmt.Sprintf(`UPDATE "Products" SET "Count" = "Count" - %d WHERE "Article" =  '%d'`, f.valuePur, f.products)
	_, er := db.Query(queryRes)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(queryRes)

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}

	log.Printf("%d purchase created ", rows)
	return nil
}

func insertDelUser(db *sql.DB, usb userBanned) error {
	query := `INSERT INTO "DelUsers"("Email", "Reason") VALUES ($1, $2)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, usb.Email, usb.Reason)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

func deleteSeries(db *sql.DB, s series) error {
	res, err := db.Exec(`DELETE FROM "Series" WHERE "Name" = ($1)`, s.Name)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			fmt.Println(count)
		}
		return nil
	}
	return err
}

func insertProduct(db *sql.DB, prod product) error {
	query := `INSERT INTO "Products"("Article", "Series", "Name", "Price", "Count", "Image", "Description") VALUES ($1, $2, $3, $4, $5, $6, $7)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	prod.Series = "Poco"
	res, err := stmt.ExecContext(ctx, prod.Article, prod.Series, prod.Name, prod.Price, prod.Count, prod.Image, prod.Description)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

func deletePost(db *sql.DB, pt post) error {
	res, err := db.Exec(`DELETE FROM "Posts" WHERE "Id" = ($1)`, pt.Id)
	if err == nil {
		_, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
		}
		return nil
	}
	return err
}

func deleteProduct(db *sql.DB, pd product) error {
	res, err := db.Exec(`DELETE FROM "Products" WHERE "Article" = ($1)`, pd.Article)
	if err == nil {
		_, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
		}
		return nil
	}
	return err
}

func brandSelect(brand string, val int) []product {

	text := fmt.Sprintf(`SELECT * FROM "Products" WHERE "Series" = '%s' LIMIT %d`, brand, val)

	rows, err := db.Query(text)
	if err != nil {
		fmt.Println("Error = brandSelect() db.go")
		panic(err)
	}

	employee := product{}
	employeesProduct := []product{}

	for rows.Next() {
		var article, count int
		var price float64
		var series, name, image, description string
		err = rows.Scan(&article, &series, &name, &price, &count, &image, &description)
		if err != nil {
			fmt.Println("Error = brandSelect() rows.Next()  db.go")
			panic(err)
		}
		employee.Article = article
		employee.Series = series
		employee.Name = name
		employee.Price = price
		employee.Count = count
		employee.Image = image
		employee.Description = description
		employeesProduct = append(employeesProduct, employee)

	}
	//defer db.Close()
	return employeesProduct
}

func valueBrandSelect(brand string) int {

	text := fmt.Sprintf(`SELECT COUNT(*) as count FROM "Products" WHERE "Series" = '%s'`, brand)

	rows, err := db.Query(text)
	if err != nil {
		fmt.Println("Error = brandSelect() db.go")
		panic(err)
	}
	var count int
	for rows.Next() {

		err = rows.Scan(&count)
		if err != nil {
			fmt.Println("Error = brandSelect() rows.Next()  db.go")
			panic(err)
		}
	}
	//defer db.Close()
	return count
}

func priceBrandSelect(brand string) []float64 {

	employeesPrice := []float64{}

	text := fmt.Sprintf(`SELECT "Price" as count FROM "Products" WHERE "Series" = '%s'`, brand)

	rows, err := db.Query(text)
	if err != nil {
		fmt.Println("Error = brandSelect() db.go")
		panic(err)
	}
	var count float64
	for rows.Next() {

		err = rows.Scan(&count)
		if err != nil {
			fmt.Println("Error = brandSelect() rows.Next() db.go")
			panic(err)
		}
		employeesPrice = append(employeesPrice, count)
	}
	//defer db.Close()
	return employeesPrice
}

func purchPriceSelect() ([]float64, []string) {
	employeesPrice := []float64{}
	employeesData := []string{}
	rows, err := db.Query(`SELECT "Date", "TotalPrice" as count FROM "Purchase"`)
	if err != nil {
		fmt.Println("Error = purchPriceSelect() db.go")
		panic(err)
	}
	var count float64
	var date string
	for rows.Next() {

		err = rows.Scan(&date, &count)
		if err != nil {
			fmt.Println("Error = brandSelect() rows.Next() db.go")
			panic(err)
		}
		employeesPrice = append(employeesPrice, count)
		employeesData = append(employeesData, date)
		//fmt.Println(count, date)
	}

	return employeesPrice, employeesData
}
