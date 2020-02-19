package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// var washpostXML []byte = (`
// <sitemapindex>
// 	<sitemap>
// 		<loc>
// 			https://www.washingtonpost.com/news-sitemaps/goingoutguide.xml
// 		</loc>
// 	</sitemap>
// 	<sitemap>
// 		<loc>
// 			https://www.washingtonpost.com/news-sitemaps/goingoutguide.xml
// 		</loc>
// 	</sitemap>
// </sitemapindex>
// `)

// Location : some text
type Location struct {
	Loc string `xml:"loc"`
}

// SitemapIndex : some text 2
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if err != nil {
		print(err)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	// fmt.Println(s.Locations)
	for _, Location := range s.Locations {
		fmt.Println(strings.TrimSpace(Location.Loc))
	}
}
