// Package: com.github.narcismpap.news
// Response Test
//
// Author: Narcis M. Pap on 22/06/2019

package main

import (
	"testing"
)

func TestLoadRequestedArticles(t *testing.T) {
	response := LoadRequestedArticles(CategoriesAll, ProvidersAll)

	// these are some really silly tests
	if len(response.Categories) != len(SharedConfig.Categories){
		t.Errorf(".Categories mismatch")
	}

	if len(response.Providers) != len(SharedConfig.Providers){
		t.Errorf(".Providers mismatch")
	}
}
