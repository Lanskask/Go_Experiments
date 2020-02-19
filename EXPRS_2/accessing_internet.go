package main

/*
<sitemapindex>
	<sitemap>
		<loc>
			https://www.washingtonpost.com/news-sitemaps/goingoutguide.xml
		</loc>
	</sitemap>
	<sitemap>
		<loc>
			https://www.washingtonpost.com/news-sitemaps/goingoutguide.xml
		</loc>
	</sitemap>
</sitemapindex>
*/

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type Location struct {
	Loc string `xml:"loc"`
}

type SitemapIndex struct {
	Locations []location `xml:"sitemap"`
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
}
