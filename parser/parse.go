package parser

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Parse parses a document with goquery and returns a element
func Parse(doc *goquery.Document, filters []string) (parsed string) {
	for _, filter := range filters {
		doc.Find(filter).Each(func(i int, s *goquery.Selection) {
			if filter == "body > pre" {
				parsed = retriveDescriptionData(doc, filter)
			}
			if _, err := strconv.Atoi(s.Text()); err == nil {
				parsed = s.Text()
			}
		})
	}
	return
}

func retriveDescriptionData(doc *goquery.Document, filter string) string {
	return retriveLastWord(retriveDescriptionFromHTML(doc, filter))
}

func retriveDescriptionFromHTML(doc *goquery.Document, filter string) string {
	return strings.TrimSpace(doc.Find(filter).Contents().Get(0).Data)
}

func retriveLastWord(words string) string {
	wordsInArray := strings.Split(words, " ")
	return wordsInArray[len(wordsInArray)-1]
}
