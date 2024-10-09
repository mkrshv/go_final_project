package main

import (
	"fmt"
	"net/http"
	"time"
)

func StartServer(port string) {
	webDir := "./web"

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	http.HandleFunc("/api/nextdate", handleNextDate)

	http.HandleFunc()

	fmt.Println("Server starting at", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

func handleNextDate(w http.ResponseWriter, r *http.Request) {
	now, _ := time.Parse("20060102", r.FormValue("now"))
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	fmt.Println(date, repeat)
	newDate, err := NextDate(now, date, repeat)
	fmt.Println(newDate, err)
	w.Write([]byte(newDate))
}

func handleTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

	}
}
