package gorss

import (
	"encoding/xml"
)

type Feed struct {
	Title       string
	Link        string
	Description string
	Items       []*Item
}

type Item struct {
	Title       string
	Link        string
	Description string
}

func (f *Feed) Xml() (string, error) {
	r := &Rss{f}
	x := r.FeedXml()
	data, err := xml.MarshalIndent(x, "", " ") 
	if err != nil {
		return "", err
	}

	s := xml.Header[:len(xml.Header)-1] + string(data)
	return s, nil
}

func (f *Feed) AddItem(item *Item) {
	f.Items = append(f.Items,item)
}