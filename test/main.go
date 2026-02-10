package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// 1. Open the local file
	s := `<p>Hello, Gemini!</p>`
	r := strings.NewReader(s)

	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	// doc is now the root of your HTML tree
	fmt.Printf("Root type: %v\n", doc.FirstChild.Data)

	// file, err := os.Open("min-test.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// // 2. Initialize goquery
	// doc, err := goquery.NewDocumentFromReader(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 3. Debugging Tests

	// // Test A: Find by ID and get text
	// title := doc.Find("#main-title").Text()
	// fmt.Printf("Title found: %s\n", title)

	// Test B: Loop through list and get attributes
	// doc.Find("ul#item-list li").Each(func(i int, s *goquery.Selection) {
	// 	id, _ := s.Attr("data-id")
	// 	fmt.Printf("Index %d: %s (ID: %s)\n", i, s.Text(), id)
	// })

	// // Test C: Finding by Text (using :contains)
	// doc.Find("a:contains('Contact')").Each(func(i int, s *goquery.Selection) {
	// 	href, _ := s.Attr("href")
	// 	fmt.Printf("Link found by text: %s\n", href)
	// })
}
