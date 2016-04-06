package pttifierLib_test

import (
	"strings"
	"testing"

	"github.com/tommady/pttifierLib"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

func TestWrapBoardPageLink(t *testing.T) {
	tests := []struct {
		inBoard  string // Regardless of capitalization
		inPage   string
		expected string
	}{
		{"Joke", "", "https://www.ptt.cc/bbs/Joke/index.html"},
		{"gossiping", "689", "https://www.ptt.cc/bbs/gossiping/index689.html"},
	}

	for _, tt := range tests {
		actual := pttifierLib.WrapBoardPageLink(tt.inBoard, tt.inPage)
		if actual != tt.expected {
			t.Errorf("WrapBoardPageLink(%q, %q): expect-> %q, actual-> %q",
				tt.inBoard,
				tt.inPage,
				tt.expected,
				actual,
			)
		}
	}
}

func TestGetActionBarNode(t *testing.T) {
	tests := []struct {
		inRoot   *html.Node
		expected *html.Node
	}{
		{nil, nil},
	}

	for _, tt := range tests {
		actual := pttifierLib.GetActionBarNode(tt.inRoot)
		if actual != tt.expected {
			t.Errorf("GetActionBarNode: expect-> nil, actual-> not nil")
		}
	}
}

func TestGetRListNode(t *testing.T) {
	tests := []struct {
		inRoot   *html.Node
		expected *html.Node
	}{
		{nil, nil},
	}

	for _, tt := range tests {
		actual := pttifierLib.GetRListNode(tt.inRoot)
		if actual != tt.expected {
			t.Errorf("GetRListNode: expect-> nil, actual-> not nil")
		}
	}
}

func TestRemoveBottumAnnouncements(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestBoardHTML))
	rListNode := pttifierLib.GetRListNode(root)
	pttifierLib.RemoveBottumAnnouncements(rListNode)

	_, found := scrape.Find(rListNode, scrape.ByClass("r-list-sep"))
	if found {
		t.Errorf("RemoveBottumAnnouncements: expect no r-list-sep node but still has")
	}
}

func TestTweetAmountStringToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"爆", 100},
		{"X6", -60},
		{"47", 47},
	}

	for _, tt := range tests {
		actual := pttifierLib.TweetAmountStringToInt(tt.input)
		if actual != tt.expected {
			t.Errorf("TweetAmountStringToInt: expect-> %d, actual-> %d", tt.expected, actual)
		}
	}
}
