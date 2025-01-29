package search

import (
	"context"
	"net/http"
)

func Search(ctx context.Context, query string) (Results, error) {
	req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
	if err != nil {
		return nil, err
	}
	//ToDo
}
