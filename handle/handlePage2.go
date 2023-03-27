package handle

import (
	"FirstVerServer/helpFunc"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// кнопка перехода на вторую страницу
func HandlePage2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/res" {
		http.Error(w, "404 page TWO not found", http.StatusNotFound)
		return
	}
	drowTableMap(w)
	drowTableSQL(w)
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "HTML/formsResult.html")

	default:
		fmt.Fprintf(w, "Only get")
	}

}

// кнопка перехода на главную страницу
func HandleReturnPage2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page TWO not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		fmt.Fprintf(w, "Only get")
	}

}

var tmpl = `
        <table cellspacing="15">
		<thead><th>Short Url</th><th>Long Url</th></thead>
            {{range $key, $value := .}}
                <tr>
                    <td>{{$key}}</td> 
					
                    <td>{{$value}}</td>
                </tr>
            {{end}}
        </table>
		<br>
    `

func drowTableMap(w http.ResponseWriter) {
	// Парсим шаблон
	t := template.Must(template.New("").Parse(tmpl))

	// Выводим содержимое map в виде HTML-таблицы
	err := t.Execute(w, helpFunc.AllUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// вывод данных из бд
func drowTableSQL(w http.ResponseWriter) {
	// Парсим шаблон
	t := template.Must(template.New("").Parse(tmpl))

	rows := helpFunc.GetSQL()
	defer helpFunc.CloseGetSql(rows)

	var urlTmp = make(map[string]string)
	for rows.Next() {
		var short, long string
		if err := rows.Scan(&short, &long); err != nil {
			log.Printf("Error scanning row: %v", err)
		}
		urlTmp[short] = long
	}
	// Выводим содержимое map в виде HTML-таблицы
	err := t.Execute(w, urlTmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
