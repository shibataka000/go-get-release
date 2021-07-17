package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func newGoqueryDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return goquery.NewDocumentFromReader(res.Body)
}

func main() {
	version := "0.12.20"
	url := fmt.Sprintf("https://releases.hashicorp.com/terraform/%s", version)
	doc, err := newGoqueryDocument(url)
	if err != nil {
		panic(err)
	}
	assetNames := []string{}
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		assetName := s.Find("a").Text()
		assetNames = append(assetNames, assetName)
	})
}
