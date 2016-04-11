package pttifierLib_test

import (
	"net"
	"strings"
	"testing"

	"github.com/tommady/pttifierLib"
	"golang.org/x/net/html"
)

func TestGetAuthor(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestContentHTML))
	p := pttifierLib.NewPostCrawler(root)

	expect := "senyek9527 ()"
	actual := p.GetAuthor()

	if expect != actual {
		t.Errorf("GetAuthor: expect-> %s, actual-> %s", expect, actual)
	}
}

func TestGetTitle(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestContentHTML))
	p := pttifierLib.NewPostCrawler(root)

	expect := "Re: [新聞] 5年7起隨機殺人 台灣社會學到甚麼？"
	actual := p.GetTitle()

	if expect != actual {
		t.Errorf("GetTitle: expect-> %s, actual-> %s", expect, actual)
	}
}

func TestGetDate(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestContentHTML))
	p := pttifierLib.NewPostCrawler(root)

	expect := "Sun Apr  3 21:04:15 2016"
	actual := p.GetDate()

	if expect != actual {
		t.Errorf("GetDate: expect-> %s, actual-> %s", expect, actual)
	}
}

func TestGetIP(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestContentHTML))
	p := pttifierLib.NewPostCrawler(root)

	expect := net.ParseIP("126.148.111.173")
	actual := p.GetIP()

	if !expect.Equal(actual) {
		t.Errorf("GetIP: expect-> %s, actual-> %s", expect.String(), actual.String())
	}
}

func TestGetURL(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestContentHTML))
	p := pttifierLib.NewPostCrawler(root)

	expect := "https://www.ptt.cc/bbs/Gossiping/M.1459688662.A.CFE.html"
	actual := p.GetURL()

	if expect != actual {
		t.Errorf("GetURL: expect-> %s, actual-> %s", expect, actual)
	}
}

func TestGetContent(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestContentHTML))
	p := pttifierLib.NewPostCrawler(root)

	actual := p.GetContent()

	if strings.Contains(actual, "還沒崩潰完？") {
		t.Errorf("GetContent: contain tweet's content")
	}
	if strings.Contains(actual, "新唐人亞太台") {
		t.Errorf("GetContent: contain quotation's content")
	}
}

func TestGetTweets(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestContentHTML))
	p := pttifierLib.NewPostCrawler(root)

	expects := []*pttifierLib.Tweet{
		{
			Author:  "skyexers",
			Content: "還沒崩潰完？",
			Date:    "04/03 21:07",
			Tag:     pttifierLib.TweetTagBoo,
		},
		{
			Author:  "hkcdc",
			Content: "不判死的確是問題",
			Date:    "04/03 21:15",
			Tag:     pttifierLib.TweetTagNormal,
		},
		{
			Author:  "pictograma",
			Content: "一樓崩潰標準示範",
			Date:    "04/03 21:15",
			Tag:     pttifierLib.TweetTagPraise,
		},
		{
			Author:  "drigo",
			Content: "學到原來台灣人書唸不夠多, 所以獵廢死",
			Date:    "04/03 21:16",
			Tag:     pttifierLib.TweetTagNormal,
		},
		{
			Author:  "ainor",
			Content: "一堆22K都沒殺人了",
			Date:    "04/03 21:19",
			Tag:     pttifierLib.TweetTagBoo,
		},
		{
			Author:  "dave01",
			Content: "沒學到  過一陣子 霉不報導 一堆人就忘了",
			Date:    "04/03 21:30",
			Tag:     pttifierLib.TweetTagNormal,
		},
		{
			Author:  "puorg",
			Content: "你太中肯了",
			Date:    "04/03 22:27",
			Tag:     pttifierLib.TweetTagPraise,
		},
		{
			Author:  "xgodtw",
			Content: "有點中肯",
			Date:    "04/03 23:51",
			Tag:     pttifierLib.TweetTagNormal,
		},
	}

	actuals := p.GetTweets()
	for i, expect := range expects {
		actual := actuals[i]
		if expect.Author != actual.Author {
			t.Errorf("GetTweets: expect-> %q, actual-> %q", expect.Author, actual.Author)
		}
		if expect.Content != actual.Content {
			t.Errorf("GetTweets: expect-> %q, actual-> %q", expect.Content, actual.Content)
		}
		if expect.Date != actual.Date {
			t.Errorf("GetTweets: expect-> %q, actual-> %q", expect.Date, actual.Date)
		}
		if expect.Tag != actual.Tag {
			t.Errorf("GetTweets: expect-> %q, actual-> %q", expect.Tag, actual.Tag)
		}
	}
}
