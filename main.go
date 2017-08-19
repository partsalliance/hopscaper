package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
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

func main() {
	fmt.Println("start")
	//get_hop_data("http://www.beersmith.com/hops/zeus.htm")
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

	fmt.Printf("%#v\n", divided)


}