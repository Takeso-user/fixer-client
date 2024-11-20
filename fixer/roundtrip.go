package fixer

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type LogRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l LogRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	_, err := fmt.Fprintf(l.logger, "[%s]  %s  %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	if err != nil {
		log.Fatal(err)
	}
	return l.next.RoundTrip(r)
}
