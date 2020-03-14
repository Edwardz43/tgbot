package crawl

// Crawler defines crawl function for specific crawler
type Crawler interface {
	Get() string
	// crawl(r *http.Response) (l []string, err error)
	// visitPages(n *html.Node)
	// visitTarget(n *html.Node)
	// forEachNode(n *html.Node, pre, post func(n *html.Node))
}
