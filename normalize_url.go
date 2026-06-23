package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// func getHTML(rawURL string) (string, error) {

// }

func normalizeURL(rawUrl string) (string, error) {
	normalizeRegex := regexp.MustCompile(`^(?:https?://)?(.*?)/?$`)

	if !normalizeRegex.MatchString(rawUrl) {
		return "", errors.New("invalid url type")
	}

	p, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	return p.Host + p.Path, nil
}

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	sl := doc.Find("html")
	title := sl.Find("h1").Text()
	if title == "" {

		title = sl.Find("h2").Text()
	}

	return title

}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	sl := doc.Find("html")
	desc := sl.Find("main").Find("p").Text()
	if desc == "" {
		desc = sl.Find("p").Text()
	}

	return desc
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Println(err.Error())
		return []string{}, nil
	}

	body := doc.Find("body")
	urls := make([]string, 0)

	body.Find("a[href]").Each(func(_ int, link *goquery.Selection) {
		href, exists := link.Attr("href")
		if !exists {
			return
		}

		hrefURL, err := url.Parse(href)
		if err != nil {
			return
		}

		fullURL := baseURL.ResolveReference(hrefURL)

		urls = append(urls, fullURL.String())

	})

	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	urls := make([]string, 0)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Println(err.Error())
		return []string{}, nil
	}

	body := doc.Find("body")

	body.Find("img[src]").Each(func(_ int, img *goquery.Selection) {
		src, exists := img.Attr("src")
		if !exists {
			return
		}

		srcURL, err := url.Parse(src)
		if err != nil {
			return
		}

		fullURL := baseURL.ResolveReference(srcURL)
		urls = append(urls, fullURL.String())
	})

	return urls, nil

}

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	heading := getHeadingFromHTML(html)

	firstParagraph := getFirstParagraphFromHTML(html)

	baseParsedURL, err := url.Parse(pageURL)
	log.Println(baseParsedURL)
	if err != nil {
		fmt.Printf("error occurred %v", err)
		return PageData{}
	}
	outGoingLinks, err := getURLsFromHTML(html, baseParsedURL)
	if err != nil {
		fmt.Printf("error occurred %v", err)
		return PageData{}
	}

	imgUrls, err := getImagesFromHTML(html, baseParsedURL)
	if err != nil {
		fmt.Printf("error occurred %v", err)
		return PageData{}
	}

	return PageData{
		URL:            baseParsedURL.String(),
		Heading:        heading,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outGoingLinks,
		ImageURLs:      imgUrls,
	}
}
