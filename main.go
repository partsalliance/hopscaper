package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func get_html_links() {
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
			return

		case tt == html.StartTagToken:
			t := z.Token()

			if t.Data == "td" {
				inner := z.Next()
				if inner == html.TextToken {
					text := (string)(z.Text())
					t := strings.TrimSpace(text)
					fmt.Println(t)
				} else if inner == html.StartTagToken {
					for _, a := range z.Token().Attr {
						if a.Key == "href" {
							if strings.Contains(a.Val, "hops/") {
								inner := z.Next()
								if inner == html.TextToken {
									fmt.Println(z.Token().Data)
									hop_name_list = append(hop_name_list, z.Token().Data)
								}
							}

						}
					}
				}
			}
		}
	}
}

func main() {
	fmt.Println("start")
	//get_hop_data("http://www.beersmith.com/hops/zeus.htm")
	get_html_links()

}
