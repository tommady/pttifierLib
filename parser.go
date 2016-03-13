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
func (rule *Rule) Parsing(root *html.Node, strMatcher StrMatcher) ([]*Result, error) {
	articles := scrape.FindAll(root, scrape.ByClass("r-ent"))

	results := []*Result{}
	for _, article := range articles {
		result := new(Result)

		title, ok := scrape.Find(article, scrape.ByTag(atom.A))
		if !ok {
			// post has been delete
			continue
		}

		result.Title = scrape.Text(title)

		// checking is this parsing title mathces the rule
		if !strMatcher(result.Title, rule.TitleKey) {
			continue
		}

		author, ok := scrape.Find(article, scrape.ByClass("author"))
		if !ok {
			// this should not be happend, unless the ptt server done
			continue
		}

		result.Author = scrape.Text(author)
		if rule.Author != "" && !strings.EqualFold(rule.Author, result.Author) {
			continue
		}

		result.URL = pttBaseURL + scrape.Attr(title, "href")

		date, ok := scrape.Find(article, scrape.ByClass("date"))
		if !ok {
			// this should not be happend, unless the ptt server done
			continue
		}

		result.Date = scrape.Text(date)
		results = append(results, result)
	}

	return results, nil
}
