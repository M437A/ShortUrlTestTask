package helpFunc

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

func createSQL() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func ClouseSQL() {
	Db.Close()
}

func Init() {
	var err error
	Db, err = createSQL()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}

// функция добавления в бд
func PutSQL(short, long string) {
	_, err := Db.Exec("INSERT INTO public.url_2(short, long) VALUES ($1, $2)", short, long)
	if err != nil {
		log.Printf("error inserting into database: %v", err)
	}

}

// получение всех данных из sql
func GetSQL() *sql.Rows {
	rows, err := Db.Query("SELECT short, long FROM public.url_2")
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}

	return rows
}

func CloseGetSql(rows *sql.Rows) {
	defer rows.Close()
}

// ищем короткую ссылку по длинной
func FindShortSQL(longUrl string) string {
	rows, err := Db.Query("SELECT short FROM public.url_2 WHERE long = $1", longUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Printf("К данной длинной ссылке нет короткой ссылки \nlong=%s\n", longUrl)
		return ""
	}

	var shortUrl string
	err = rows.Scan(&shortUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Короткая ссылка найдена в базе, она равна ", shortUrl)
	return shortUrl
}

func FindLongSQL(shortURL string) string {
	rows, err := Db.Query("SELECT long FROM public.url_2 WHERE short = $1", shortURL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Printf("К данной короткой ссылке нет длинной ссылке \nshort=%s", shortURL)
		return ""
	}

	var longURL string
	err = rows.Scan(&longURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Длинная ссылка найдена в базе, она равна ", longURL)
	return longURL
}
