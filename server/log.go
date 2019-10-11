// Package: com.github.narcismpap.news
// Logging
//
// Author: Narcis M. Pap on 22/06/2019

package main

import (
	"log"
	"time"
)

func logError(e error) {
	// log to Sentry in production
	log.Printf("[ERROR] %v", e)
}

func logMessage(msg string) {
	log.Printf("%s: %s", time.Now().Format("2006-01-02 15:04:05.000"), msg)
}
