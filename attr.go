package attr

// Original Author: Siong-Ui Te (https://siongui.github.io/2016/04/15/go-getElementById-via-net-html-package/)
// Modified by Logan Yokum (lyokum@nd.edu)

import (
	"golang.org/x/net/html"
)

func GetAttr(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}

	return "", false
}

func checkAttr(n *html.Node, attr string, val string) bool {
	if n.Type == html.ElementNode {
		s, ok := GetAttr(n, attr)

		// attr found
		if ok && s == val {
			return true
		}
	}

	// not found
	return false
}

/* General Search */
func GetElementByAttr(n *html.Node, attr string, val string) *html.Node {
	// check curr node
	if checkAttr(n, attr, val) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := GetElementByAttr(c, attr, val)

		// attr found further down
		if result != nil {
			return result
		}
	}

	// attr not found
	return nil
}

/* Search Wrappers */
func GetElementById(n *html.Node, id string) *html.Node {
	return GetElementByAttr(n, "id", id)
}

func GetElementByName(n *html.Node, name string) *html.Node {
	return GetElementByAttr(n, "name", name)
}
