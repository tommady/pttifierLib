package pttifierLib

import "strings"

type StrMatcher func(s, chars string) bool
type IntMatcher func(n, comparison int) bool
type RuleSetting func(*Parser)
type Result BaseInfo

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
	err error
}

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

func (p *Parser) Parsing(posts []*BoardInfoAndArticle) (results []*Result) {
	if posts == nil {
		return
	}

	resultCh := make(chan *Result, len(posts))
	for _, post := range posts {
		go func(post *BoardInfoAndArticle) {
			if !p.compareTitle(post.Title) ||
				!p.compareAuthor(post.Author) ||
				!p.compareTweetAmount(post.TweetAmount) ||
				!p.comparePostContent(post.Content) {
				resultCh <- nil
				return
			}

			resultCh <- &Result{
				URL:    post.URL,
				Title:  post.Title,
				Author: post.Author,
				Date:   post.Date,
			}
		}(post)
	}

	for i := 0; i < len(posts); i++ {
		select {
		case r := <-resultCh:
			if r != nil {
				results = append(results, r)
			}
		}
	}

	return results
}

func (p *Parser) compareTitle(title string) bool {
	if p.Title != "" && !p.TitleMatcher(title, p.Title) {
		return false
	}

	return true
}

func (p *Parser) compareAuthor(author string) bool {
	if p.Author != "" && !p.AuthorMatcher(author, p.Author) {
		return false
	}

	return true
}

func (p *Parser) compareTweetAmount(tweetAmount int) bool {
	if p.TweetAmount != 0 && !p.TweetAmountMatcher(tweetAmount, p.TweetAmount) {
		return false
	}

	return true
}

func (p *Parser) comparePostContent(content string) bool {
	if p.Content != "" && !p.ContentMatcher(content, p.Content) {
		return false
	}

	return true
}
