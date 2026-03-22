package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func main() {
	htmlTokenizer()
	// parserTest()
	// writer()
	// htmlReader()
}

func parserTest() {
	s := `<p>Hello, Gemini!</p>`
	r := strings.NewReader(s)

	fmt.Println("s: ", s)

	_, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
}

func htmlTokenizer() {
	// input := `<a href="https://google.com"> Google </a> <p> Hello </p> <a href="/about"> About </a> `
	input := `
		<div id="inventory-container">
			<article class="data-card">
				<div class="info">
					<span class="price-tag">$899.99</span>
					<hr/>
				</div>
			</article>
		</div>
	`

	tokenizer := html.NewTokenizer(strings.NewReader(input))

	for {
		// 1. Advance to the next token
		tokenType := tokenizer.Next()

		// 2. Check for the end of the document
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break // End of file reached gracefully
			}
			fmt.Printf("Error: %v", err)
			break
		}

		// 3. Process the token
		token := tokenizer.Token()
		switch tokenType {
		case html.StartTagToken:
			{
				fmt.Println("----------------------------")
				fmt.Println("tag ", token.Data)

				// Check if the tag is an anchor <a>
				for _, attr := range token.Attr {
					fmt.Println("key:", attr.Key, "val:", attr.Val)
				}
			}
		case html.TextToken:
			{
				fmt.Println("------------- Text -------------------")
				fmt.Println("Type:", token.Type)
				fmt.Println("Text ", token.Data)
			}
		case html.EndTagToken:
			{
				fmt.Println("-----------------")
				fmt.Println("End of tag ", token.Data)
			}
		case html.SelfClosingTagToken:
			{
				fmt.Println("==================")
				fmt.Println("Self closing tag ", token.Data)
			}
		}
	}
}

func htmlReader() {
	file, err := os.Open("min-test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 2. Initialize goquery
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Debugging Tests

	// Test A: Find by ID and get text
	title := doc.Find("#main-title").Text()
	fmt.Printf("Title found: %s\n", title)

	// Test B: Loop through list and get attributess
	doc.Find("ul#item-list li").Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("data-id")
		fmt.Printf("Index %d: %s (ID: %s)\n", i, s.Text(), id)
	})

	// Test C: Finding by Text (using :contains)
	doc.Find("a:contains('Contact')").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		fmt.Printf("Link found by text: %s\n", href)
	})
}

func writer() {
	input := `<div class="nav">Hello</div>`
	// This initializes the 'r' (io.Reader) and 'buf' in the struct
	z := html.NewTokenizer(strings.NewReader(input))

	for {
		tt := z.Next() // This updates 'tt', 'raw', 'data', and 'err' internally

		if tt == html.ErrorToken {
			// This checks the 'z.err' field
			break
		}

		// z.Token() uses the 'raw' and 'data' spans to
		// convert bytes in 'buf' into a readable struct
		token := z.Token()

		if tt == html.StartTagToken {
			fmt.Printf("Tag: %s\n", token.Data) // 'Data' comes from 'z.data' span
			for _, a := range token.Attr {
				fmt.Printf(" - Attr: %s = %s\n", a.Key, a.Val)
			}
		}

		if tt == html.TextToken {
			fmt.Printf("Text: %s\n", token.Data)
		}
	}
}
