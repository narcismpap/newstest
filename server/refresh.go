// Package: com.github.narcismpap.news
// Async Refresh
//
// Author: Narcis M. Pap on 22/06/2019

package main

import (
	"fmt"
	"time"
)

// Setup and refresh all data sources on-the-fly
func SourcesTriggerRefresh() {
	// setup sources from config
	if !SharedSources.isReady {
		sLen := len(SharedConfig.Feeds)
		SharedSources.Sources = make([]Source, sLen)

		for i, src := range SharedConfig.Feeds {
			SharedSources.Sources[i] = Source{
				RSSFeed:  src.Source,
				Category: src.Category,
				Provider: src.Provider,
				Ready:    false,
			}
		}

		logMessage(fmt.Sprintf("Configured %d sources", len(SharedSources.Sources)))
		SharedSources.isReady = true
	}

	// let's reload our sources now
	for j := range SharedSources.Sources {
		// go's best capability is concurrency, let's use that to load all providers in parallel
		go func(pos int) {
			logMessage(fmt.Sprintf("Now refreshing: %s %s\n", SharedSources.Sources[pos].Provider, SharedSources.Sources[pos].Category))

			SharedSources.Sources[pos].RefreshMutex.Lock() // prevent race

			err := SharedSources.Sources[pos].Refresh()
			SharedSources.Sources[pos].RefreshTime = time.Now()
			SharedSources.Sources[pos].Ready = true

			if err != nil {
				// allow stale resources for one cycle
				SharedSources.Sources[pos].RefreshStatus = RefreshStatusError
				logError(err)
			} else {
				SharedSources.Sources[pos].RefreshStatus = RefreshStatusOK
			}

			SharedSources.Sources[pos].RefreshMutex.Unlock() // we're done
		}(j)
	}


}

// Triggers a refresh
func (s *Source) Refresh() error {
	if s.Provider == "BBC" { // BBC has a way around their media:thumbnail tags
		return ProviderBBCRefresh(s)
	}

	return ProviderGenericRefresh(s)
}
