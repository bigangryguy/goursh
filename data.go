package main

import "time"

// Category represents a hierarchical link category.
type Category struct {
	ID          int       `json:"id"`
	ParentID    int       `json:"parentid"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
}

// Link represents an original URL and shortcode to use with a URL shortening service.
type Link struct {
	Shortcode   string    `json:"shortcode"`
	URL         string    `json:"url"`
	CategoryID  int       `json:"categoryid"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
}

// LinkVisit represents a single visit to a link.
type LinkVisit struct {
	Shortcode string    `json:"shortcode"`
	Visit     time.Time `json:"visit"`
}
