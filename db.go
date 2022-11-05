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

	Queries["Select#Manager"], err = db.Prepare(`SELECT "Name", "Role" FROM "Manager" WHERE "Email"=$1 AND "Password"=$2`)
	if err != nil {
		fmt.Println("Ошибка запроса Select#Manager ")
	}

	// data := fmt.Sprintf(`INSERT INTO "Users" (Email, Password, Name) VALUES (%s, %s, %s)`, us.email, us.password, us.name)
	// fmt.Println(data)

}

func (ct *categoty) Select() error {
	stmt, ok := Queries["Select#Category"]
	if !ok {
		err = errors.New("db.go Select() - Select#Category")
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

// func (us *user) Insert() error {

// 	Queries["Insert#Users"], err = db.Prepare(`INSERT INTO "Users" (Email, Password, Name) VALUES ($1, $2, $3)`)
// 	if err != nil {
// 		fmt.Println("Ошибка запроса Insert#Users ")
// 	}

// 	stmtUsers, ok := Queries["Insert#Users"]
// 	if !ok {
// 		err = errors.New("db.go Insert() - Insert#Users")
// 		return err
// 	}

// 	_, er := stmtUsers.Exec(us.email, us.password, us.name)
// 	if er != nil {
// 		return err
// 	}

// 	return nil
// }

// INSERT INTO "Users"("Email", "Password", "Name") VALUES ('dl.donnu@gmail.com', 'idonotknow', 'University')

func insert(db *sql.DB, us user) error {
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
