// Package: com.github.narcismpap.news
// Views
//
// Author: Narcis M. Pap on 22/06/2019

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type NewsListResponse struct {
	Articles   []*SourceArticle
	Categories []SourceCategory
	Providers  []SourceProvider
}

func ArticlesFilterHandler(w http.ResponseWriter, r *http.Request) {
	urlVars := mux.Vars(r)

	// sanity checks
	if _, ok := urlVars["category"]; !ok {
		prepareRequestErrorResponse(w, "category is missing")
		return
	}

	if _, ok := urlVars["provider"]; !ok {
		prepareRequestErrorResponse(w, "provider is missing")
		return
	}

	allArticles := LoadRequestedArticles(SourceCategory(urlVars["category"]), SourceProvider(urlVars["provider"]))

	logMessage(fmt.Sprintf("--> /articles/%s/%s with %d results", urlVars["provider"], urlVars["category"], len(allArticles.Articles)))
	prepareListResponse(w, http.StatusOK, allArticles)
}

func ArticlesAllHandler(w http.ResponseWriter, r *http.Request) {
	allArticles := LoadRequestedArticles(CategoriesAll, ProvidersAll)

	logMessage(fmt.Sprintf("--> /all with %d results", len(allArticles.Articles)))
	prepareListResponse(w, http.StatusOK, allArticles)
}
