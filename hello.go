package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func readHtmlFromFile(fileName string) (string, error) {

	dat, err := os.ReadFile("index.html")
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func html2md(text string) (data []string) {

	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		return
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode && !slices.Contains([]atom.Atom{atom.A, atom.Script}, n.Parent.DataAtom) {
			retStr := node2Str(n)
			if len(retStr) > 0 {
				fmt.Print(retStr)
				if n.Parent != nil && n.Parent.DataAtom == atom.P && n.Parent.LastChild == n {
					fmt.Print("  ")
				}
				fmt.Println()
			}
		} else {
			node2Str(n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.DataAtom == 0 {

			} else {
				f(c)
			}

		}
	}
	f(doc)
	return
}

func node2Str(n *html.Node) (data string) {
	//fmt.Println("in node2Str")
	//fmt.Println(n.Type)
	//fmt.Print("[", n.Data, n.DataAtom)
	//fmt.Printf(" %d ]", n.DataAtom)
	//fmt.Print([]byte(n.Data))
	switch {
	case n.DataAtom == atom.A:
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Printf("[%s](%s)\n", node2Str(n.FirstChild), a.Val)
			}
		}
	case n.DataAtom == atom.Img:
		for _, a := range n.Attr {
			if a.Key == "src" {
				fmt.Printf("![%s](%s)\n", "image", a.Val)
			}
		}
	case n.DataAtom == atom.Ul:

	case n.DataAtom == atom.H1:
		fmt.Printf("# ")

	case n.DataAtom == atom.H2:
		fmt.Printf("## ")
	case n.DataAtom == atom.H3:
		fmt.Printf("### ")

	case n.DataAtom == atom.Li:
		fmt.Printf("* ")
	case n.DataAtom == atom.P:

	case n.Type == html.TextNode:
		return strings.TrimSpace(n.Data)

	default:
		return
	}
	return
}

func main() {

	fileName := "index.html"
	text, err := readHtmlFromFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	html2md(text)
}
