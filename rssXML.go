package gorss

import (
	"encoding/xml"
	// "fmt"
)

type RssOut struct {
    XMLName xml.Name `xml:"rss"`
    Version string   `xml:"version,attr"`
    Channel *RssFeed
}

type RssFeed struct {
    XMLName xml.Name `xml:"channel"`
    Title string `xml:"title"`
    Link string `xml:"link"`
    Description string `xml:"description"`
    Items []*RssItem
}

type RssItem struct {
	XMLName xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
}

func newItem(post *Item) *RssItem {
	item := &RssItem {
		Title : post.Title,
		Link : post.Link,
		Description: post.Description,
	}
	return item
}

type Rss struct{
	*Feed
}

func (r *Rss) CreateFeed() *RssFeed {
	channel := &RssFeed {
		Title : "Big K's blog",
		Link : "http://bigk.me",
		Description: "Big K's blog",
	}

	for _, i := range r.Items {
		channel.Items = append(channel.Items, newItem(i))
	}
	return channel
}

func (r *Rss) FeedXml() *RssOut {
	return r.CreateFeed().FeedXml()

}

func (r *RssFeed) FeedXml() *RssOut {
	return &RssOut{Version: "2.0", Channel: r}
}