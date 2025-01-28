package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gdinesh/websearch/userip"
)

func handleSearch(w http.ResponseWriter, req *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	query := req.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	userIP, err := userip.FromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx = userip.NewContext(ctx, userIP)

	start := time.Now()
	results, err := search.Search(ctx, query)
	elapsed := time.Since(start)

	if err := resultsTemplate.Execute(w, struct {
		Results          google.Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}
}
