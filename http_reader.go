package main

import (
	"fmt"
	"io"
//	"io/ioutil"
	"log"
        "net/http"
	"os"
//	"strings"
//	"golang.org/x/net/html"
       )

func check(e error) {
    if e != nil {
	log.Fatal(e)
    }
}

func main() {
	page, err := os.Create("citystar_ru.html")
	check(err)
	defer page.Close()

	s_page, err := os.Create("citystar_ru_.html")
	check(err)
	defer s_page.Close()
	
	res, err := http.Get("http://citystar.ru/detal.htm?d=37&nm=%CE%E1%FA%FF%E2%EB%E5%ED%E8%FF%20%E2%20%F0%F3%E1%F0%E8%EA%E5%20%22%D1%E4%E0%EC%20%EA%EE%EC%EC%E5%F0%F7%E5%F1%EA%F3%FE%20%ED%E5%E4%E2%E8%E6%E8%EC%EE%F1%F2%FC%20%E2%20%E3.%20%CC%E0%E3%ED%E8%F2%EE%E3%EE%F0%F1%EA%E5%22")
	check(err)
	defer res.Body.Close()
	
	size, err := io.Copy(page, res.Body)
	check(err)
	fmt.Println("File size: ", size)
	
/*	s, err := string(ioutil.ReadAll(res.Body))
	check(err)
	doc, err := html.Parse(strings.NewReader(s))
	check(err)
	
	size, err := io.Copy(s_page, s)
	check(err)
	fmt.Println("File size: ", size)

	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c:= n.FirstChild; c != nil; c = n.NextSibling {
			f(c)
		}
	}
	f(doc)
*/
}