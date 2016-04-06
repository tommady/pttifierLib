package pttifierLib

import (
	"errors"
	"strings"

	"golang.org/x/net/html"
)

type StrMatcher func(s, chars string) bool
type IntMatcher func(n, comparison int) bool
type RuleSetting func(*Parser)
type Parsers []*Parser
type Results []*PostBaseInfo

type TextRule struct {
	Title       string
	Author      string
	Content     string
	TweetAmount int
}

type MatcherRule struct {
	TitleMatcher       StrMatcher
	AuthorMatcher      StrMatcher
	ContentMatcher     StrMatcher
	TweetAmountMatcher IntMatcher
}

type Parser struct {
	TextRule
	MatcherRule
	skipThisParse bool
	err           error
}

var (
	ErrRootNodeNil = errors.New("pttifier.parser: input root node is nil")
)

func NewParser(settings ...RuleSetting) *Parser {
	p := new(Parser)
	for _, setting := range settings {
		setting(p)
	}

	if p.TitleMatcher == nil {
		p.TitleMatcher = strings.Contains
	}
	if p.AuthorMatcher == nil {
		p.AuthorMatcher = strings.EqualFold
	}
	if p.TweetAmountMatcher == nil {
		p.TweetAmountMatcher = LessThanOrEqualToComparison
	}
	if p.ContentMatcher == nil {
		p.ContentMatcher = strings.Contains
	}

	return p
}

func (p *Parser) setErr(err error) {
	if p.err == nil {
		p.err = err
	}
}

func (p *Parser) Err() error {
	return p.err
}

func SetParserTitle(title string) RuleSetting {
	return func(p *Parser) {
		p.Title = title
	}
}

func SetParserAuthor(author string) RuleSetting {
	return func(p *Parser) {
		p.Author = author
	}
}

func SetParserTweetAmount(tweetAmount string) RuleSetting {
	return func(p *Parser) {
		p.TweetAmount = TweetAmountStringToInt(tweetAmount)
	}
}

func SetParserContent(content string) RuleSetting {
	return func(p *Parser) {
		p.Content = content
	}
}

func SetParserTitleMatcher(matcher StrMatcher) RuleSetting {
	return func(p *Parser) {
		p.TitleMatcher = matcher
	}
}

func SetParserAuthorMatcher(matcher StrMatcher) RuleSetting {
	return func(p *Parser) {
		p.AuthorMatcher = matcher
	}
}

func SetParserContentMatcher(matcher StrMatcher) RuleSetting {
	return func(p *Parser) {
		p.ContentMatcher = matcher
	}
}

func (p *Parser) Parsing(root *html.Node) (results Results) {
	board := NewBoardCrawler(root)

	for _, article := range board.GetArticlesInfos() {
		p.skipThisParse = false
		p.compareTitle(article.Title)
		p.compareAuthor(article.Author)
		p.compareTweetAmount(article.TweetAmount)
		if p.skipThisParse {
			continue
		}

		results = append(results, article)
	}

	return results
}

func (p *Parser) compareTitle(title string) {
	if p.skipThisParse {
		return
	}

	if p.Title != "" && !p.TitleMatcher(title, p.Title) {
		p.skipThisParse = true
	}
}

func (p *Parser) compareAuthor(author string) {
	if p.skipThisParse {
		return
	}

	if p.Author != "" && !p.AuthorMatcher(author, p.Author) {
		p.skipThisParse = true
	}
}

func (p *Parser) compareTweetAmount(tweetAmount int) {
	if p.skipThisParse {
		return
	}

	if p.TweetAmount != 0 && !p.TweetAmountMatcher(tweetAmount, p.TweetAmount) {
		p.skipThisParse = true
	}
}

func (p *Parser) comparePostContent(URL string) {
	if p.skipThisParse {
		return
	}

	root, err := GetNodeFromLink(URL)
	if err != nil {
		p.skipThisParse = true
		return
	}

	postCrawler := NewPostCrawler(root)
	content := postCrawler.GetArticleContents()
	if p.Content != "" && !p.ContentMatcher(content, p.Content) {
		p.skipThisParse = true
	}
}
