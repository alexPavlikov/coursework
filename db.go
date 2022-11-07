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

type post struct {
	Id    int
	Image string
	Title string
	Text  string
	Data  string

	Rows []post
}

var (
	db      *sql.DB
	err     error
	Queries map[string]*sql.Stmt
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
	Queries["Select#Category"], err = db.Prepare(`SELECT "Name" FROM "Category" ORDER BY "Name"`)
	if err != nil {
		fmt.Println("Ошибка запроса Select#Category ")
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

func (ct *categoty) Select() error {
	stmt, ok := Queries["Select#Category"]
	if !ok {
		err = errors.New("db.go Select() - Select#Category")
		return err
	}
	// defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&ct.Name)
		if err != nil {
			fmt.Println("Error - Scan", err.Error())
		}

		ct.Rows = append(ct.Rows, categoty{Name: ct.Name})
	}

	return nil
}

func (m *manager) Select() error {
	stmtUser, ok := Queries["Select#User"]
	if !ok {
		err = errors.New("db.go Select() - Select#User")
		return err
	}
	// defer stmtUser.Close()

	stmtManager, ok := Queries["Select#Manager"]
	if !ok {
		err = errors.New("db.go Select() - Select#Manager")
		return err
	}
	// defer stmtManager.Close()

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

	rows := stmtPost.QueryRow()
	if err != nil {
		return err
	}

	err := rows.Scan(&posts.Id, &posts.Image, &posts.Title, &posts.Text, &posts.Data)
	if err != nil {
		fmt.Println(err.Error())
	}

	// for rows.Next() {
	// 	err = rows.Scan(&posts.Image)
	// 	if err != nil {
	// 		fmt.Println("Error - Scan", err.Error())
	// 	}

	// 	posts.Rows = append(posts.Rows, post{Image: posts.Image})
	// }

	return nil
}

// Users insert in database
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

// Post insert in database
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

	// rows, ok := Queries["Select#Users"]
	// if !ok {
	// 	err = errors.New("db.go Select() - Select#User")
	// 	return err
	// }

	employee := user{}
	employees := []user{}

	for rows.Next() {
		var name, pass, email string
		err = rows.Scan(&email, &name, &pass)
		if err != nil {
			fmt.Println("Error = dbSelect() rows.Next()  db.go")
			panic(err)
		}
		employee.Email = email
		employee.Name = name
		employee.Password = pass
		employees = append(employees, employee)

	}
	// defer db.Close()
	return employees
}
