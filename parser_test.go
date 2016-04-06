package pttifierLib_test

import (
	"strings"
	"testing"

	"github.com/tommady/pttifierLib"
	"golang.org/x/net/html"
)

func TestParsing(t *testing.T) {
	root, _ := html.Parse(strings.NewReader(TestBoardHTML))
	parseRules := pttifierLib.Parsers{}

	parseRules = append(parseRules, pttifierLib.NewParser(
		pttifierLib.SetParserTitle("失望"),
	))
	parseRules = append(parseRules, pttifierLib.NewParser(
		pttifierLib.SetParserTitle("小樂透"),
		pttifierLib.SetParserAuthor("[彩券]"),
	))
	parseRules = append(parseRules, pttifierLib.NewParser(
		pttifierLib.SetParserAuthor("[彩券]"),
	))

	expectResults := []pttifierLib.Results{
		{
			{
				"https://www.ptt.cc/bbs/LoL/M.1457622551.A.650.html",
				"[閒聊] 對於XG挺失望的",
				"iamfenixsc",
				"3/10",
				1,
			},
		},
		{
			{
				"https://www.ptt.cc/bbs/LoL/M.1457622048.A.Z80.html",
				"[公告] LoL 小樂透開獎",
				"[彩券]",
				"3/11",
				11,
			},
		},
		{
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
		},
	}

	for i, parseRule := range parseRules {
		actualResults := parseRule.Parsing(root)
		if len(actualResults) == 0 {
			t.Errorf("Parsing[%d]: expect at least 1 result but found 0", i)
		} else {
			for j := 0; j < len(actualResults); j++ {
				actualResult := actualResults[j]
				expectResult := expectResults[i][j]

				if actualResult.Title != expectResult.Title {
					t.Errorf("Parsing[%d]: expected-> %q, actual-> %q",
						i,
						expectResult.Title,
						actualResult.Title,
					)
				}

				if actualResult.URL != expectResult.URL {
					t.Errorf("Parsing[%d]: expected-> %q, actual-> %q",
						i,
						expectResult.URL,
						actualResult.URL,
					)
				}

				if actualResult.Author != expectResult.Author {
					t.Errorf("Parsing[%d]: expected-> %q, actual-> %q",
						i,
						expectResult.Author,
						actualResult.Author,
					)
				}

				if actualResult.Date != expectResult.Date {
					t.Errorf("Parsing[%d]: expected-> %q, actual-> %q",
						i,
						expectResult.Date,
						actualResult.Date,
					)
				}
			}
		}
	}
}
