package pttifierLib

import (
	"errors"
	"net"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

var (
	ErrMainContainerNodeNil = errors.New("pttifierLib.postCrawler: main container node is nil")
)

type Tweet struct {
	Author  string
	Content string
	Date    string
	Tag     string
}

type PostCrawler struct {
	mainContainerNode *html.Node
	err               error
	skipThisTweet     bool
}

const (
	TweetTagPraise = "推"
	TweetTagNormal = "→"
	TweetTagBoo    = "噓"
)

func NewPostCrawler(root *html.Node) *PostCrawler {
	p := new(PostCrawler)

	p.mainContainerNode = GetRListNode(root)
	if p.mainContainerNode == nil {
		p.setErr(ErrMainContainerNodeNil)
	}

	return p
}

func (p *PostCrawler) setErr(err error) {
	if p.err == nil {
		p.err = err
	}
}

func (p *PostCrawler) Err() error {
	return p.err
}

func (p *PostCrawler) GetAuthor() string {
	return p.getArticleMetalineData("作者")
}

func (p *PostCrawler) GetTitle() string {
	return p.getArticleMetalineData("標題")
}

func (p *PostCrawler) GetDate() string {
	return p.getArticleMetalineData("時間")
}

func (p *PostCrawler) getArticleMetalineData(keyword string) (data string) {
	if p.err == ErrMainContainerNodeNil {
		return
	}

	n, ok := scrape.Find(p.mainContainerNode, func(n *html.Node) bool {
		return scrape.Attr(n, "class") == "article-metaline" &&
			scrape.Text(n.FirstChild) == keyword
	})
	if ok {
		data = scrape.Text(n.LastChild)
	}

	return
}

func (p *PostCrawler) GetIP() (ip net.IP) {
	if p.err == ErrMainContainerNodeNil {
		return
	}

	n, ok := scrape.Find(p.mainContainerNode, func(n *html.Node) bool {
		return scrape.Attr(n, "class") == "f2" &&
			strings.HasPrefix(scrape.Text(n), "※ 發信站: 批踢踢實業坊(ptt.cc)")
	})
	if ok {
		s := scrape.Text(n)
		ts := strings.TrimPrefix(s, "※ 發信站: 批踢踢實業坊(ptt.cc), 來自: ")
		ip = net.ParseIP(ts)
	}

	return
}

func (p *PostCrawler) GetURL() (url string) {
	if p.err == ErrMainContainerNodeNil {
		return
	}

	n, ok := scrape.Find(p.mainContainerNode, func(n *html.Node) bool {
		return scrape.Attr(n, "class") == "f2" &&
			strings.HasPrefix(scrape.Text(n), "※ 文章網址: ")
	})
	if ok {
		s := scrape.Text(n)
		url = strings.TrimPrefix(s, "※ 文章網址: ")
	}

	return
}

func (p *PostCrawler) GetContent() string {
	if p.err == ErrMainContainerNodeNil {
		return ""
	}

	content := scrape.FindAll(p.mainContainerNode, func(n *html.Node) bool {
		return n.Type == html.TextNode &&
			scrape.Attr(n.Parent, "class") != "article-meta-tag" &&
			scrape.Attr(n.Parent, "class") != "article-meta-value" &&
			scrape.Attr(n.Parent, "class") != "f2" &&
			scrape.Attr(n.Parent.Parent, "class") != "f2" &&
			scrape.Attr(n.Parent, "class") != "f6" &&
			scrape.Attr(n.Parent.Parent, "class") != "f6" &&
			scrape.Attr(n.Parent.Parent, "class") != "push"
	})

	joiner := func(ns []*html.Node) string {
		c := ""
		for _, n := range ns {
			c += strings.TrimSpace(n.Data)
		}
		return c
	}

	return joiner(content)
}

func (p *PostCrawler) GetTweets() (tweets []*Tweet) {
	if p.mainContainerNode == nil {
		return
	}

	ts := scrape.FindAll(p.mainContainerNode, scrape.ByClass("push"))
	for _, t := range ts {
		p.skipThisTweet = false

		author := p.getTweetAuthor(t)
		content := p.getTweetContent(t)
		date := p.getTweetDate(t)
		tag := p.getTweetTag(t)

		if p.skipThisTweet {
			continue
		}

		tweet := new(Tweet)
		tweet.Author = author
		tweet.Content = content
		tweet.Date = date
		tweet.Tag = tag
		tweets = append(tweets, tweet)
	}

	return
}

func (p *PostCrawler) getTweetAuthor(tweet *html.Node) (tweetAuthor string) {
	if p.skipThisTweet {
		return
	}

	a, ok := scrape.Find(tweet, scrape.ByClass("push-userid"))
	if !ok {
		p.skipThisTweet = true
		return
	}

	tweetAuthor = scrape.Text(a)
	return
}

func (p *PostCrawler) getTweetContent(tweet *html.Node) (tweetContent string) {
	if p.skipThisTweet {
		return
	}

	c, ok := scrape.Find(tweet, scrape.ByClass("push-content"))
	if !ok {
		p.skipThisTweet = true
		return
	}

	tweetContent = strings.Trim(scrape.Text(c), ": ")
	return
}

func (p *PostCrawler) getTweetDate(tweet *html.Node) (tweetDate string) {
	if p.skipThisTweet {
		return
	}

	d, ok := scrape.Find(tweet, scrape.ByClass("push-ipdatetime"))
	if !ok {
		p.skipThisTweet = true
		return
	}

	tweetDate = scrape.Text(d)
	return
}

func (p *PostCrawler) getTweetTag(tweet *html.Node) (tweetTag string) {
	if p.skipThisTweet {
		return
	}

	t, ok := scrape.Find(tweet, scrape.ByClass("push-tag"))
	if !ok {
		p.skipThisTweet = true
		return
	}

	tweetTag = scrape.Text(t)
	return
}
