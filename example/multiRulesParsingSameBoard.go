package main

import (
	"fmt"
	"log"

	ptlib "github.com/tommady/pttifierLib"
)

func main() {
	link := ptlib.WrapBoardPageLink("WomenTalk", "5014")
	root, err := ptlib.GetNodeFromLink(link)
	if err != nil {
		log.Fatalf("GG on get page: %v", err)
	}

	board := ptlib.NewBoardCrawler(root)
	posts := board.GetPostsInfosAndArticles()
	if board.Err() != nil {
		log.Fatalf("GG on board: %v", board.Err())
	}

	parserList := []*ptlib.Parser{
		ptlib.NewParser(
			ptlib.SetParserTitle("女"),
		),
		ptlib.NewParser(
			ptlib.SetParserTitle("男"),
			ptlib.SetParserAuthor("a2006lkk"),
		),
	}

	results := []*ptlib.Result{}
	resultCh := make(chan []*ptlib.Result, len(parserList))
	for _, parser := range parserList {
		go func(parser *ptlib.Parser) {
			results := parser.Parsing(posts)
			resultCh <- results
		}(parser)
	}

	for i := 0; i < len(parserList); i++ {
		select {
		case rs := <-resultCh:
			results = append(results, rs...)
		}
	}

	for _, r := range results {
		fmt.Println(r.Date, r.Title, r.Author)
	}
}
