// Package: com.github.narcismpap.news
// Configuration
//
// Author: Narcis M. Pap on 22/06/2019

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

// load config and perform sanity checks
func setupConfig() NewsConfig {
	var config NewsConfig

	// parse config
	configSrc, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic("unable to read config.json")
	}

	if err2 := json.Unmarshal(configSrc, &config); err2 != nil {
		panic("unable to parse config.json")
	}

	// sanity checks
	if len(config.Providers) == 0 {
		panic("no providers configured")
	}

	if len(config.Categories) == 0 {
		panic("no categories configured")
	}

	if len(config.Feeds) == 0 {
		panic("no feeds configured")
	}

	for pos, feed := range config.Feeds {
		if len(feed.Source) == 0 || !isValidUrl(feed.Source) {
			panic(fmt.Sprintf("Feed Source [%s] is not a valid URL", feed.Source))
		}

		if len(feed.Provider) == 0 {
			panic(fmt.Sprintf("Feed #%d has no provider", pos))
		}

		if len(feed.Category) == 0 {
			panic(fmt.Sprintf("Feed #%d has no category", pos))
		}
	}

	return config
}

// check if feed is valid URL
func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}
