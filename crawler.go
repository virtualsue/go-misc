package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
	"io"
)

func main() {

	default_domain := "https://golang.org"
	domainp := flag.String("domain", default_domain, "Specify the domain name to crawl. Default is "+default_domain)
	flag.Parse()
	fmt.Println("Crawling domain", *domainp)
    crawl(*domainp)

}

func crawl(domain string) {
	urls := make(map[string]bool)
	resp, err := http.Get(domain)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(page)
	for {
		t := tokenizer.Next()
		switch t {
		case html.ErrorToken:
			// End of document
			return
		case html.StartTagToken:
			token := tokenizer.Token()

			isAnchor := token.Data == "a"
			if isAnchor {
				href := getHref(token)
				fmt.Println("href is ", href)
				if (href != "") {
					// Determine if link should be followed
					crawl
				}
			}
		}
	}
}

func getHref(token html.Token) (href string) {
	href = ""
	for _, a := range token.Attr {
		if a.Key == "href" {
			href = a.Val
		}
	}
	return
}
