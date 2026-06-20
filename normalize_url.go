package main

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func normalizeURL(rawUrl string) (string, error) {
	normalizeRegex := regexp.MustCompile(`^(?:https?://)?(.*?)/?$`)

	if !normalizeRegex.MatchString(rawUrl) {
		return "", errors.New("invalid url type")
	}

	p, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	// fmt.Printf("raw %v\n", p.Path)
	// fmt.Printf("raw %v\n", p.Host)
	// fmt.Printf("raw combined %v%v\n", p.Host, p.Path)

	// cleaned := normalizeRegex.ReplaceAllString(rawUrl, "$1")

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
	urls := make([]string, 0)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Println(err.Error())
		return []string{}, nil
	}

	body := doc.Find("body")

	body.Find("a[href]").Each(func(_ int, link *goquery.Selection) {
		href, exists := link.Attr("href")
		if exists {
			urls = append(urls, href)

		}

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
		if exists {
			fullImgLink := baseURL.Scheme + `://` + baseURL.Host + src
			urls = append(urls, fullImgLink)
		}
	})

	return urls, nil

}
