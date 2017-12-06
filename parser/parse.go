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
				t := strings.TrimSpace(doc.Find(filter).Contents().Get(2).Data)
				parsed = t
			}
			if _, err := strconv.Atoi(s.Text()); err == nil {
				parsed = s.Text()
			}
		})
	}
	return
}
