package handle

import (
	"FirstVerServer/helpFunc"
	"fmt"
	"log"
	"strings"
)

func CreateHandelSql() {

	rows := helpFunc.GetSQL()
	defer helpFunc.CloseGetSql(rows)

	fmt.Printf("Данные которые добавили из бд\n")

	for rows.Next() {
		var short, long string
		if err := rows.Scan(&short, &long); err != nil {
			log.Printf("Error scanning row(first): %v", err)
		}
		prUrl := "localhost:8080"
		posrUrl := strings.TrimPrefix(short, prUrl)
		helpFunc.CreateShortUrlHand(posrUrl, long)

		fmt.Printf("Short: %s, Long: %s\n", short, long)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}
}
