package main

import (
	"example/hello/crawler"
	"fmt"
	"sort"
)

func main() {
	// var wgForLinks sync.WaitGroup

	// var h1s []string
	// visitedURLs := crawler.NewLinkMap()
	// h1Ch := make(chan string)

	// baseURL := "http://books.toscrape.com/"

	// doc, err := crawler.GetDocumentFromURL(baseURL)
	// if err != nil {
	// 	log.Fatal("could not get root doc", err)
	// }

	// wgForLinks.Add(1)

	// // Next step: Extraction of H1s. Channel to collect H1s?
	// go crawler.ScrapeHTMLDocument(doc, visitedURLs, baseURL, &wgForLinks, h1Ch)

	// for {
	// 	newH1 := <-h1Ch
	// 	h1s = append(h1s, newH1)
	// 	fmt.Println("crawled", newH1, len(h1s))
	// }

	// wgForLinks.Wait()
	crawl("http://books.toscrape.com/", 32)
}

func crawl(baseURL string, concurrency int) {
	type urlResult struct {
		baseURL string
		urls    []string
	}
	type url string
	type job struct {
		url      url
		title    string
		complete bool
		err      error
	}
	jobs := map[url]*job{}

	addJob := func(newURL string) {
		jobs[url(newURL)] = &job{
			url:      url(newURL),
			complete: false,
		}
	}

	addJob(baseURL)

	chanJobComplete := make(chan job)
	chanJobNewUrls := make(chan urlResult)

	running := 0

	for {
		complete := 0
		for _, nextJob := range jobs {
			if nextJob.complete {
				complete++
			} else if running < concurrency {
				running++
				go func(currentJob job) {
					title, urls, err := crawler.ScrapeHTMLDocument(string(currentJob.url))
					currentJob.complete = true
					currentJob.err = err
					currentJob.title = title
					chanJobNewUrls <- urlResult{
						baseURL: string(currentJob.url),
						urls:    urls,
					}
					chanJobComplete <- currentJob
				}(*nextJob)
			}
		}
		if complete == len(jobs) {
			fmt.Println("we are done")
			// spew.Dump(jobs)
			stringURLs := []string{}
			for u := range jobs {
				stringURLs = append(stringURLs, string(u)) // string(u) -> casting to string?
			}
			sort.Strings(stringURLs)
			for i, stringURL := range stringURLs {
				job := jobs[url(stringURL)]
				fmt.Println(i, stringURL, "	", job.title, "	", job.err)
			}
			return
		}
		select {
		case newURLsResults := <-chanJobNewUrls:
			for _, newURL := range newURLsResults.urls {
				newURL, err := crawler.NormalizeLink(newURLsResults.baseURL, newURL)
				if err != nil {
					fmt.Println("could not normalize url", err)
					continue
				}
				_, exists := jobs[url(newURL)]
				if !exists {
					addJob(newURL)
				}
			}
		case completedJob := <-chanJobComplete:
			running--
			jobs[completedJob.url] = &completedJob
		}

	}
}
