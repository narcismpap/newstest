// Package: com.github.narcismpap.news
// Base
//
// Author: Narcis M. Pap on 22/06/2019

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	RefreshWindow = time.Minute * 10 // our built-in async refresh
	ConfigFile    = "./config.json"  // relative to binary, panics if FnF
)

// Holds the newsfeeds
type NewsConfig struct {
	Categories []SourceCategory `json:"categories"`
	Providers  []SourceProvider `json:"providers"`

	Feeds []struct {
		Id       string         `json:"id"`
		Category SourceCategory `json:"category"`
		Provider SourceProvider `json:"provider"`
		Source   string         `json:"source"`
	} `json:"feeds"`
}

// Holds active data sources
type DataSources struct {
	isReady bool     // web server holds on false
	Sources []Source // holds a mutex-capable feed list, in-memory storage. resets upon service restart
}

var SharedConfig = setupConfig() // holds the config.json contents
var SharedSources = DataSources{isReady: false}

// HTTP REST Router service
func setupRequestRouting() {
	r := mux.NewRouter()

	r.HandleFunc("/all", ArticlesAllHandler)
	r.HandleFunc("/articles/{provider}/{category}", ArticlesFilterHandler)

	log.Fatal(http.ListenAndServe(":9000", r))
}

// Long-running async refresh mechanism (runs outside of the request-response cycle)
func setupRefresh() {
	go func() {
		for {
			SourcesTriggerRefresh()

			logMessage(fmt.Sprintf("async refresh wait, next at %s", (time.Now().Add(RefreshWindow)).Format("15:04:05.000")))
			time.Sleep(RefreshWindow)
		}

	}()
}

func main() {
	setupRefresh()
	setupRequestRouting()
}
