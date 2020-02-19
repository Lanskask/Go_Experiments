package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

var washPostXMLNews = (`
<urlset>
  <publication>
			<news:name>Washington Post</news:name>
			<news:language>en</news:language>
	</publication>
	<publication>
			<news:name>Washington Post</news:name>
			<news:language>en</news:language>
	</publication>
</urlset>
`)

type URLSet struct {
	XMLName xml.Name      `xml:"urlset" json:"-"`
	URLList []Publication `xml:"publication" json="url"`
}

// type Publication struct {
// 	// XMLName         xml.Name        `xml:"url"  json="-"`
// 	Name     string `xml:"news name,omitempty" prefix="news"`
// 	Language string `xml:"news language,omitempty" prefix="news"`
// }

type Publication struct {
	// XMLName         xml.Name        `xml:"url"  json="-"`
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
