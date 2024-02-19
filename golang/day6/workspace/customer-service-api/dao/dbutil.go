package dao

import (
	"api/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Driver   string `json:"driver"`
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func connect() *sql.DB {
	file, err := os.Open("config.json")
	utils.CheckForError(err)

	var config Config
	json.NewDecoder(file).Decode(&config)

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User, config.Password, config.Hostname, config.Port, config.Database)

	db, err := sql.Open("mysql", connStr)
	utils.CheckForError(err)

	return db
}
