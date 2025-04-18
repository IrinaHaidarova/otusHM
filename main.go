package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет Otus!"))
	w.WriteHeader(http.StatusOK) // одно и тоже что и 200
	w.WriteHeader(200)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed) // вернем 405
		return
	}
	w.Write([]byte("Привет!"))
	w.WriteHeader(http.StatusOK) // одно и тоже что и 200
}

func main() {
	http.HandleFunc("/health/", handler)
	http.HandleFunc("/2", handler2)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
