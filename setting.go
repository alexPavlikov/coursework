package main

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
)

type Setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgName     string
	Data       string
	Assets     string
	Html       string
}

var cfg Setting

func init() {
	//Открыть файл
	file, err := os.Open("setting.cfg")
	if err != nil {
		log.Fatal("Error - Открыть файл ", err.Error())
	}
	defer file.Close()

	var stat fs.FileInfo
	stat, err = file.Stat()
	if err != nil {
		log.Fatal("Error - Статистика файла ", err.Error())
	}

	reabByte := make([]byte, stat.Size())

	_, err = file.Read(reabByte)
	if err != nil {
		log.Fatal("Error - Чтение файла ", err.Error())
	}

	err = json.Unmarshal(reabByte, &cfg)
	if err != nil {
		log.Fatal("Error - Unmarshal файла ", err.Error())
	}

}
