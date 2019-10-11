// Package: com.github.narcismpap.news
// BBC-specific Parser
//
// Author: Narcis M. Pap on 22/06/2019

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BBCRss struct {
	XMLName xml.Name   `xml:"rss"`
	Channel BBCChannel `xml:"channel"`
}

type BBCChannel struct {
	Items []BBCItem `xml:"item"`
}

type BBCImage struct {
	XMLName xml.Name `xml:"image"`
	Url     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Width   string   `xml:"width"`
	Height  string   `xml:"height"`
}

type BBCItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Media       BBCMedia `xml:"thumbnail"`
}

type BBCMedia struct {
	XMLName xml.Name `xml:"thumbnail"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Url     string   `xml:"url,attr"`
}

func ProviderLoadContents(url string) ([]byte, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { // without anon func we can miss this err
		err := resp.Body.Close()
		if err != nil {
			logError(err)
		}
	}()

	response, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return response, err
}

// BBC Importer
func ProviderBBCRefresh(s *Source) error {
	var bbcFeed BBCRss

	rep, err := ProviderLoadContents(s.RSSFeed)
	if err != nil {
		return err
	}

	err2 := xml.Unmarshal(rep, &bbcFeed)
	if err2 != nil {
		return err2
	}

	articles := make([]*SourceArticle, len(bbcFeed.Channel.Items))

	for i, item := range bbcFeed.Channel.Items {
		pubTime, err := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", item.PubDate)
		if err != nil {
			logError(err)
			continue // skip
		}

		articles[i] = &SourceArticle{
			Title: item.Title,
			Body:  item.Description,
			Time:  &pubTime,
			Media: item.Media.Url,
			URL:   item.Link,

			Category: s.Category,
			Provider: s.Provider,
		}

		logMessage(fmt.Sprintf("[+] parser_bbc [%v] from %s\n", item.Title, articles[i].Provider))
	}

	s.Items = articles
	return nil
}
