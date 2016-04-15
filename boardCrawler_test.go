package pttifierLib_test

import (
	"sort"
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

type GetPostsInfosSortWrapper []*pttifierLib.BoardInfo

func (w GetPostsInfosSortWrapper) Len() int           { return len(w) }
func (w GetPostsInfosSortWrapper) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w GetPostsInfosSortWrapper) Less(i, j int) bool { return w[i].TweetAmount < w[j].TweetAmount }

func TestGetPostsInfos(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestBoardHTML))
	b := pttifierLib.NewBoardCrawler(root)

	expects := GetPostsInfosSortWrapper{
		{
			pttifierLib.BaseInfo{
				URL:    "https://www.ptt.cc/bbs/LoL/M.1457622048.A.D80.html",
				Title:  "[公告] LoL 樂透開獎",
				Author: "[彩券]",
				Date:   "3/10",
			},
			12,
		},
		{
			pttifierLib.BaseInfo{
				URL:    "https://www.ptt.cc/bbs/LoL/M.1457622048.A.Z80.html",
				Title:  "[公告] LoL 小樂透開獎",
				Author: "[彩券]",
				Date:   "3/11",
			},
			11,
		},
		{
			pttifierLib.BaseInfo{
				URL:    "https://www.ptt.cc/bbs/LoL/M.1457622282.A.D4F.html",
				Title:  "Re: [問題] AlphaGo打LoL的話會怎麼樣",
				Author: "lzainside",
				Date:   "3/10",
			},
			4,
		},
		{
			pttifierLib.BaseInfo{
				URL:    "https://www.ptt.cc/bbs/LoL/M.1457622407.A.C98.html",
				Title:  "[問題] 關於四名中路的比較！",
				Author: "lizst980074",
				Date:   "3/10",
			},
			2,
		},
		{
			pttifierLib.BaseInfo{
				URL:    "https://www.ptt.cc/bbs/LoL/M.1457622551.A.650.html",
				Title:  "[閒聊] 對於XG挺失望的",
				Author: "iamfenixsc",
				Date:   "3/10",
			},
			1,
		},
	}

	actuals := b.GetPostsInfos()
	sort.Sort(GetPostsInfosSortWrapper(actuals))
	sort.Sort(expects)

	for i, expect := range expects {
		actual := actuals[i]
		if actual.URL != expect.URL {
			t.Errorf("GetPostsInfos: expect-> %q, actual-> %q", expect.URL, actual.URL)
		}
		if actual.Title != expect.Title {
			t.Errorf("GetPostsInfos: expect-> %q, actual-> %q", expect.Title, actual.Title)
		}
		if actual.Author != expect.Author {
			t.Errorf("GetPostsInfos: expect-> %q, actual-> %q", expect.Author, actual.Author)
		}
		if actual.Date != expect.Date {
			t.Errorf("GetPostsInfos: expect-> %q, actual-> %q", expect.Date, actual.Date)
		}
		if actual.TweetAmount != expect.TweetAmount {
			t.Errorf("GetPostsInfos: expect-> %q, actual-> %q", expect.TweetAmount, actual.TweetAmount)
		}
	}
}
