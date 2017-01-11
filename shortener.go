package main

import (
	"errors"
	"math/rand"
	"time"
)

// Random string generation code from an excellent StackOverflow answer: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

const shortcodeLen = 7

const (
	lengthLessThanOneError = "string length must be greater than zero"
	linkEmptyError         = "URL cannot be empty"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

// randString returns a string of length n containing random characters selected from the letterBytes constant
func randString(n int) (string, error) {
	if n <= 0 {
		return "", errors.New(lengthLessThanOneError)
	}
	b := make([]byte, n)
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b), nil
}

// CreateLink creates a new link with shortcode from the passed long URL and description.
func CreateLink(url, desc string, categoryID int) (Link, error) {
	if url == "" {
		return Link{}, errors.New(linkEmptyError)
	}
	shortcode, err := randString(shortcodeLen)
	if err != nil {
		return Link{}, err
	}
	return Link{
		Shortcode:   shortcode,
		URL:         url,
		CategoryID:  categoryID,
		Description: desc,
	}, nil
}
