package pttifierLib

import (
	"errors"
	"strconv"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type BoardInfoAndArticle struct {
	*BoardInfo
	Content string   `json:"content"`
	Tweets  []*Tweet `json:"tweets"`
}

type BoardInfo struct {
	BaseInfo
	TweetAmount int `json:"tweet_amount"`
}

type BoardCrawler struct {
	actionBarNode   *html.Node
	rListNode       *html.Node
	err             error
	skipThisArticle bool
}

var (
	ErrActionBarNodeNil = errors.New("pttifier.boardCrawler: action bar node is nil")
	ErrRListNodeNil     = errors.New("pttifier.boardCrawler: R list node is nil")
)

func NewBoardCrawler(root *html.Node) *BoardCrawler {
	b := new(BoardCrawler)

	b.actionBarNode = GetActionBarNode(root)
	if b.actionBarNode == nil {
		b.setErr(ErrActionBarNodeNil)
	}

	b.rListNode = GetRListNode(root)
	if b.rListNode == nil {
		b.setErr(ErrRListNodeNil)
	}

	RemoveBottumAnnouncements(b.rListNode)
	return b
}

func (b *BoardCrawler) setErr(err error) {
	if b.err == nil {
		b.err = err
	}
}

func (b *BoardCrawler) Err() error {
	return b.err
}

func (b *BoardCrawler) GetCurrPageLinkNum() int {
	strPreLink := b.getPageLink("上頁")
	strNewestLink := b.getPageLink("最新")
	strNewestLink = strings.TrimSuffix(strNewestLink, ".html")
	strPreLink = strings.TrimPrefix(strPreLink, strNewestLink)
	strPreLink = strings.TrimSuffix(strPreLink, ".html")
	intPreLink, err := strconv.Atoi(strPreLink)
	if err != nil {
		return 0
	}

	return intPreLink + 1
}

func (b *BoardCrawler) GetPrevPageLink() string {
	return b.getPageLink("上頁")
}

func (b *BoardCrawler) GetNextPageLink() string {
	return b.getPageLink("下頁")
}

func (b *BoardCrawler) getPageLink(pageText string) (pageLink string) {
	if b.err == ErrActionBarNodeNil {
		return
	}

	ns := scrape.FindAll(b.actionBarNode, scrape.ByClass("wide"))
	for _, n := range ns {
		text := scrape.Text(n)
		link := scrape.Attr(n, "href")
		if link != "" {
			if strings.Contains(text, pageText) {
				pageLink = PttBaseURL + link
				break
			}
		}
	}

	return
}

func (b *BoardCrawler) GetPostsInfosAndArticles() []*BoardInfoAndArticle {
	if b.err == ErrRListNodeNil {
		return nil
	}
	infos := b.GetPostsInfos()
	infosAndArticles := []*BoardInfoAndArticle{}
	for _, info := range infos {
		post, err := GetNodeFromLink(info.URL)
		if err != nil {
			b.setErr(err)
			return nil
		}
		postCrawler := NewPostCrawler(post)
		content := postCrawler.GetContent()
		tweets := postCrawler.GetTweets()
		if postCrawler.Err() != nil {
			b.setErr(err)
			return nil
		}
		infosAndArticles = append(infosAndArticles, &BoardInfoAndArticle{
			BoardInfo: info,
			Content:   content,
			Tweets:    tweets,
		})
	}
	return infosAndArticles
}

func (b *BoardCrawler) GetPostsInfos() (infos []*BoardInfo) {
	if b.err == ErrRListNodeNil {
		return
	}

	articles := scrape.FindAll(b.rListNode, scrape.ByClass("r-ent"))
	for i := len(articles) - 1; i >= 0; i-- {
		b.skipThisArticle = false

		title := b.getTitle(articles[i])
		author := b.getAuthor(articles[i])
		url := b.getURL(articles[i])
		date := b.getDate(articles[i])
		tweetAmount := b.getTweetAmount(articles[i])

		if b.skipThisArticle {
			continue
		}

		info := new(BoardInfo)
		info.Title = title
		info.URL = url
		info.Author = author
		info.Date = date
		info.TweetAmount = tweetAmount
		infos = append(infos, info)
	}

	return
}

func (b *BoardCrawler) getTitle(article *html.Node) (title string) {
	if b.skipThisArticle {
		return
	}
	t, ok := scrape.Find(article, scrape.ByTag(atom.A))
	if !ok {
		// post has been deleted
		b.skipThisArticle = true
		return
	}

	title = scrape.Text(t)
	return
}

func (b *BoardCrawler) getURL(article *html.Node) (url string) {
	if b.skipThisArticle {
		return
	}

	t, ok := scrape.Find(article, scrape.ByTag(atom.A))
	if !ok {
		// post has been deleted
		b.skipThisArticle = true
		return
	}

	url = PttBaseURL + scrape.Attr(t, "href")
	return
}

func (b *BoardCrawler) getAuthor(article *html.Node) (author string) {
	if b.skipThisArticle {
		return
	}

	a, ok := scrape.Find(article, scrape.ByClass("author"))
	if !ok {
		b.skipThisArticle = true
		return
	}

	author = scrape.Text(a)
	return
}

func (b *BoardCrawler) getDate(article *html.Node) (date string) {
	if b.skipThisArticle {
		return
	}

	d, ok := scrape.Find(article, scrape.ByClass("date"))
	if !ok {
		b.skipThisArticle = true
		return
	}

	date = scrape.Text(d)
	return
}

func (b *BoardCrawler) getTweetAmount(article *html.Node) (tweetAmount int) {
	if b.skipThisArticle {
		return
	}

	n, ok := scrape.Find(article, scrape.ByClass("nrec"))
	if !ok {
		b.skipThisArticle = true
		return
	}

	tweetAmount = TweetAmountStringToInt(scrape.Text(n))
	return
}
