package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type categoty struct {
	Name string
	Rows []categoty
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
