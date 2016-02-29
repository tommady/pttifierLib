package pttifierLib

import (
    "net/http"
)

type UserInfo struct {
    Email    string
}

type Rule struct {
    UserInfo
	Board    string
	TitleKey string
}

type Result struct {
    UserInfo
	Url   string
	Title string
}

const (
	pttBaseUrl         = "https://www.ptt.cc/bbs/"
	defaultParsingPage = "/index"
)

func DoCrawling(rule *Rule, resultCh chan Result, errorCh chan PttiferErr) {
	// doing crawling
}

// for some specific board need over 18 years old check
func getPttRespond(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle error
	}

	req.Header.Set("Cookie", "over18=1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle error
	}

	return res
}
