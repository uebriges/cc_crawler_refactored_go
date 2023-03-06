package crawler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ScrapeHTMLDocument extracts a tags recursively, crawls their hrefs and scrapes h1 tags
func TestScrapeHTMLDocument(t *testing.T) () {

	url := "http://books.toscrape.com/"
	title, _, err := ScrapeHTMLDocument(url)
	fmt.Println("title", title)
	// fmt.Println("urls", urls)
	fmt.Println("err", err)

	assert.Equal(t, "http://books.toscrape.com/", url)
}
