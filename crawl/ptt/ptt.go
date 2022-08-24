package ptt

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// BoardMap mapps bot command to specific PTT board
var BoardMap = map[string]string{
	"/b":    "Beauty",
	"/c":    "C_Chat",
	"/e":    "e-shopping",
	"/s":    "Stock",
	"/se":   "sex",
	"/g":    "Gossiping",
	"/l":    "Lifeismoney",
	"/n":    "NBA",
	"/bs":   "Baseball",
	"/wh":   "nCoV2019",
	"/pc":   "PC_Shopping",
	"/help": "List",
}

// GetInstance return a instance of PTT crawler
func GetInstance(board string) *Crawler {
	return &Crawler{board: board}
}

// Crawler represents instance of beauty crawler
type Crawler struct {
	board    string
	target   string
	links    []string
	posts    []*post
	lastPage int
	//pageSize int
}

type post struct {
	Title   string
	Nrecord string
	URL     string
}

func (p *post) toString() string {
	return fmt.Sprintf("%v  推文數:%v  %v", p.Title, p.Nrecord, p.URL)
}

var (
	cookie   = http.Cookie{Name: "over18", Value: "1"}
	pageSize = 10
)

func errHandler(msg string, err error) {
	if err != nil {
		log.Printf("%s : [%v]\n", msg, err)
	}
}

func (c *Crawler) listCMDs() string {
	var l string

	for k, v := range BoardMap {
		l += fmt.Sprintf("%s : %s\n", k, v)
	}

	return l
}

// Get gets the target
func (c *Crawler) Get() string {
	if c.board == "List" {
		return c.listCMDs()
	}

	c.target = fmt.Sprintf("https://www.ptt.cc/bbs/%s/index.html", c.board)
	urls, err := c.getURL()
	if err != nil {
		//TODO
	}

	var links []string

	var wg sync.WaitGroup

	for index := 0; index < len(urls); index++ {
		client2 := &http.Client{}
		wg.Add(1)
		url := urls[index]
		go func(url string) {
			reqTarget, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Println(err)
			}
			reqTarget.AddCookie(&cookie)

			respTarget, err := client2.Do(reqTarget)
			if err != nil {
				log.Println(err)
			}

			err = c.crawlPost(respTarget, &wg)
			if err != nil {
				log.Printf("Error crawling : %v", url)
			}
		}(url)
	}
	wg.Wait()

	//todo sort
	sort.Slice(c.posts, func(i, j int) bool {
		recordI := c.posts[i].Nrecord
		recordJ := c.posts[j].Nrecord

		if recordI == "爆" {
			return true
		} else if recordJ == "爆" {
			return false
		}
		return recordI > recordJ
	})

	for i := 0; i < len(c.posts); i++ {
		links = append(links, c.posts[i].toString())
		ss := strings.Join(links, "\n")
		l := len(ss)
		if l > 2000 {
			links = links[0 : len(links)-1]
		}
	}

	c.posts = make([]*post, 0)

	return strings.Join(links, "\n")
}

func (c *Crawler) getURL() ([]string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", c.target, nil)

	errHandler("Http Request init fail", err)

	req.AddCookie(&cookie)

	resp, err := client.Do(req)

	errHandler("Http send fail", err)

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("parsing HTML: %v", err)
	}

	c.lastPage = c.getLastPage(doc)

	var urls = make([]string, pageSize)

	for i := 0; i < pageSize; i++ {
		urls[i] = fmt.Sprintf("https://www.ptt.cc/bbs/%s/index%d.html", c.board, c.lastPage-i)
	}

	c.links = make([]string, 0)

	return urls, nil
}

func (c *Crawler) crawlPost(r *http.Response, wg *sync.WaitGroup) error {
	// counter++
	doc, err := html.Parse(r.Body)
	if err != nil {
		wg.Done()
		return fmt.Errorf("parsing HTML: %v", err)
	}

	r.Body.Close()

	forEachNode(doc, c.visitTarget, nil)

	wg.Done()
	// log.Printf("Done [%d]", counter)
	return nil
}

func (c *Crawler) visitPages(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "btn wide" {
				href := n.Attr[1].Val
				c.links = append(c.links, href)
			}
		}

	}
}

func (c *Crawler) visitTarget(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "class" &&
				a.Val == "nrec" &&
				n.FirstChild != nil &&
				n.NextSibling.NextSibling.FirstChild.NextSibling != nil {

				record := n.FirstChild.FirstChild.Data

				if record != "爆" {
					count, err := strconv.Atoi(record)
					if err != nil || count < 50 {
						break
					}
				}

				title := n.NextSibling.
					NextSibling.
					FirstChild.
					NextSibling.
					LastChild.Data

				if title[1:7] == "公告" {
					break
				}

				url := "https://www.ptt.cc/" +
					n.NextSibling.
						NextSibling.
						FirstChild.
						NextSibling.Attr[0].Val

				p := &post{
					Title:   title,
					Nrecord: record,
					URL:     url,
				}
				log.Printf("title : [%v], record : [%v], url : [%v]", title, record, url)
				c.posts = append(c.posts, p)
			}
		}

	}
}

// getLastPage return the last page of the board
func (c *Crawler) getLastPage(n *html.Node) int {
	forEachNode(n, c.visitPages, nil)
	s := c.links[1]

	reg, err := regexp.Compile("index([0-9]+?)\\.html")
	if err != nil {
		log.Println(err)
	}

	sa := reg.FindStringSubmatch(s)

	if sa == nil {
		//TODO log
		return 0
	}

	i, _ := strconv.Atoi(sa[1])

	return i + 1
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
