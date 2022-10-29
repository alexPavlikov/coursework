package main

import "fmt"

type Setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgName     string
}

func init() {
	fmt.Println("Hello setting")

}
