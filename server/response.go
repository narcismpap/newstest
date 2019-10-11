// Package: com.github.narcismpap.news
// Response Manager
//
// Author: Narcis M. Pap on 22/06/2019

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

// Loads all relevant articles based on filtering criteria
func LoadRequestedArticles(category SourceCategory, provider SourceProvider) NewsListResponse {
	articles := make([]*SourceArticle, 0)

	for sID := range SharedSources.Sources {
		if SharedSources.Sources[sID].Ready {
			if category != CategoriesAll && category != SharedSources.Sources[sID].Category {
				continue
			}

			if provider != ProvidersAll && provider != SharedSources.Sources[sID].Provider {
				continue
			}

			for _, art := range SharedSources.Sources[sID].Items {
				articles = append(articles, art)
			}
		}
	}

	// order results by time
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Time.Unix() > articles[j].Time.Unix()
	})

	return NewsListResponse{
		Articles:   articles,
		Providers:  SharedConfig.Providers,
		Categories: SharedConfig.Categories,
	}
}

// Renders 500 error
func prepareServerErrorResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
}

// Renders Error Response
func prepareRequestErrorResponse(w http.ResponseWriter, reason string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	logMessage(fmt.Sprintf("--> user error: %s", reason))

	response, err := json.Marshal(map[string]string{"error": reason})
	if err != nil {
		logError(err)
		prepareServerErrorResponse(w)
		return
	}

	_, err2 := w.Write(response)

	if err2 != nil {
		logError(err2)
		prepareServerErrorResponse(w)
		return
	}
}

// Renders JSON response
func prepareListResponse(w http.ResponseWriter, code int, articles NewsListResponse) {
	response, err := json.Marshal(articles)

	if err != nil {
		logError(err)
		prepareServerErrorResponse(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err2 := w.Write(response)

	if err2 != nil {
		logError(err2)
		prepareServerErrorResponse(w)
		return
	}
}
