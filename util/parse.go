package util

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func Parse(doc *goquery.Document, filters []string) (parsed string) {
	for _, filter := range filters {
		doc.Find(filter).Each(func(i int, s *goquery.Selection) {
			if _, err := strconv.Atoi(s.Text()); err == nil {
				parsed = s.Text()
			}
		})
	}
	return
}
