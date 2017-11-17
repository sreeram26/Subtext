package transliterate

import (
	"fmt"
	"net/http"
)

type SubText struct {
	englishSubText  string
	language        string
	languageSubText string
}

func ConvertEngSubText(englishSubText string) string {
	return englishSubText
}

func HandleTransliterateQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintf(w, "In Progress!! Wait")
}
