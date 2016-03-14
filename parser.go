package pttifierLib

import (
	"net/http"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// StrMatcher returns true when special rule of s and chars are found
type StrMatcher func(s, chars string) bool

// Rule is the rule for parsing posts
//
// TitleKey is the keywords that will be searched as if this is part of the parsing title,
//
// Author is a simple condition that NULL("") means for no author check, otherwise,
// it is not case-sensitive checking using strings.EqualFold
type Rule struct {
	TitleKey string
	Author   string
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
	targetURL := pttBaseCrawlingURL + targetBoard + defaultParsingPage + pageNum
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

// Parsing returns the result of extracting ptt web page by given rule and page.
func (rule *Rule) Parsing(root *html.Node, titleMatcher StrMatcher) (results []*Result) {
	articles := scrape.FindAll(root, scrape.ByClass("r-ent"))

	for _, article := range articles {
		skipThisParse = false

		title := getTitle(article)
		rule.compareTitle(title, titleMatcher)

		author := getAuthor(article)
		rule.compareAuthor(author, strings.EqualFold)

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

// compareTitle returns true if the parsing title is obay the strMatcher
func (rule *Rule) compareTitle(title string, strMatcher StrMatcher) {
	if !skipThisParse {
		if !strMatcher(title, rule.TitleKey) {
			skipThisParse = true
		}
	}
}

// compareAuthor returns true if the parsing author is obay the strMatcher
func (rule *Rule) compareAuthor(author string, strMatcher StrMatcher) {
	if !skipThisParse {
		if rule.Author != "" && !strMatcher(author, rule.Author) {
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
		return title
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
		return url
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
		return author
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
		return date
	}

	return
}
