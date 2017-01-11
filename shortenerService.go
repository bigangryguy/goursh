package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/go-kit/kit/endpoint"
)

const (
	urlEmptyError = "URL cannot be empty"
)

// ShortenerService provides operations for shortening and managing URLs.
type ShortenerService interface {
	Shorten(shortenRequest) (Link, error)
}

type loggingMiddleware struct {
	logger *log.Logger
	next   ShortenerService
}

type shortenerService struct{}

type shortenRequest struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	CategoryID  int    `json:"categoryid"`
}

type shortenResponse struct {
	Link  Link  `json:"link"`
	Error error `json:"error"`
}

func (shortenerService) Shorten(req shortenRequest) (Link, error) {
	if req.URL == "" {
		return Link{}, errors.New(urlEmptyError)
	}

	return CreateLink(req.URL, req.Description, req.CategoryID)
}

func (mw loggingMiddleware) Shorten(req shortenRequest) (link Link, err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(log.Fields{
			"method": "Shorten",
			"input":  req,
			"output": link,
			"error":  err,
			"took":   time.Since(begin),
		}).Info("Shorten service request")
	}(time.Now())

	link, err = mw.next.Shorten(req)
	return
}

func makeShortenEndpoint(svc ShortenerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(shortenRequest)
		l, err := svc.Shorten(req)
		return shortenResponse{
			Link:  l,
			Error: err,
		}, nil
	}
}

func decodeShortenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
