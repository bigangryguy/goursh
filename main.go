package main

import (
	"net/http"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	ctx := context.Background()

	logger := log.New()

	var svc ShortenerService
	svc = shortenerService{}
	svc = loggingMiddleware{
		logger: logger,
		next:   svc,
	}

	shortenHandler := httptransport.NewServer(
		ctx,
		makeShortenEndpoint(svc),
		decodeShortenRequest,
		encodeResponse,
	)

	http.Handle("/shorten", shortenHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
