package pttifierLib_test

import (
	"strings"
	"testing"

	"github.com/tommady/pttifierLib"
	"golang.org/x/net/html"
)

func TestGetArticleContents(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestContentHTML))
	p := pttifierLib.NewPostCrawler(root)

	actual := p.GetArticleContents()

	if strings.Contains(actual, "還沒崩潰完？") {
		t.Errorf("GetArticleContents: contain tweet's content")
	}
	if strings.Contains(actual, "新唐人亞太台") {
		t.Errorf("GetArticleContents: contain quotation's content")
	}
}

func TestGetTweets(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestContentHTML))
	p := pttifierLib.NewPostCrawler(root)

	expects := []*pttifierLib.Tweet{
		{
			"skyexers",
			"還沒崩潰完？",
			"04/03 21:07",
			pttifierLib.TweetTagBoo,
		},
		{
			"hkcdc",
			"不判死的確是問題",
			"04/03 21:15",
			pttifierLib.TweetTagNormal,
		},
		{
			"pictograma",
			"一樓崩潰標準示範",
			"04/03 21:15",
			pttifierLib.TweetTagPraise,
		},
		{
			"drigo",
			"學到原來台灣人書唸不夠多, 所以獵廢死",
			"04/03 21:16",
			pttifierLib.TweetTagNormal,
		},
		{
			"ainor",
			"一堆22K都沒殺人了",
			"04/03 21:19",
			pttifierLib.TweetTagBoo,
		},
		{
			"dave01",
			"沒學到  過一陣子 霉不報導 一堆人就忘了",
			"04/03 21:30",
			pttifierLib.TweetTagNormal,
		},
		{
			"puorg",
			"你太中肯了",
			"04/03 22:27",
			pttifierLib.TweetTagPraise,
		},
		{
			"xgodtw",
			"有點中肯",
			"04/03 23:51",
			pttifierLib.TweetTagNormal,
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
