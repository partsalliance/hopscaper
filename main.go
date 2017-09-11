package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"log"
	"os"
	"encoding/csv"
)

func get_html_links() []string {
	url := "http://beersmith.com/hop-list/"
	response, _ := http.Get(url)
	hop_name_list := []string{}
	z := html.NewTokenizer(response.Body)
	for z.Token().Data != "html" {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			response.Body.Close()
			fmt.Println("finished")
			break

		case tt == html.StartTagToken:
			t := z.Token()

			if t.Data == "td" {
				inner := z.Next()
				if inner == html.TextToken {
					text := (string)(z.Text())
					t := strings.TrimSpace(text)
					if t != "" {
						hop_name_list = append(hop_name_list, t)
					}

				} else if inner == html.StartTagToken {
					for _, a := range z.Token().Attr {
						if a.Key == "href" {
							if strings.Contains(a.Val, "hops/") {
								inner := z.Next()
								if inner == html.TextToken {
									hop_name_list = append(hop_name_list, z.Token().Data)
								}
							}

						}
					}
				}
			}
		}
	}
	return hop_name_list
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func main() {
	fmt.Println("start")
	hops := get_html_links()

	var divided [][]string

	chunkSize := 4
	for i := 0; i < len(hops); i += chunkSize {

		end := i + chunkSize

		if end > len(hops) {
			end = len(hops)
		}

		divided = append(divided, hops[i:end])
	}

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, hop := range divided {
		err := writer.Write(hop)
		checkError("Cannot write to file", err)
	}

}
