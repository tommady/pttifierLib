package pttifierLib_test

import (
	"strings"
	"testing"

	"github.com/tommady/pttifierLib"
	"golang.org/x/net/html"
)

func TestGetPrevPageLink(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestBoardHTML))
	b := pttifierLib.NewBoardCrawler(root)

	expect := pttifierLib.PttBaseURL + "/bbs/LoL/index4380.html"
	actual := b.GetPrevPageLink()
	if actual != expect {
		t.Errorf("GetPrevPageLink: expect-> %q, actual-> %q", expect, actual)
	}
}

func TestGetNextPageLink(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestBoardHTML))
	b := pttifierLib.NewBoardCrawler(root)

	expect := pttifierLib.PttBaseURL + "/bbs/LoL/index4382.html"
	actual := b.GetNextPageLink()
	if actual != expect {
		t.Errorf("GetPrevPageLink: expect-> %q, actual-> %q", expect, actual)
	}
}

func TestGetArticlesInfos(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestBoardHTML))
	b := pttifierLib.NewBoardCrawler(root)

	expects := []*pttifierLib.PostBaseInfo{
		{
			"https://www.ptt.cc/bbs/LoL/M.1457622048.A.D80.html",
			"[公告] LoL 樂透開獎",
			"[彩券]",
			"3/10",
			11,
		},
		{
			"https://www.ptt.cc/bbs/LoL/M.1457622048.A.Z80.html",
			"[公告] LoL 小樂透開獎",
			"[彩券]",
			"3/11",
			11,
		},
		{
			"https://www.ptt.cc/bbs/LoL/M.1457622282.A.D4F.html",
			"Re: [問題] AlphaGo打LoL的話會怎麼樣",
			"lzainside",
			"3/10",
			4,
		},
		{
			"https://www.ptt.cc/bbs/LoL/M.1457622407.A.C98.html",
			"[問題] 關於四名中路的比較！",
			"lizst980074",
			"3/10",
			2,
		},
		{
			"https://www.ptt.cc/bbs/LoL/M.1457622551.A.650.html",
			"[閒聊] 對於XG挺失望的",
			"iamfenixsc",
			"3/10",
			1,
		},
	}

	actuals := b.GetArticlesInfos()

	for i, expect := range expects {
		actual := actuals[i]
		if actual.URL != expect.URL {
			t.Errorf("GetArticlesInfos: expect-> %q, actual-> %q", expect.URL, actual.URL)
		}
		if actual.Title != expect.Title {
			t.Errorf("GetArticlesInfos: expect-> %q, actual-> %q", expect.Title, actual.Title)
		}
		if actual.Author != expect.Author {
			t.Errorf("GetArticlesInfos: expect-> %q, actual-> %q", expect.Author, actual.Author)
		}
		if actual.Date != expect.Date {
			t.Errorf("GetArticlesInfos: expect-> %q, actual-> %q", expect.Date, actual.Date)
		}
		if actual.TweetAmount != expect.TweetAmount {
			t.Errorf("GetArticlesInfos: expect-> %q, actual-> %q", expect.TweetAmount, actual.TweetAmount)
		}
	}
}
