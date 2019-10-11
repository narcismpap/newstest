// Package: com.github.narcismpap.news
// Generic RSS Parser
//
// Author: Narcis M. Pap on 22/06/2019

package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

// Generic RSS importer, using mmcdole/gofeed
func ProviderGenericRefresh(s *Source) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(s.RSSFeed)

	if err != nil {
		return err
	}

	articles := make([]*SourceArticle, len(feed.Items))

	for i, item := range feed.Items {
		image := ""
		if item.Image != nil {
			image = item.Image.URL
		}

		articles[i] = &SourceArticle{
			Title:    item.Title,
			Body:     item.Description,
			Time:     item.PublishedParsed,
			Media:    image,
			URL:      item.Link,
			Category: s.Category,
			Provider: s.Provider,
		}

		logMessage(fmt.Sprintf("[+] parser_generic [%v] from %s\n", item.Title, articles[i].Provider))
	}

	s.Items = articles
	return nil
}
