package crawler

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeHTMLDocument extracts a tags recursively, crawls their hrefs and scrapes h1 tags
func ScrapeHTMLDocument(
	baseURL string,
) (title string, urls []string, err error) {

	document, err := GetDocumentFromURL(baseURL)
	if err != nil {
		return "", nil, err
	}

	aTags := document.Find("a")

	aTags.Each(func(i int, aTag *goquery.Selection) {
		href, hasHref := aTag.Attr("href")

		if hasHref {
			fmt.Println("href", href)
			urls = append(urls, href)
		}
	})

	h1s := extractElementsFromDocument(document, "h1")

	if len(h1s) > 0 {
		title = h1s[0]
	}

	return title, urls, err
}
