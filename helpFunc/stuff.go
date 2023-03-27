package helpFunc

import (
	"fmt"
	"math/rand"
	"net/http"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	_ "github.com/lib/pq"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// [longUrl]shortUrl
var AllUrl = make(map[string]string)

/*
создание нового Url заданной длины
По заданию длина должна быть как можно короче (минимальная длина =1), в программе используется
длина =6, чтобы избежать создания одинаковых url
*/
func shortenURL(length int, url string) string {
	shortURL := ""
	for i := 0; i < length; i++ {
		shortURL += string(letters[rand.Intn(len(letters))])
	}
	return url + shortURL
}

func NewUrl(longUrl string, url string) (string, int) {

	if !isURL(longUrl) {
		return "", 0
	}

	// проверяем, есть ли longUrl в базе данных
	testFlag := strings.Split(longUrl, " ")
	shortUrlString := testLongUrl(longUrl)
	if shortUrlString != "" {
		return shortUrlString, 2
	}

	shortUrlString = shortenURL(6, url)

	// проверяю есть ли флаг
	if testFlag[0] == "-d" {
		//записываю в бд
		longUrl = testFlag[1]
		PutSQL(shortUrlString, longUrl)

		prUrl := "localhost:8080"
		posrUrl := strings.TrimPrefix(shortUrlString, prUrl)

		CreateShortUrlHand(posrUrl, longUrl)

	} else {
		// записываю в память программы
		AllUrl[shortUrlString] = longUrl
	}
	return shortUrlString, 1
}

func testLongUrl(longUrl string) string {
	tmp := strings.Split(longUrl, " ")
	if len(tmp) == 2 {
		return FindShortSQL(tmp[1])
	}
	for key, val := range AllUrl {
		if val == longUrl {
			return key
		}
	}
	return ""
}

func isURL(longUrl string) bool {
	// Создание регулярного выражения для проверки веб-ссылки
	re := regexp.MustCompile(`^(http://|https://)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)
	if strings.Contains(longUrl, "-d") {
		if re.MatchString(strings.Split(longUrl, " ")[1]) {
			return true
		}
	} else {
		if re.MatchString(longUrl) {
			return true
		}
	}
	return false
}

// открываем ссылку на любой операционке
func OpenURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows", "darwin":
		err = exec.Command("cmd", "/c", "start", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

func CreateShortUrlHand(short, long string) {
	http.HandleFunc(short, func(w http.ResponseWriter, r *http.Request) {
		err := OpenURL(long)
		if err != nil {
			fmt.Println(err)
		}
	})
}
