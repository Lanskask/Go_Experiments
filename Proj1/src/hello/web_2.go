package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const newsText string = `
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`

type NewsAggPage struct {
	Title, News string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go is Neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{
		Title: "It's a title",
		News:  newsText}

	t, err := template.ParseFiles("./basic_templating.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
