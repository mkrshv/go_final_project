package main

import (
	"database/sql"
	todo "final_project/pkg"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = dbCheck()
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT, date TEXT, title TEXT, comment TEXT, repeat TEXT)")
	if err != nil {
		panic(err)
	}

	port := ":" + os.Getenv("TODO_PORT")
	if port == ":" {
		port = ":7540"
	}

	todo.StartServer(port)

}

// Проверяет наличие базы данных в директории выполнения приложения, создает БД если ее нет и возвращает путь до нее.
func dbCheck() string {
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	_, err = os.Stat(dbFile)

	if err != nil {
		os.Create("scheduler.db")
	}

	return dbFile
}
