package main

import (
	"FirstVerServer/handle"
	"FirstVerServer/helpFunc"

	"net/http"
)

func main() {
	helpFunc.Init()
	defer helpFunc.ClouseSQL()

	//создание хендлеров из бд
	go handle.CreateHandelSql()

	http.HandleFunc("/static/", handle.HandleStatic)
	http.HandleFunc("/res", handle.HandlePage2)
	http.HandleFunc("/", MainHandler)

	http.ListenAndServe(":8080", nil)
}
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// развертывание главной старницы
	http.ServeFile(w, r, "HTML/forms.html")
}
