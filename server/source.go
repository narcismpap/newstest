// Package: com.github.narcismpap.news
// Data Sources
//
// Author: Narcis M. Pap on 22/06/2019

package main

import (
	"sync"
	"time"
)

const (
	RefreshStatusOK    = 1
	RefreshStatusError = 2

	CategoriesAll = SourceCategory("*")
	ProvidersAll  = SourceProvider("*")
)

type SourceProvider string // sanity
type SourceCategory string // sanity

// Holds a single parsed article
type SourceArticle struct {
	Title string
	Body  string
	Time  *time.Time

	Media string
	URL   string

	Category SourceCategory
	Provider SourceProvider
}

// Feed data source - holds global content
type Source struct {
	RSSFeed string
	Items   []*SourceArticle

	Category SourceCategory
	Provider SourceProvider

	RefreshMutex  sync.Mutex // prevents concurrency clashes
	RefreshTime   time.Time  // records last sync
	RefreshStatus int8       // status code
	Ready         bool       // false on init, true upon seeding
}
