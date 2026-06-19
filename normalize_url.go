package main

import (
	"errors"
	"net/url"
	"regexp"
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
