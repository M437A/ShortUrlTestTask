package handle

import (
	"FirstVerServer/helpFunc"
	"fmt"
	"net/http"
	"strings"
)

var url string

// кнопки добавления нового url и поиска готового
func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/static/" {
		http.Error(w, "404 page ONE not found", http.StatusNotFound)
		return
	}
	//префикс url
	url = r.Host + "/"

	switch r.Method {
	case "GET":

		shortUrl := r.FormValue("shortUrl")

		fmt.Fprintf(w, "Short url: %s", shortUrl)

		if shortUrl == "" {
			http.Error(w, "Вы ничего не ввели", http.StatusBadRequest)
			return
		}

		//ищем длинную
		longUrl := helpFunc.AllUrl[shortUrl]

		if longUrl == "" {
			longUrl = helpFunc.FindLongSQL(shortUrl)
		}
		if longUrl == "" {
			fmt.Fprintf(w, "\nURL адрес не найден")
		} else {

			err := helpFunc.OpenURL(longUrl)
			if err != nil {
				fmt.Println(err)
			}

		}

	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "PorseForm err - %v", err)
			return
		}
		//считываем ссылку
		longUrl := r.FormValue("longUrl")
		//попробовать сделать так, чтобы NewUrl ничего не возврашала
		//получаем новую ссылку типа: localhost:8080/dWzkmq и добавление в базу либо в память, создает нужные хэндлеры
		webCode, b := helpFunc.NewUrl(longUrl, url)
		// убираем localhost:8080
		strShort := postfUrl(webCode)
		switch b {
		case 1: //новый url
			helpFunc.CreateShortUrlHand(strShort, longUrl)

		case 2: //существуюший URL
			http.Redirect(w, r, strShort, http.StatusSeeOther)

		default:
			fmt.Fprintf(w, "Вы ввели не правлильную ссылка, она не будет сохраненна\n")
		}
		fmt.Fprintf(w, "New URL- %s", webCode)

	default:

		fmt.Fprintf(w, "Only get and post")
	}

}

func postfUrl(short string) string {
	return "/" + strings.TrimPrefix(short, url)
}
