package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

var washPostXMLNews = (`
<urlset>
	<url>
		<news:news>
			<news:publication>
				<news:name>Washington Post</news:name>
				<news:language>en</news:language>
			</news:publication>
			<news:publication_date>2020-02-16T19:33:00Z</news:publication_date>
			<news:title>
				<![CDATA[ Trump kicks off Daytona 500 as grand marshal ]]>
			</news:title>
			<news:keywords>
				<![CDATA[
				, trump, the great american race, rv, keep america great, air force one
				]]>
			</news:keywords>
		</news:news>
	</url>

	<url>
		<news:news>
			<news:publication>
				<news:name>Washington Post</news:name>
				<news:language>en</news:language>
			</news:publication>
			<news:publication_date>2020-02-16T20:31:57Z</news:publication_date>
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
	</url>
</urlset>
`)

type URLSet struct {
	// XMLName xml.Name `xml:"urlset" json:"-"`
	URLList []Url `xml:"url" json="url"`
}

type Url struct {
	News News `xml:"news"`
}

type News struct {
	// XMLName         xml.Name    `xml:"url" json:"-"`
	Publications    Publication `xml:"publication" json="url"`
	PublicationDate string      `xml:"publication_date" json="url"`
	Title           string      `xml:"title" json="url"`
	Keywords        string      `xml:"keywords" json="url"`
}

// type Publication struct {
// 	// XMLName         xml.Name        `xml:"url"  json="-"`
// 	Name     string `xml:"news name,omitempty" prefix="news"`
// 	Language string `xml:"news language,omitempty" prefix="news"`
// }

type Publication struct {
	// XMLName  xml.Name `xml:"publication"  json="-"`
	Name     string `xml:"name"`
	Language string `xml:"language"`
}

// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

func main() {
	run3()
}
func run3() {
	var URLSet URLSet
	xml.Unmarshal([]byte(washPostXMLNews), &URLSet)

	fmt.Println(URLSet)

	jsonData, _ := json.MarshalIndent(URLSet, "", " ")
	fmt.Println("\n\njsonData: ")
	fmt.Println(string(jsonData))
}
