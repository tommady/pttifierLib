package pttifierLib

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

const (
	PttBaseURL         = "https://www.ptt.cc"
	PttBaseCrawlingURL = "https://www.ptt.cc/bbs/"
	DefaultParsingPage = "/index"
)

type BaseInfo struct {
	URL    string
	Title  string
	Author string
	Date   string
}

func WrapBoardPageLink(targetBoard, pageNum string) string {
	return PttBaseCrawlingURL + targetBoard + DefaultParsingPage + pageNum + ".html"
}

func GetNodeFromLink(targetURL string) (*html.Node, error) {
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed on new request", err)
	}

	// for some specific board need over 18 years old checks
	req.Header.Set("Cookie", "over18=1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed on default client do", err)
	}

	root, err := html.Parse(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("html parse web page fail", err)
	}

	return root, nil
}

func GetActionBarNode(root *html.Node) *html.Node {
	if root == nil {
		return nil
	}

	n, ok := scrape.Find(root, scrape.ByClass("action-bar"))
	if !ok {
		return nil
	}

	return n
}

func GetRListNode(root *html.Node) *html.Node {
	if root == nil {
		return nil
	}

	n, ok := scrape.Find(root, scrape.ByClass("bbs-screen"))
	if !ok {
		return nil
	}

	return n
}

func RemoveBottumAnnouncements(rListNode *html.Node) {
	if rListNode == nil {
		return
	}

	n, ok := scrape.Find(rListNode, scrape.ByClass("r-list-sep"))
	if !ok {
		// This board page is not contain bottum announcements, no need to bother with it
		return
	}

	var tmpNext *html.Node
	for c := n; c != nil; c = tmpNext {
		tmpNext = c.NextSibling
		rListNode.RemoveChild(c)
	}
}

func LessThanComparison(n, comparison int) bool {
	if n < comparison {
		return true
	}
	return false
}

func LessThanOrEqualToComparison(n, comparison int) bool {
	if n <= comparison {
		return true
	}
	return false
}

func EqualToComparison(n, comparison int) bool {
	if n == comparison {
		return true
	}
	return false
}

func GreaterThanComparison(n, comparison int) bool {
	if n > comparison {
		return true
	}
	return false
}

func GreaterThanOrEqualToComparison(n, comparison int) bool {
	if n >= comparison {
		return true
	}
	return false
}

func TweetAmountStringToInt(strTweetAmount string) (intTweetAmount int) {
	if strTweetAmount == "爆" {
		intTweetAmount = 100
	} else if strings.HasPrefix(strTweetAmount, "X") {
		if strTweetAmount == "XX" {
			intTweetAmount = -100
		} else {
			if n, err := strconv.Atoi(string(strTweetAmount[1])); err == nil {
				intTweetAmount = -10 * n
			}
		}
	} else {
		if n, err := strconv.Atoi(strTweetAmount); err == nil {
			intTweetAmount = n
		}
	}
	return
}
