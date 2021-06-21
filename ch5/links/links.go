package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func Extract(url string) (links []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("link: getting %s:%s", url, resp.Status)
	}
	node, err := html.Parse(resp.Body)
	if err != nil{
		return nil, fmt.Errorf("link: parsing %s as HTML: %v",url,err)
	}
	visitNode := func(n *html.Node) {
		if n.Type==html.ElementNode && n.Data=="a"{
			for _, attribute := range n.Attr {
				if attribute.Key!="href"{
					continue
				}
				link, err := resp.Request.URL.Parse(attribute.Val)
				if err == nil{
					links = append(links,link.String())
				}
			}
		}
	}
	ForeachNode(visitNode,node,nil)
	return links,nil
}

func ForeachNode(pre func(n *html.Node), n *html.Node ,post func(n *html.Node)){
	if pre!=nil{
		pre(n)
	}
	for i := n.FirstChild; i != nil; i=i.NextSibling {
		ForeachNode(pre,i,post)
	}
	if post!=nil{
		post(n)
	}
}
