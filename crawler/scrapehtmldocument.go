package crawler

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeHTMLDocument extracts a tags recursively, crawls their hrefs and scrapes h1 tags
func ScrapeHTMLDocument(
	//document *goquery.Document,
	baseURL string,
	// docURL string,
) (title string, urls []string, err error) {

	document, err := GetDocumentFromURL(baseURL)
	if err != nil {
		return "", nil, err
	}

	aTags := document.Find("a")

	//scrapeErrors := []error{}

	aTags.Each(func(i int, aTag *goquery.Selection) {
		href, hasHref := aTag.Attr("href")

		if hasHref {
			fmt.Println("href", href)
			urls = append(urls, href)
		}

		// normalizedLink, err := normalizeLink(baseURL, href)
		// if err != nil {
		// 	scrapeErrors = append(scrapeErrors, err)
		// 	return
		// }

		//if hasHref && !linkMap.containsLink(normalizedLink) {

		// Add link to linkMap
		//			linkMap.add(normalizedLink)

		// currentDoc, err := GetDocumentFromURL(normalizedLink)
		// if err != nil {
		// 	//return err
		// 	scrapeErrors = append(scrapeErrors, err)
		// 	return
		// }
		//h1s := extractElementsFromDocument(currentDoc, "h1")

		// for _, h1 := range h1s {
		// 	fmt.Printf("Add h1: %s\n", h1)
		// 	// h1Ch <- h1 // channel needs to be closed at some point <- happens outside
		// }

		// } else {
		// 	fmt.Println("skipping", normalizedLink)
		// }
	})
	h1s := extractElementsFromDocument(document, "h1")
	if len(h1s) > 0 {
		title = h1s[0]
	}

	return title, urls, err
}
