package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/html"
	"strings"
)

type Page struct {
	linksTo []string
	reachedFrom []string
	assets []string
}

var (
	visited map[string]bool
	pages []Page
)

func main() {

	default_domain := "https://golang.org"
	domainp := flag.String("domain", default_domain, "Specify the domain name to crawl. Default is "+default_domain)
	flag.Parse()
	fmt.Println("Crawling domain", *domainp)
    startCrawl(*domainp)

}

func startCrawl(domain string) {
	links := getLinks(domain)
	for url, follow := range links {
		fmt.Println(url)
		if follow == true {
			links := getLinks(url)
		}
	}
}

func getLinks(domain string) (map[string]bool){
	links := make(map[string]bool)
	resp, err := http.Get(domain)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		t := tokenizer.Next()
		switch t {
		case html.ErrorToken:
			// End of document
			return links
		case html.StartTagToken:
			token := tokenizer.Token()

			isAnchor := token.Data == "a"
			if isAnchor {
				href := getHref(token)
				if href != "" && href != "/" && strings.HasPrefix(href, "#") != true {
					link, follow := normaliseHref(href, domain)
					links[link] = follow
				}
			}
		}
	}
}

func normaliseHref(href string, domain string) (string, bool) {
	newHref := href
	follow := true
	if strings.HasPrefix(href, "/") {
		newHref = domain + href
	} else if strings.HasPrefix(href, domain) != true {
		follow = false
	}
	return newHref, follow
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
