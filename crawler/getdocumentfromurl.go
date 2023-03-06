package crawler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetDocumentFromURL loads and returns document from URL
func GetDocumentFromURL(url string) (*goquery.Document, error) {
	fmt.Printf("Getting doc from: %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		//		log.Fatalf()
		return nil, fmt.Errorf("unexpected status code %d: %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
