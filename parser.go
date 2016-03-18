package pttifierLib

import (
	"errors"
	"net/http"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// StrMatcher returns true when special rule of s and chars are found
type StrMatcher func(s, chars string) bool

// Rule is the rule for parsing posts
type Rule struct {
	TitleKey      string
	Author        string
	TitleMatcher  StrMatcher
	AuthorMatcher StrMatcher
}

// Result is the parsing result that will be returned
type Result struct {
	URL    string
	Title  string
	Author string
	Date   string
}

const (
	pttBaseURL         = "https://www.ptt.cc"
	pttBaseCrawlingURL = "https://www.ptt.cc/bbs/"
	defaultParsingPage = "/index"
)

var (
	skipThisParse bool
)

// GetRootNode returns the root node of ginving Board and board page number,
// if want to start from default index then pageNum should be NULL ("")
func GetRootNode(targetBoard, pageNum string) (*html.Node, error) {
	targetURL := pttBaseCrawlingURL + targetBoard + defaultParsingPage + pageNum + ".html"
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, ReportError("failed on new request", err)
	}

	// for some specific board need over 18 years old checks
	req.Header.Set("Cookie", "over18=1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, ReportError("failed on default client do", err)
	}

	root, err := html.Parse(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, ReportError("html parse web page fail", err)
	}

	return root, nil
}

// GetActionBarNode returns <div class="action-bar"> from givin root node
func GetActionBarNode(root *html.Node) (*html.Node, error) {
	if root == nil {
		return nil, ReportError("Get nil root, can not extracts out Action Bar Node", nil)
	}

	n, ok := scrape.Find(root, scrape.ByClass("action-bar"))
	if !ok {
		return nil, ReportError("Can not find Action Bar Node from givin node", nil)
	}

	return n, nil
}

// GetRListNode returns <div class="r-list-container bbs-screen"> from givin root node
func GetRListNode(root *html.Node) (*html.Node, error) {
	if root == nil {
		return nil, ReportError("Get nil root, can not extracts out R List Node", nil)
	}

	n, ok := scrape.Find(root, scrape.ByClass("bbs-screen"))
	if !ok {
		return nil, ReportError("Can not find R List Node from givin node", nil)
	}

	return n, nil
}

// GetPrevNextPageLink returns previous and next page's links,
func GetPrevNextPageLink(root *html.Node) (prev string, next string, err error) {
	if root == nil {
		return "", "", ReportError("Get nil root, can not extracts out previous and next page's links", nil)
	}

	ns := scrape.FindAll(root, scrape.ByClass("wide"))
	for _, n := range ns {
		text := scrape.Text(n)
		link := scrape.Attr(n, "href")
		if link != "" {
			if strings.Contains(text, "上頁") {
				prev = pttBaseURL + link
			} else if strings.Contains(text, "下頁") {
				next = pttBaseURL + link
			}
		}
	}

	return prev, next, nil
}

// RemoveBottumAnnouncements returns the whole html tree without bottum announcements if contains,
func RemoveBottumAnnouncements(root *html.Node) (err error) {
	if root == nil {
		return ReportError("Gvin root is nil", nil)
	}

	n, ok := scrape.Find(root, scrape.ByClass("r-list-sep"))
	if !ok {
		// gvin root is not contain bottum announcements, no need to bother with it
		return nil
	}

	// WARNING: if givin root node is not <div class="r-list-container bbs-screen"> or returns by
	// function GetRListNode, may causes panic
	defer func() {
		if r := recover(); r != nil {
			msg := "root node is not <r-list-container bbs-screen> node"
			switch x := r.(type) {
			case string:
				err = ReportError(msg, errors.New(x))
			case error:
				err = ReportError(msg, x)
			default:
				err = ReportError(msg, errors.New("Unknown panic"))
			}
		}
	}()

	var tmpNext *html.Node
	for c := n; c != nil; c = tmpNext {
		tmpNext = c.NextSibling
		root.RemoveChild(c)
	}

	return nil
}

// Parsing returns the result of extracting ptt web page by given rule and page
func (rule *Rule) Parsing(root *html.Node) (results []*Result) {
	if root == nil {
		return
	}

	articles := scrape.FindAll(root, scrape.ByClass("r-ent"))

	for _, article := range articles {
		skipThisParse = false

		title := getTitle(article)
		rule.compareTitle(title)

		author := getAuthor(article)
		rule.compareAuthor(author)

		url := getURL(article)
		date := getDate(article)

		if skipThisParse {
			continue
		}

		result := new(Result)
		result.Title = title
		result.URL = url
		result.Author = author
		result.Date = date
		results = append(results, result)
	}

	return results
}

// compareTitle returns true if the parsing title is obay the strMatcher,
// if rule's TitleKey == "" means a post with any title will be accepted as result
func (rule *Rule) compareTitle(title string) {
	if !skipThisParse {
		if rule.TitleKey != "" && !rule.TitleMatcher(title, rule.TitleKey) {
			skipThisParse = true
		}
	}
}

// compareAuthor returns true if the parsing author is obay the strMatcher,
// if rule's Author == "" means a post with any author will be accepted as result
func (rule *Rule) compareAuthor(author string) {
	if !skipThisParse {
		if rule.Author != "" && !rule.AuthorMatcher(author, rule.Author) {
			skipThisParse = true
		}
	}
}

// getTitle returns the post published title
func getTitle(article *html.Node) (title string) {
	if !skipThisParse {
		t, ok := scrape.Find(article, scrape.ByTag(atom.A))
		if !ok {
			// post has been deleted
			skipThisParse = true
			return
		}

		title = scrape.Text(t)
	}

	return
}

// getURL returns the post published url link
func getURL(article *html.Node) (url string) {
	if !skipThisParse {
		t, ok := scrape.Find(article, scrape.ByTag(atom.A))
		if !ok {
			// post has been deleted
			skipThisParse = true
			return
		}

		url = pttBaseURL + scrape.Attr(t, "href")
	}

	return
}

// getAuthor returns the post published author
func getAuthor(article *html.Node) (author string) {
	if !skipThisParse {
		a, ok := scrape.Find(article, scrape.ByClass("author"))
		if !ok {
			// this should not be happend, unless the ptt server suddenly down
			skipThisParse = true
			return
		}

		author = scrape.Text(a)
	}

	return
}

// getDate returns the post published day
func getDate(article *html.Node) (date string) {
	if !skipThisParse {
		d, ok := scrape.Find(article, scrape.ByClass("date"))
		if !ok {
			// this should not be happend, unless the ptt server suddenly down
			skipThisParse = true
			return
		}

		date = scrape.Text(d)
	}

	return
}
