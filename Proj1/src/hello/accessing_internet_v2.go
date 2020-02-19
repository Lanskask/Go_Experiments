package main

import (
	"encoding/json"
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

var washPostXMLNews = (`
<urlset>
	<url>
		<loc>
			https://www.washingtonpost.com/politics/trump-kicks-off-daytona-500-as-grand-marshal/2020/02/16/53bb323e-50d5-11ea-9e47-59804be1dcfb_story.html
		</loc>
		<lastmod>2020-02-16T20:31:57Z</lastmod>
		<news:news>
			<news:publication>
				<news:name>Washington Post</news:name>
				<news:language>en</news:language>
			</news:publication>
			<news:publication_date>2020-02-16T20:31:57Z</news:publication_date>
				<news:title>
					<![CDATA[ Trump kicks off Daytona 500 as grand marshal ]]>
				</news:title>
				<news:keywords>
					<![CDATA[
					, trump, the great american race, rv, keep america great, air force one
					]]>
				</news:keywords>
		</news:news>
		<changefreq>hourly</changefreq>
	</url>
	<url>
			<loc>
				https://www.washingtonpost.com/politics/trumps-soft-touch-with-chinas-xi-worries-advisers-who-say-more-is-needed-to-combat-coronavirus-outbreak/2020/02/16/93de385a-5019-11ea-9b5c-eac5b16dafaa_story.html
			</loc>
			<lastmod>2020-02-16T19:33:00Z</lastmod>

			<news:news>
				<news:publication>
					<news:name>Washington Post</news:name>
					<news:language>en</news:language>
				</news:publication>
				<news:publication_date>2020-02-16T19:33:00Z</news:publication_date>
				<news:title>
					<![CDATA[
					Trump’s soft touch with China’s Xi worries advisers who say more is needed to combat coronavirus outbreak
					]]>
				</news:title>
				<news:keywords>
					<![CDATA[
					coronavirus, trump, xi, outbreak, Fauci, Azar, China
					]]>
				</news:keywords>
			</news:news>

			<changefreq>hourly</changefreq>
	</url>
</urlset>
`)

type URLSet struct {
	XMLName xml.Name `xml:"urlset" json:"-"`
	URLList []Url2   `xml:"url" json="url"`
}

type Url2 struct {
	// XMLName         xml.Name        `xml:"url"  json="-"`

	Loc        string `xml:"loc" json="loc"`
	Lastmod    string `xml:"lastmod" json="lastmod"`
	News       News2  `xml:"news" json="news"`
	Changefreq string `xml:"changefreq" json="changefreq"`
	// Publication     NewsPublication `xml:"publication"`
	// PublicationDate string          `xml:"publication_date"`
	// Title           string          `xml:"title"`
	// Keywords        string          `xml:"keywords"`
}

type News2 struct {
	NewsPublication NewsPublication `xml:"publication"`
	PublicationDate string          `xml:"publication_date" json="url"`
	Title           string          `xml:"title" json="url"`
	Keywords        string          `xml:"keywords" json="url"`
}

type NewsPublication struct {
	// XMLName  xml.Name `xml:"publication"`
	Name     string `xml:"name"`
	Language string `xml:"language"`
}

// // Location : some text
// type Location struct {
// 	Loc string `xml:"loc"`
// }
//
// // SitemapIndex : some text 2
// type SitemapIndex struct {
// 	Locations []Location `xml:"sitemap"`
// }

// SitemapIndex : some text 2
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles   []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	// Titles    []string `xml:"url>news>title"`
	// Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

func main() {
	if 11 > 10 {
		run1()
	} else if 2 > 10 {
		run2()
	} else if 3 > 10 {
		MurshalFuncTest2()
	} else {
		run3()
	}
}

func run1() {
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)

	resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if err != nil {
		print(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		fmt.Println("Location: ", Location)
		resp, err := http.Get(strings.TrimSpace(Location))
		if err != nil {
			print(err.Error())
		}
		fmt.Println("p1")
		bytes, err2 := ioutil.ReadAll(resp.Body)
		fmt.Println("p2")
		if err2 != nil {
			print(err2.Error())
		}
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)

		fmt.Println(n)

		for idx, _ := range n.Titles {
			fmt.Println(idx, "\n")
			news_map[n.Titles[idx]] =
				NewsMap{
					n.Keywords[idx],
					n.Locations[idx]}
		}
	}

	for title, data := range news_map {
		fmt.Println("\n\n\n", title)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}

func run2() {
	// var s SitemapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/politics.xml")

	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &n)

	for index := range n.Titles {
		// fmt.Println(
		// 	n.Titles[index],
		// 	n.Locations[index],
		// 	n.Keywords[index])
		fmt.Println(index, n.Locations[index])
	}

}

func run3() {
	var URLSet URLSet
	xml.Unmarshal([]byte(washPostXMLNews), &URLSet)

	fmt.Println(URLSet)

	jsonData, _ := json.MarshalIndent(URLSet, "", " ")
	fmt.Println("\n\njsonData: ")
	fmt.Println(string(jsonData))
}

func MurshalFuncTest2() {
	fmt.Println("MurshalFuncTest2 func: \n\n")
	var news News
	xml.Unmarshal([]byte(washPostXMLNews), &news)

	fmt.Println(news)

	// jsonData, _ := json.MarshalIndent(news, "", " ")
	// fmt.Println("\n\njsonData: ")
	// fmt.Println(string(jsonData))

	fmt.Println("\n\nResult: \n\n")
	for index, _ := range news.Titles {
		fmt.Println(index,
			strings.TrimSpace(news.Locations[index]), "\n",
			strings.TrimSpace(news.Titles[index]), "\n",
			strings.TrimSpace(news.Keywords[index]), "\n")
	}
}
